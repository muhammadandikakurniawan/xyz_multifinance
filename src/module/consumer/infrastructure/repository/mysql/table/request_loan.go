package table

import (
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
)

func SetupRequestLoanTable(e entity.RequestLoanEntity) RequestLoanTable {
	tbl := RequestLoanTable{
		Id:             e.Id,
		ConsumerId:     e.ConsumerId,
		IsApproved:     e.IsApproved,
		ContractNumber: e.ContractNumber,
		OTR:            e.OTR,
		AdminFee:       e.AdminFee,
		Installment:    e.Installment,
		Interest:       e.Interest,
		AssetName:      e.AssetName,
		CreatedAt:      e.CreatedAt.UnixMilli(),
	}
	if e.UpdatedAt != nil {
		tbl.UpdatedAt = e.UpdatedAt.UnixMilli()
	}
	if e.DeletedAt != nil {
		tbl.DeletedAt = e.DeletedAt.UnixMilli()
	}
	return tbl
}

type RequestLoanTable struct {
	Id             int64   `gorm:"column:id"`
	ConsumerId     string  `gorm:"column:consumer_id"`
	IsApproved     *bool   `gorm:"column:is_approved"`
	ContractNumber string  `gorm:"column:contract_number"`
	OTR            float64 `gorm:"column:otr"`
	AdminFee       float64 `gorm:"column:admin_fee"`
	Installment    float64 `gorm:"column:installment"`
	Interest       float64 `gorm:"column:interest"`
	AssetName      string  `gorm:"column:asset_name"`
	CreatedAt      int64   `gorm:"column:created_at"`
	UpdatedAt      int64   `gorm:"column:updated_at"`
	DeletedAt      int64   `gorm:"column:deleted_at"`
}

func (*RequestLoanTable) TableName() string {
	return "request_loan"
}

func (tbl RequestLoanTable) TransformToEntity() entity.RequestLoanEntity {
	e := entity.RequestLoanEntity{
		Id:             tbl.Id,
		ConsumerId:     tbl.ConsumerId,
		IsApproved:     tbl.IsApproved,
		ContractNumber: tbl.ContractNumber,
		OTR:            tbl.OTR,
		AdminFee:       tbl.AdminFee,
		Installment:    tbl.Installment,
		Interest:       tbl.Interest,
		AssetName:      tbl.AssetName,
		CreatedAt:      time.UnixMilli(tbl.CreatedAt),
	}
	if tbl.UpdatedAt > 0 {
		t := time.UnixMilli(tbl.UpdatedAt)
		e.UpdatedAt = &t
	}
	if tbl.DeletedAt > 0 {
		t := time.UnixMilli(tbl.DeletedAt)
		e.DeletedAt = &t
	}
	return e
}
