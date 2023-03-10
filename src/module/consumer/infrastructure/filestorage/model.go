package filestorage

type UploadFileOpt struct {
	Base64String string
	Filename     string
	Bucket       string
	ContentType  string
}

type UploadResultOpt struct {
	FilePath string
}
