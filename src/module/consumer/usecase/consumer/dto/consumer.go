package dto

import (
	"strings"
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
)

type ConsumerId struct {
	Id string `json:"id" validate:"required"`
}

type BaseConsumerDto struct {
	ConsumerId
	NIK          string  `json:"nik" validate:"required"`
	Fullname     string  `json:"fullname" validate:"required"`
	Legalname    string  `json:"legalname" validate:"required"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
}

// swagger:model
type RequestCreateNewConsumerDto struct {
	NIK            string  `json:"nik" validate:"required"`
	Fullname       string  `json:"fullname" validate:"required"`
	Legalname      string  `json:"legalname" validate:"required"`
	PlaceOfBirth   string  `json:"place_of_birth" validate:"required"`
	DateOfBirth    string  `json:"date_of_birth" validate:"required"`
	Salary         float64 `json:"salary" validate:"required"`
	KtpImageBase64 string  `json:"ktp_base64" validate:"required"`
	SelfeBase64    string  `json:"selfie_base64" validate:"required"`
}

func (r RequestCreateNewConsumerDto) TransformToEntity(ktpImgUrl, selfieImgUrl string) (e entity.ConsumerEntity, err error) {
	dobFormat := "02/01/2006"
	dob, err := time.Parse(dobFormat, r.DateOfBirth)
	if err != nil {
		err = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "invalid date of birth", "invalid date of birth")
		return
	}

	ktpImgUrl = strings.ReplaceAll(ktpImgUrl, " ", "")
	if ktpImgUrl == "" {
		err = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "invalid ktp", "invalid ktp")
		return
	}

	selfieImgUrl = strings.ReplaceAll(selfieImgUrl, " ", "")
	if selfieImgUrl == "" {
		err = sharedErr.NewAppError(sharedErr.ERROR_BAD_REQUEST, "invalid selfie image", "invalid selfie image")
		return
	}

	e = entity.ConsumerEntity{
		NIK:          r.NIK,
		Fullname:     r.Fullname,
		Legalname:    r.Legalname,
		PlaceOfBirth: r.PlaceOfBirth,
		DateOfBirth:  dob.UnixMilli(),
		KtpImageUrl:  ktpImgUrl,
		SelfieUrl:    selfieImgUrl,
		Salary:       r.Salary,
	}
	return
}

type ConsumerDto struct {
	BaseConsumerDto
	KtpImageUrl string `json:"ktp" validate:"required"`
	SelfeUrl    string `json:"selfie" validate:"required"`
}
