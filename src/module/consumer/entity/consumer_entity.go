package entity

import (
	"strings"
	"time"

	sharedError "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
)

type ConsumerEntity struct {
	Id           string     `json:"id"`
	NIK          string     `json:"nik"`
	Fullname     string     `json:"full_name"`
	Legalname    string     `json:"legal_name"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfBirth  int64      `json:"date_of_birth"`
	Salary       float64    `json:"salary"`
	KtpImageUrl  string     `json:"ktp"`
	SelfieUrl    string     `json:"selfie"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`

	TenorLimits          []TenorLimitEntity        `json:"tenor_limits"`
	MapTenorLimitByMonth map[int]*TenorLimitEntity `json:"-"`

	ListRequestLoan    []RequestLoanEntity          `json:"list_request_loan"`
	MapRequestLoanById map[int64]*RequestLoanEntity `json:"-"`
}

func (e *ConsumerEntity) AddTenorLimit(tl TenorLimitEntity) {
	if e.MapTenorLimitByMonth == nil {
		e.MapTenorLimitByMonth = map[int]*TenorLimitEntity{}
	}

	e.TenorLimits = append(e.TenorLimits, tl)
	e.MapTenorLimitByMonth[tl.Month] = &e.TenorLimits[len(e.TenorLimits)-1]
}

func (e *ConsumerEntity) AddRequestLoan(req RequestLoanEntity) {
	if e.MapRequestLoanById == nil {
		e.MapRequestLoanById = map[int64]*RequestLoanEntity{}
	}

	e.ListRequestLoan = append(e.ListRequestLoan, req)
	e.MapRequestLoanById[req.Id] = &e.ListRequestLoan[len(e.ListRequestLoan)-1]
}

func (e ConsumerEntity) ValidateNewConsumer() (err error) {
	errorValidationMessages := []string{}

	e.NIK = strings.ReplaceAll(e.NIK, " ", "")
	if e.NIK == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid nik")
	}

	e.Fullname = strings.ReplaceAll(e.Fullname, " ", "")
	if e.NIK == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid full name")
	}

	e.Legalname = strings.ReplaceAll(e.Legalname, " ", "")
	if e.NIK == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid legal name")
	}

	e.PlaceOfBirth = strings.ReplaceAll(e.PlaceOfBirth, " ", "")
	if e.NIK == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid place of birth")
	}

	e.KtpImageUrl = strings.ReplaceAll(e.KtpImageUrl, " ", "")
	if e.KtpImageUrl == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid ktp")
	}

	e.SelfieUrl = strings.ReplaceAll(e.SelfieUrl, " ", "")
	if e.SelfieUrl == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid selfie")
	}

	if e.DateOfBirth <= 0 {
		errorValidationMessages = append(errorValidationMessages, "invalid date of birth")
	}

	invalidData := len(errorValidationMessages) > 0
	if invalidData {
		err = sharedError.NewValidationError(strings.Join(errorValidationMessages, ", "))
	}
	return
}
