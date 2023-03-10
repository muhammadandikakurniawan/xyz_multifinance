package dto

import "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"

type TenorLimitDto struct {
	ConsumerId string  `json:"consumer_id" validate:"required"`
	Month      int     `json:"month" validate:"required"`
	Value      float64 `json:"value" validate:"required"`
}

func (d TenorLimitDto) TransformToEntity(consumerId string) entity.TenorLimitEntity {
	return entity.TenorLimitEntity{
		ConsumerId: consumerId,
		Month:      d.Month,
		LimitValue: d.Value,
	}
}

// swagger:model
type AddTenorLmitRequestDto struct {
	ConsumerId string          `json:"consumer_id" validate:"required"`
	Limits     []TenorLimitDto `json:"limits" validate:"required"`
}
