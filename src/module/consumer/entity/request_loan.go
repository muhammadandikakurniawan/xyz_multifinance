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

type RequestLoanApprovalStatus string

var (
	RequestLoanApprovalStatus_APPROVED RequestLoanApprovalStatus = "APPROVED"
	RequestLoanApprovalStatus_REJECTED RequestLoanApprovalStatus = "REJECTED"
	RequestLoanApprovalStatus_PENDING  RequestLoanApprovalStatus = "PENDING"
)

type SearchRequestLoanFilterModel struct {
	ConsumerId     string                    `json:"consumer_id"`
	ConsumerName   string                    `json:"consumer_name"`
	ContractNumber string                    `json:"contract_number"`
	AssetName      string                    `json:"asset_name"`
	ApprovalStatus RequestLoanApprovalStatus `json:"approval_status"`
}
