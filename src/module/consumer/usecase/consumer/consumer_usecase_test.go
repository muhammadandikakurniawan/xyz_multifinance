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
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		fileStorageClient.On("UploadBase64", mock.Anything, mock.Anything).Return(filestorage.UploadResultOpt{FilePath: "img-url"}, nil)
		consumerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
		// END Setup behaviour

		consumerUsecase := NewConsumerUsecase(appConfig, jsonValidator, consumerRepo, fileStorageClient)

		requestData := dto.RequestCreateNewConsumerDto{
			DateOfBirth:    "11/10/2000",
			Fullname:       "consumer two",
			Legalname:      "consumer two legal name",
			NIK:            "9087960984956784",
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

		consumerUsecase := NewConsumerUsecase(appConfig, jsonValidator, consumerRepo, fileStorageClient)

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

		consumerUsecase := NewConsumerUsecase(appConfig, jsonValidator, consumerRepo, fileStorageClient)

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
