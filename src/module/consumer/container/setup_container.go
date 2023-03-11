package container

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/filestorage"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/repository/mysql"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/moduleregistry"
)

func SetupContainer(moduleRegistry moduleregistry.ModuleRegistry) Container {
	appConfig := config.LoadByEnv()
	aesCBCCrypto := NewCryptoUtility(appConfig)
	validate := validator.New()
	db := NewMysqlDb(appConfig)
	consumerRepository := mysql.NewConsumerRepository(db)
	minio := NewMinioClient(appConfig)
	fileStorage := filestorage.NewFileBucketClient(minio)
	consumerUsecase := consumer.NewConsumerUsecase(appConfig, aesCBCCrypto, validate, consumerRepository, fileStorage)
	container := newContainer(moduleRegistry, appConfig, consumerUsecase)
	return container
}

func registerSharedModuleRegistry(container Container) {
	moduleregistry.RegisterSharedModule(container.moduleRegistry, "consumer/approve-request-loan", container.ConsumerUsecase.ApproveRequestLoan)
	moduleregistry.RegisterSharedModule(container.moduleRegistry, "consumer/add-tenor-limit", container.ConsumerUsecase.AddTenorLimit)
}
