package filestorage

import "context"

type FileStorage interface {
	UploadBase64(ctx context.Context, opt UploadFileOpt) (result UploadResultOpt, err error)
}
