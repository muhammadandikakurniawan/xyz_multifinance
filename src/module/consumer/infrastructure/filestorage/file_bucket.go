package filestorage

import (
	"context"
	"encoding/base64"
	"log"
	"strings"

	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/minio"
)

func NewFileBucketClient(client minio.Minio) FileStorage {
	return &fileBucketClient{
		client: client,
	}
}

type fileBucketClient struct {
	client minio.Minio
}

func (f *fileBucketClient) UploadBase64(ctx context.Context, opt UploadFileOpt) (result UploadResultOpt, err error) {
	base64Str := strings.ReplaceAll(opt.Base64String, " ", "")
	if base64Str == "" {
		err = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "base64 cannot be empty", "base64 cannot be empty")
		return
	}

	fileBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return
	}

	uploadInfo, err := f.client.UploadBytes(ctx, opt.Bucket, opt.Filename, fileBytes, opt.ContentType)
	if err != nil {
		log.Println(err)
		return
	}
	result.FilePath = strings.Join([]string{f.client.GetExternalHost(), uploadInfo.Bucket, uploadInfo.Key}, "/")
	return
}

func (f *fileBucketClient) CreateDirectory(ctx context.Context, dirname string) (err error) {
	isExists, err := f.client.IsBucketExists(ctx, dirname)
	if err != nil {
		return
	}
	if isExists {
		return
	}
	return f.client.CreateBucket(ctx, dirname)
}
