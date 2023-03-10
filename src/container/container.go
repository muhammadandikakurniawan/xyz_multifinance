package container

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/config"
	consumerModuleContainer "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/moduleregistry"
)

type Container struct {
	modulerRegistry         moduleregistry.ModuleRegistry
	AppConfig               config.AppConfig
	ConsumerModuleContainer consumerModuleContainer.Container
}

func SetupContainer() Container {
	modulerRegistry := moduleregistry.NewModuleRegistry()
	consumerModuleContainer := consumerModuleContainer.SetupContainer(modulerRegistry)
	return Container{
		modulerRegistry:         modulerRegistry,
		AppConfig:               config.LoadConfig(consumerModuleContainer.AppConfig),
		ConsumerModuleContainer: consumerModuleContainer,
	}
}
