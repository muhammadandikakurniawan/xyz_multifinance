package table

import (
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
)

func SetupTenorLimitTable(e entity.TenorLimitEntity) TenorLimitTable {
	tbl := TenorLimitTable{
		ConsumerId: e.ConsumerId,
		Month:      e.Month,
		LimitValue: e.LimitValue,
		CreatedAt:  e.CreatedAt.UnixMilli(),
	}
	if e.UpdatedAt != nil {
		tbl.UpdatedAt = e.UpdatedAt.UnixMilli()
	}
	if e.DeletedAt != nil {
		tbl.DeletedAt = e.DeletedAt.UnixMilli()
	}
	return tbl
}

type TenorLimitTable struct {
	ConsumerId string  `gorm:"column:consumer_id"`
	Month      int     `gorm:"column:month"`
	LimitValue float64 `gorm:"column:limit_value"`
	CreatedAt  int64   `gorm:"column:created_at"`
	UpdatedAt  int64   `gorm:"column:updated_at"`
	DeletedAt  int64   `gorm:"column:deleted_at"`
}

func (*TenorLimitTable) TableName() string {
	return "tenor_limit"
}

func (tbl TenorLimitTable) TransformToEntity() entity.TenorLimitEntity {
	e := entity.TenorLimitEntity{
		ConsumerId: tbl.ConsumerId,
		Month:      tbl.Month,
		LimitValue: tbl.LimitValue,
		CreatedAt:  time.UnixMilli(tbl.CreatedAt),
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
