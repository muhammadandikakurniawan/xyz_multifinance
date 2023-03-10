// go:build wireinject
//go:build wireinject
// +build wireinject

package container

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/filestorage"
	mysqlRepo "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/repository/mysql"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/database"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/minio"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/moduleregistry"
	"gorm.io/gorm"
)

type Container struct {
	AppConfig       config.AppConfig
	ConsumerUsecase consumer.ConsumerUsecase
	moduleRegistry  moduleregistry.ModuleRegistry
}

func NewMysqlDb(cfg config.AppConfig) *gorm.DB {
	db, err := database.NewMysql(cfg.DatabaseMysqlConnectionString, gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewMinioClient(cfg config.AppConfig) minio.Minio {
	client, err := minio.NewMinio(cfg.MinioConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func newContainer(
	moduleRegistry moduleregistry.ModuleRegistry,
	AppConfig config.AppConfig,
	ConsumerUsecase consumer.ConsumerUsecase,
) Container {
	c := Container{
		moduleRegistry:  moduleRegistry,
		AppConfig:       AppConfig,
		ConsumerUsecase: ConsumerUsecase,
	}
	return c
}

func InitializeDependencyContainer() Container {
	wire.Build(
		moduleregistry.NewModuleRegistry,
		config.LoadByEnv,
		NewMysqlDb,
		NewMinioClient,
		validator.New,
		mysqlRepo.NewConsumerRepository,
		filestorage.NewFileBucketClient,
		consumer.NewConsumerUsecase,
		newContainer,
	)
	return Container{}
}
