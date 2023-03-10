package minio

import (
	"time"
)

type MinioConfig struct {
	Timeout      time.Duration `json:"timeout"`
	Host         string        `json:"host"`
	KeyId        string        `json:"keyId"`
	AccessKey    string        `json:"accessKey"`
	Ssl          bool          `json:"ssl"`
	ExternalHost string        `json:"external_host"`
}
