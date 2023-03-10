package config

import (
	consumerConfig "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/config"
)

type AppConfig struct {
	HttpPort             int
	ConsumerModuleConfig consumerConfig.AppConfig
}
