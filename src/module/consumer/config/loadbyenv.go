package config

import (
	"os"
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/minio"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/utility"
	"github.com/spf13/cast"
)

func LoadByEnv() AppConfig {
	return AppConfig{
		DatabaseMysqlConnectionString: utility.GetRequiredEnv("MODULE_CONSUMER_DB_MYSQL_STRCONNECTION"),
		CryptoAesKey:                  utility.GetRequiredEnv("CRYPTO_AES_KEY"),
		CryptoAesIV:                   utility.GetRequiredEnv("CRYPTO_AES_IV"),
		MinioConfig: minio.MinioConfig{
			Timeout:      time.Duration(cast.ToFloat64(os.Getenv("MINIO_TIMEOUT"))),
			Host:         utility.GetRequiredEnv("MINIO_HOST"),
			KeyId:        utility.GetRequiredEnv("MINIO_KEY_ID"),
			AccessKey:    utility.GetRequiredEnv("MINIO_ACESS_KEY"),
			Ssl:          cast.ToBool(utility.GetEnv("MINIO_SSL", "false")),
			ExternalHost: utility.GetRequiredEnv("MINIO_EXTERNAL_HOST"),
		},
		BucketConsumerImage: utility.GetRequiredEnv("BUCKET_CONSUMER_IMAGE"),
	}
}
