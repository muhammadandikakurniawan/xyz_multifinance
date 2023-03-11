package consumer_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/filestorage"
	. "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/crypto/aes"
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	CRYPTO_AES_KEY = "AESKEYXYZMULTIFINANCE12023030923"
	CRYPTO_AES_IV  = "acfa7a047800b2f2"
	aesCrypto, _   = aes.NewAesCbc(CRYPTO_AES_IV, CRYPTO_AES_KEY)
)

func Test_Register(t *testing.T) {
	t.Run("successs", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{
			BucketConsumerImage: "image",
		}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour
		fileStorageClient.On("CreateDirectory", mock.Anything, mock.Anything).Return(nil)
		fileStorageClient.On("UploadBase64", mock.Anything, mock.Anything).Return(filestorage.UploadResultOpt{FilePath: "img-url"}, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.RequestCreateNewConsumerDto{
			DateOfBirth:    "11/10/2000",
			Fullname:       "consumer two",
			Legalname:      "consumer two legal name",
			NIK:            "xM3NJ/yVw5HMz1i1NRbfA4+w79ofUhh3lodenqEW3z8=",
			PlaceOfBirth:   "bekasi",
			Salary:         6500000,
			KtpImageBase64: "ktpbase64",
			SelfeBase64:    "selfiebase64",
		}
		result, err := consumerUsecase.Register(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.SUCCESS), result.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{
			BucketConsumerImage: "image",
		}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.RequestCreateNewConsumerDto{
			DateOfBirth:    "2001/05/24",
			Fullname:       "asdasd asd",
			Legalname:      "consumer two legal name",
			NIK:            "",
			PlaceOfBirth:   "bekasi",
			Salary:         6500000,
			KtpImageBase64: "ktpbase64",
			SelfeBase64:    "selfiebase64",
		}
		result, err := consumerUsecase.Register(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.ERROR_BAD_REQUEST), result.StatusCode)
	})
}

func Test_AddTenorLimit(t *testing.T) {
	t.Run("successs", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{
			BucketConsumerImage: "image",
		}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour

		aggregateRoot := entity.ConsumerEntity{
			Id: "0a68ef44-53eb-4374-8503-7b61658be2af",
		}
		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			ConsumerId: aggregateRoot.Id,
			Month:      1,
			LimitValue: 200000,
		})
		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			ConsumerId: aggregateRoot.Id,
			Month:      2,
			LimitValue: 300000,
		})
		getResult := consumer.BuildConsumerAggregate(aggregateRoot)

		consumerRepo.On("FindTenorLimitByConsumerId", mock.Anything, mock.Anything).Return(&getResult, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.AddTenorLmitRequestDto{
			ConsumerId: aggregateRoot.Id,
			Limits: []dto.TenorLimitDto{
				{
					Month: 1,
					Value: 200000,
				},
				{
					Month: 2,
					Value: 300000,
				},
			},
		}
		result, err := consumerUsecase.AddTenorLimit(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.SUCCESS), result.StatusCode)
	})
}

func Test_RequestLoan(t *testing.T) {
	t.Run("successs", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour

		aggregateRoot := entity.ConsumerEntity{
			Id: "0a68ef44-53eb-4374-8503-7b61658be2af",
		}
		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			ConsumerId: aggregateRoot.Id,
			Month:      1,
			LimitValue: 200000,
		})
		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			ConsumerId: aggregateRoot.Id,
			Month:      2,
			LimitValue: 300000,
		})
		getResult := consumer.BuildConsumerAggregate(aggregateRoot)

		consumerRepo.On("FindTenorLimitByConsumerId", mock.Anything, mock.Anything).Return(&getResult, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.RequestLoanDto{
			ConsumerId:  aggregateRoot.Id,
			AssetName:   "asset test",
			OTR:         35000000,
			Installment: 1500000,
			Interest:    500000,
			AdminFee:    450000,
		}
		result, err := consumerUsecase.RequestLoan(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.SUCCESS), result.StatusCode)
	})

	t.Run("failed, tenor limit not found", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour

		aggregateRoot := entity.ConsumerEntity{
			Id: "0a68ef44-53eb-4374-8503-7b61658be2af",
		}
		getResult := consumer.BuildConsumerAggregate(aggregateRoot)

		consumerRepo.On("FindTenorLimitByConsumerId", mock.Anything, mock.Anything).Return(&getResult, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.RequestLoanDto{
			ConsumerId:  aggregateRoot.Id,
			AssetName:   "asset test",
			OTR:         35000000,
			Installment: 1500000,
			Interest:    500000,
			AdminFee:    450000,
		}

		result, err := consumerUsecase.RequestLoan(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.ERROR_BAD_REQUEST), result.StatusCode)
		assert.Equal(t, "tenor limit not found", result.Message)

	})
}

func Test_ApproveRequestLoan(t *testing.T) {
	t.Run("successs", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour

		aggregateRoot := entity.ConsumerEntity{
			Id: "0a68ef44-53eb-4374-8503-7b61658be2af",
		}
		aggregateRoot.AddRequestLoan(entity.RequestLoanEntity{
			Id:         1,
			ConsumerId: aggregateRoot.Id,
		})
		getResult := consumer.BuildConsumerAggregate(aggregateRoot)

		consumerRepo.On("FindRequestLoanById", mock.Anything, mock.Anything).Return(&getResult, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.ApprovalResponseDataDto{
			Id:         1,
			IsApproved: true,
		}
		result, err := consumerUsecase.ApproveRequestLoan(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, true, result.Success)
		assert.Equal(t, http.StatusOK, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.SUCCESS), result.StatusCode)
		assert.Equal(t, true, result.Data.IsApproved)
	})

	t.Run("failed, request already approve", func(t *testing.T) {

		ctx := context.Background()
		appConfig := config.AppConfig{}
		jsonValidator := validator.New()
		consumerRepo := &mocks.ConsumerRepository{}
		fileStorageClient := &mocks.FileStorage{}

		// START Setup behaviour

		aggregateRoot := entity.ConsumerEntity{
			Id: "0a68ef44-53eb-4374-8503-7b61658be2af",
		}
		isApprove := true
		aggregateRoot.AddRequestLoan(entity.RequestLoanEntity{
			Id:         1,
			ConsumerId: aggregateRoot.Id,
			IsApproved: &isApprove,
		})
		getResult := consumer.BuildConsumerAggregate(aggregateRoot)

		consumerRepo.On("FindRequestLoanById", mock.Anything, mock.Anything).Return(&getResult, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, aesCrypto, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.ApprovalResponseDataDto{
			Id:         1,
			IsApproved: true,
		}
		result, err := consumerUsecase.ApproveRequestLoan(ctx, requestData)
		assert.Nil(t, err)
		assert.Equal(t, false, result.Success)
		assert.Equal(t, http.StatusBadRequest, result.HttpStatusCode)
		assert.Equal(t, string(sharedErr.ERROR_BAD_REQUEST), result.StatusCode)
		assert.Equal(t, "cannot set approval for this request", result.Message)

	})
}
