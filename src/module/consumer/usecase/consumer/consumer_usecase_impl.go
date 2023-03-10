package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/filestorage"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/repository"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
	"golang.org/x/sync/errgroup"
)

func NewConsumerUsecase(
	appConfig config.AppConfig,
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
		structValidation:   structValidation,
		consumerRepository: consumerRepository,
		fileStorage:        fileStorage,
	}
}

type consumerUsecaseImpl struct {
	appConfig          config.AppConfig
	structValidation   *validator.Validate
	consumerRepository repository.ConsumerRepository
	fileStorage        filestorage.FileStorage
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

	// upload ktp and selfie
	ktpImgUrl, selfieImgUrl, err := uc.UploadKtpAndSelfie(ctx, requestData)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

	consumerData, err := requestData.TransformToEntity(ktpImgUrl, selfieImgUrl)
	if err != nil {
		result.SetError(err)
		err = nil
		return
	}

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

	requestLoanData := requestData.TransformToEntity()
	if err = consumerAggregate.AddRequestLoan(&requestLoanData); err != nil {
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

func (uc consumerUsecaseImpl) ApproveRequestLoan(ctx context.Context, requestData dto.ConsumerDto) (result model.BaseResponseModel[dto.ApprovalResponseDataDto], err error) {

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
