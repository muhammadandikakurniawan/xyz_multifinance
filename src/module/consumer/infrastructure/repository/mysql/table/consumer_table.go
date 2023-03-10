package table

import (
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"golang.org/x/sync/errgroup"
)

func SetupConsumerTable(e entity.ConsumerEntity) ConsumerTable {
	tbl := ConsumerTable{
		Id:           e.Id,
		NIK:          e.NIK,
		Fullname:     e.Fullname,
		Legalname:    e.Legalname,
		PlaceOfBirth: e.PlaceOfBirth,
		DateOfBirth:  e.DateOfBirth,
		Salary:       e.Salary,
		KtpImageUrl:  e.KtpImageUrl,
		SelfieUrl:    e.SelfieUrl,
		CreatedAt:    e.CreatedAt.UnixMilli(),
	}
	if e.UpdatedAt != nil {
		tbl.UpdatedAt = e.UpdatedAt.UnixMilli()
	}
	if e.DeletedAt != nil {
		tbl.DeletedAt = e.DeletedAt.UnixMilli()
	}
	return tbl
}

type ConsumerTable struct {
	Id              string             `gorm:"column:id"`
	NIK             string             `gorm:"column:nik"`
	Fullname        string             `gorm:"column:full_name"`
	Legalname       string             `gorm:"column:legal_name"`
	PlaceOfBirth    string             `gorm:"column:place_of_birth"`
	DateOfBirth     int64              `gorm:"column:date_of_birth"`
	Salary          float64            `gorm:"column:salary"`
	KtpImageUrl     string             `gorm:"column:ktp_image_url"`
	SelfieUrl       string             `gorm:"column:selfie_image_url"`
	CreatedAt       int64              `gorm:"column:created_at"`
	UpdatedAt       int64              `gorm:"column:updated_at"`
	DeletedAt       int64              `gorm:"column:deleted_at"`
	TenorLimits     []TenorLimitTable  `gorm:"foreignKey:consumer_id;references:id"`
	ListRequestLoan []RequestLoanTable `gorm:"foreignKey:consumer_id;references:id"`
}

func (*ConsumerTable) TableName() string {
	return "consumer"
}

func (tbl ConsumerTable) TransformToEntity() entity.ConsumerEntity {
	e := entity.ConsumerEntity{
		Id:           tbl.Id,
		NIK:          tbl.NIK,
		Fullname:     tbl.Fullname,
		Legalname:    tbl.Legalname,
		PlaceOfBirth: tbl.PlaceOfBirth,
		DateOfBirth:  tbl.DateOfBirth,
		Salary:       tbl.Salary,
		KtpImageUrl:  tbl.KtpImageUrl,
		SelfieUrl:    tbl.SelfieUrl,
		CreatedAt:    time.UnixMilli(tbl.CreatedAt),
	}
	if tbl.UpdatedAt > 0 {
		t := time.UnixMilli(tbl.UpdatedAt)
		e.UpdatedAt = &t
	}
	if tbl.DeletedAt > 0 {
		t := time.UnixMilli(tbl.DeletedAt)
		e.DeletedAt = &t
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		for _, tenorLimit := range tbl.TenorLimits {
			e.AddTenorLimit(tenorLimit.TransformToEntity())
		}
		return nil
	})

	eg.Go(func() error {
		for _, req := range tbl.ListRequestLoan {
			e.AddRequestLoan(req.TransformToEntity())
		}
		return nil
	})

	eg.Wait()

	return e
}
