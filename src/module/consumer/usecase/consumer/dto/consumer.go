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
	// Unique identifier for this consumer, need ecrypted using aes cbc
	// min: 16
	// max: 20
	NIK          string `json:"nik" validate:"required"`
	Fullname     string `json:"fullname" validate:"required"`
	Legalname    string `json:"legalname" validate:"required"`
	PlaceOfBirth string `json:"place_of_birth" validate:"required"`
	// The birth date for this consumer, ex : dd/mm/yyyy (29/12/1999)
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
	selfieImgUrl = strings.ReplaceAll(selfieImgUrl, " ", "")
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
	KtpImageUrl    string          `json:"ktp" validate:"required"`
	SelfieUrl      string          `json:"selfie" validate:"required"`
	ListTenorLimit []TenorLimitDto `json:"list_tenor_limit,omitempty"`
}

func (m *ConsumerDto) SetupFromEntity(e entity.ConsumerEntity) {
	m.Id = e.Id
	m.Fullname = e.Fullname
	m.Legalname = e.Legalname
	m.NIK = e.NIK
	m.PlaceOfBirth = e.PlaceOfBirth
	m.DateOfBirth = time.UnixMilli(e.DateOfBirth).Format("02/01/2006")
	m.Salary = e.Salary
	m.KtpImageUrl = e.KtpImageUrl
	m.SelfieUrl = e.SelfieUrl

	for _, tl := range e.TenorLimits {
		m.ListTenorLimit = append(m.ListTenorLimit, TenorLimitDto{
			Month: tl.Month,
			Value: tl.LimitValue,
		})
	}
}
