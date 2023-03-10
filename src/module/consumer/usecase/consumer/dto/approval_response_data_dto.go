package dto

type ApprovalRequestaDto struct {
	IsApproved     bool `json:"is_approved" validate:"required"`
	ContractNumber bool `json:"contract_number" validate:"required"`
}

type ApprovalResponseDataDto struct {
	IsApproved bool `json:"is_approved" validate:"required"`
}
