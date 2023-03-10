package dto

import "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"

// swagger:model
type RequestLoanDto struct {
	Id             int64   `json:"id"`
	ConsumerId     string  `json:"consumer_id" validate:"required"`
	IsApproved     *bool   `json:"is_approved,omitempty"`
	ContractNumber string  `json:"contract_number"`
	OTR            float64 `json:"otr" validate:"required"`
	AdminFee       float64 `json:"admin_fee" validate:"required"`
	Installment    float64 `json:"installment" validate:"required"`
	Interest       float64 `json:"interest" validate:"required"`
	AssetName      string  `json:"asset_name" validate:"required"`
}

func (r RequestLoanDto) TransformToEntity() entity.RequestLoanEntity {
	return entity.RequestLoanEntity{
		Id:             r.Id,
		ConsumerId:     r.ConsumerId,
		IsApproved:     r.IsApproved,
		ContractNumber: r.ContractNumber,
		OTR:            r.OTR,
		AdminFee:       r.AdminFee,
		Installment:    r.Installment,
		Interest:       r.Interest,
		AssetName:      r.AssetName,
	}
}
