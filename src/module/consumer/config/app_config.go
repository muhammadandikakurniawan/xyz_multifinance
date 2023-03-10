package config

import "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/minio"

type AppConfig struct {
	DatabaseMysqlConnectionString string
	CryptoAesKey                  string
	CryptoAesIV                   string
	MinioConfig                   minio.MinioConfig
	BucketConsumerImage           string
}
