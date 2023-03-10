package entity

import "time"

type RequestLoanEntity struct {
	Id             int64      `json:"id"`
	ConsumerId     string     `json:"consumer_id"`
	IsApproved     *bool      `json:"is_approved"`
	ContractNumber string     `json:"contract_number"`
	OTR            float64    `json:"otr"`
	AdminFee       float64    `json:"admin_fee"`
	Installment    float64    `json:"installment"`
	Interest       float64    `json:"interest"`
	AssetName      string     `json:"asset_name"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
