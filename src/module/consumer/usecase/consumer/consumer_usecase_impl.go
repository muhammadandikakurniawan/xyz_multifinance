package consumer

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/filestorage"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/repository"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/crypto/aes"
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
	"golang.org/x/sync/errgroup"
)

func NewConsumerUsecase(
	appConfig config.AppConfig,
	encrypter aes.AesCBCCrypto,
	structValidation *validator.Validate,
	consumerRepository repository.ConsumerRepository,
	fileStorage filestorage.FileStorage,
) ConsumerUsecase {

	if structValidation == nil {
		panic("structValidation cannot be null")
	}
	if consumerRepository == nil {
		panic("consumerRepository cannot be null")
	}
	if fileStorage == nil {
		panic("fileStorage cannot be null")
	}

	return &consumerUsecaseImpl{
		appConfig:          appConfig,
		encrypter:          encrypter,
		structValidation:   structValidation,
		consumerRepository: consumerRepository,
		fileStorage:        fileStorage,
	}
}

type consumerUsecaseImpl struct {
	appConfig          config.AppConfig
	encrypter          aes.AesCBCCrypto
	structValidation   *validator.Validate
	consumerRepository repository.ConsumerRepository
	fileStorage        filestorage.FileStorage
}

func (uc consumerUsecaseImpl) DecryptNik(nik string) (plainTextNik string, err error) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(nik)
	if err != nil {
		return
	}
	plainTextNiB, err := uc.encrypter.Decrypt(rawDecodedText)
	if err != nil {
		return
	}
	plainTextNik = string(plainTextNiB)
	return
}

func (uc consumerUsecaseImpl) Register(ctx context.Context, requestData dto.RequestCreateNewConsumerDto) (result model.BaseResponseModel[dto.ConsumerId], err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	if validationErr := uc.structValidation.Struct(requestData); validationErr != nil {
		validationErr = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "invalid request", validationErr.Error())
		result.SetError(validationErr)
		return
	}

	// decrypt nik
	requestData.NIK, err = uc.DecryptNik(requestData.NIK)
	if err != nil {
		err = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "invalid nik", "invalid nik")
		result.SetError(err)
		err = nil
		return
	}

	var ktpImgUrl, selfieImgUrl string
	consumerData, err := requestData.TransformToEntity(ktpImgUrl, selfieImgUrl)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	// upload ktp and selfie
	ktpImgUrl, selfieImgUrl, err = uc.UploadKtpAndSelfie(ctx, requestData)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}
	consumerData.KtpImageUrl = ktpImgUrl
	consumerData.SelfieUrl = selfieImgUrl

	consumerAggregate, err := consumer.CreateNewConsumerAggregate(consumerData)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	err = uc.consumerRepository.Save(ctx, &consumerAggregate)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	result.Data = dto.ConsumerId{
		Id: consumerAggregate.GetAggregateRoot().Id,
	}

	return
}

func (uc consumerUsecaseImpl) UploadKtpAndSelfie(ctx context.Context, requestData dto.RequestCreateNewConsumerDto) (ktpImageUrl, selfieImageUrl string, err error) {

	if err = uc.fileStorage.CreateDirectory(ctx, uc.appConfig.BucketConsumerImage); err != nil {
		return
	}

	eg := errgroup.Group{}
	eg.Go(func() (err error) {
		uploadRes, err := uc.fileStorage.UploadBase64(ctx, filestorage.UploadFileOpt{
			Bucket:       uc.appConfig.BucketConsumerImage,
			Filename:     fmt.Sprintf("%d-ktp.jpg", time.Now().UnixNano()),
			ContentType:  "image/jpg",
			Base64String: requestData.KtpImageBase64,
		})
		if err != nil {
			return
		}
		ktpImageUrl = uploadRes.FilePath
		return
	})

	eg.Go(func() (err error) {
		uploadRes, err := uc.fileStorage.UploadBase64(ctx, filestorage.UploadFileOpt{
			Bucket:       uc.appConfig.BucketConsumerImage,
			Filename:     fmt.Sprintf("%d-selfie.jpg", time.Now().UnixNano()),
			ContentType:  "image/jpg",
			Base64String: requestData.SelfeBase64,
		})
		if err != nil {
			return
		}
		selfieImageUrl = uploadRes.FilePath
		return
	})

	err = eg.Wait()

	return
}

func (uc consumerUsecaseImpl) RequestLoan(ctx context.Context, requestData dto.RequestLoanDto) (result model.BaseResponseModel[dto.RequestLoanDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	consumerAggregate, err := uc.consumerRepository.FindTenorLimitByConsumerId(ctx, requestData.ConsumerId)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	if consumerAggregate == nil {
		result.StatusCode = string(sharedErr.ERROR_BAD_REQUEST)
		result.HttpStatusCode = sharedErr.ERROR_BAD_REQUEST.ToHttpStatus()
		result.Message = "consumer not found"
		result.ErrorMessage = result.Message
		result.Success = false
		return
	}

	requestLoanData := requestData.TransformToEntity()
	if err = consumerAggregate.AddRequestLoan(&requestLoanData); err != nil {
		result.SetError(err)
		err = nil
		return
	}

	if err = uc.consumerRepository.Save(ctx, consumerAggregate); err != nil {
		result.SetError(err)
		err = nil
		return
	}

	requestData.ContractNumber = requestLoanData.ContractNumber
	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	result.Data = requestData
	return
}

func (uc consumerUsecaseImpl) ApproveRequestLoan(ctx context.Context, requestData dto.ApprovalResponseDataDto) (result model.BaseResponseModel[dto.ApprovalResponseDataDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	consumerAggregate, err := uc.consumerRepository.FindRequestLoanById(ctx, requestData.Id)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	if consumerAggregate == nil {
		result.StatusCode = string(sharedErr.ERROR_BAD_REQUEST)
		result.HttpStatusCode = sharedErr.ERROR_BAD_REQUEST.ToHttpStatus()
		result.Message = "data not found"
		result.ErrorMessage = result.Message
		result.Success = false
		return
	}

	err = consumerAggregate.ApproveRequestLoan(requestData.Id, requestData.IsApproved)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	err = uc.consumerRepository.Save(ctx, consumerAggregate)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}
	requestData.IsApproved = *consumerAggregate.GetAggregateRoot().MapRequestLoanById[requestData.Id].IsApproved
	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	result.Data = requestData
	return
}

func (uc consumerUsecaseImpl) AddTenorLimit(ctx context.Context, requestData dto.AddTenorLmitRequestDto) (result model.BaseResponseModel[dto.AddTenorLmitRequestDto], err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	consumerAggregate, err := uc.consumerRepository.FindTenorLimitByConsumerId(ctx, requestData.ConsumerId)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	if consumerAggregate == nil {
		result.StatusCode = string(sharedErr.ERROR_BAD_REQUEST)
		result.HttpStatusCode = sharedErr.ERROR_BAD_REQUEST.ToHttpStatus()
		result.Message = "consumer not found"
		result.ErrorMessage = result.Message
		result.Success = false
		return
	}

	tenorLimitEntities := []*entity.TenorLimitEntity{}
	for _, tl := range requestData.Limits {
		tlEntity := tl.TransformToEntity(requestData.ConsumerId)
		tenorLimitEntities = append(tenorLimitEntities, &tlEntity)
	}

	if err = consumerAggregate.AddTenorLimit(tenorLimitEntities...); err != nil {
		result.SetError(err)
		err = nil
		return
	}

	err = uc.consumerRepository.Save(ctx, consumerAggregate)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	result.Data = requestData
	return
}

func (uc consumerUsecaseImpl) GetListRequestLoan(ctx context.Context, requestData dto.GetListRequestLoanRequestDto) (result model.BaseResponseModel[[]dto.RequestLoanDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	listConsumer, paginationRes, err := uc.consumerRepository.SearchListRequestLoan(ctx, requestData.Pagination, requestData.Filter)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	resultData := dto.TransformFromConsumerEntities(listConsumer)
	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	result.PaginationResponseModel = paginationRes
	result.Data = resultData
	return
}

func (uc consumerUsecaseImpl) GetConsumer(ctx context.Context, consumerId string) (result model.BaseResponseModel[*dto.ConsumerDto], err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			result.SetError(err)
			err = nil
		}
	}()

	consumerAggregate, err := uc.consumerRepository.FindTenorLimitByConsumerId(ctx, consumerId)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	if consumerAggregate == nil {
		result.StatusCode = string(sharedErr.ERROR_BAD_REQUEST)
		result.HttpStatusCode = sharedErr.ERROR_BAD_REQUEST.ToHttpStatus()
		result.Message = "consumer not found"
		result.ErrorMessage = result.Message
		result.Success = false
		return
	}
	consumerAggregate.EncryptNik(uc.encrypter)
	result.Data = &dto.ConsumerDto{}
	result.Data.SetupFromEntity(consumerAggregate.GetAggregateRoot())
	result.Success = true
	result.Message = "successs"
	result.StatusCode = string(sharedErr.SUCCESS)
	result.HttpStatusCode = sharedErr.SUCCESS.ToHttpStatus()
	return
}
