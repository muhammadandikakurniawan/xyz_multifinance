package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(config MinioConfig) (res Minio, err error) {
	client, err := minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.KeyId, config.AccessKey, ""),
		Secure: config.Ssl,
	})
	if err != nil {
		return
	}
	res = &minioImpl{client, config}
	return
}

type Minio interface {
	Read(ctx context.Context, bucketName, fileName string) (io.Reader, error)
	Upload(ctx context.Context, path, bucketName, filename string) (*minio.UploadInfo, error)
	CreateBucket(ctx context.Context, bucketName string) (err error)
	IsBucketExists(ctx context.Context, bucketName string) (exists bool, err error)
	UploadBytes(ctx context.Context, bucketName, filename string, bytesData []byte, contentType string) (*minio.UploadInfo, error)
	GetHost() (host string)
	GetExternalHost() (host string)
}

type minioImpl struct {
	client *minio.Client
	config MinioConfig
}

func (o *minioImpl) GetHost() (host string) {
	host = o.client.EndpointURL().Scheme + "://" + o.client.EndpointURL().Host
	return
}

func (o *minioImpl) GetExternalHost() (host string) {
	return o.config.ExternalHost
}

func (o *minioImpl) Read(ctx context.Context, bucketName, fileName string) (io.Reader, error) {
	_, err := o.client.StatObject(ctx, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}

	object, err := o.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Println(fmt.Sprintf("cant read minio object : %v", err))
	}
	return object, nil
}

func (o *minioImpl) Upload(ctx context.Context, path, bucketName, filename string) (*minio.UploadInfo, error) {
	//check and create bucket
	if err := o.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := o.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Println("BUCKET EXISTS:", bucketName)
		} else {
			return nil, err
		}
	} else {
		log.Println("BUCKET CREATED:", bucketName)

		if err := o.client.SetBucketPolicy(context.Background(), bucketName, o.policy(bucketName)); err != nil {
			return nil, err
		}
	}

	//upload file to minio
	fileExt := strings.Split(filename, ".")
	info, err := o.client.FPutObject(
		ctx,
		bucketName,
		filename,
		path,
		minio.PutObjectOptions{ContentType: fileExt[len(fileExt)-1]},
	)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (o *minioImpl) CreateBucket(ctx context.Context, bucketName string) (err error) {
	exists, err := o.client.BucketExists(ctx, bucketName)
	if err != nil {
		return
	}
	if !exists {
		err = o.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return
		}
		err = o.client.SetBucketPolicy(context.Background(), bucketName, o.policy(bucketName))
	}

	return
}

func (o *minioImpl) IsBucketExists(ctx context.Context, bucketName string) (exists bool, err error) {
	exists, err = o.client.BucketExists(ctx, bucketName)
	return
}

func (o *minioImpl) UploadBytes(ctx context.Context, bucketName, filename string, bytesData []byte, contentType string) (*minio.UploadInfo, error) {
	//check and create bucket
	if err := o.CreateBucket(ctx, bucketName); err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	//upload file to minio
	info, err := o.client.PutObject(ctx, bucketName, filename, reader, int64(reader.Len()), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (o *minioImpl) policy(bucket string) string {
	if bucket == "" {
		return ""
	}

	return `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Action": [
					"s3:GetObject"
				],
				"Effect": "Allow",
				"Principal": {
					"AWS": ["*"]
				},
				"Resource": [
					"arn:aws:s3:::` + bucket + `/*"
				],
				"Sid": ""
			}
		]
	}`
}
