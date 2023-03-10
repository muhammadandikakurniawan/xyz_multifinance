package entity

import "time"

type TenorLimitEntity struct {
	ConsumerId string     `json:"consumer_id"`
	Month      int        `json:"month"`
	LimitValue float64    `json:"limit_value"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
