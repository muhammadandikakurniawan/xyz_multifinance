package dto

type ApprovalRequestaDto struct {
	IsApproved     bool `json:"is_approved" validate:"required"`
	ContractNumber bool `json:"contract_number" validate:"required"`
}

// swagger:model
type ApprovalResponseDataDto struct {
	Id         int64 `json:"id" validate:"required"`
	IsApproved bool  `json:"is_approved" validate:"required"`
}
