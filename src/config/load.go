package config

import (
	consumerConfig "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/utility"
	"github.com/spf13/cast"
)

func LoadConfig(
	consumerConfig consumerConfig.AppConfig,
) AppConfig {
	return AppConfig{
		HttpPort:             cast.ToInt(utility.GetRequiredEnv("HTTP_PORT")),
		ConsumerModuleConfig: consumerConfig,
	}
}
