package dto

import (
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
)

// swagger:model
type RequestLoanDto struct {
	Id             int64     `json:"id"`
	ConsumerId     string    `json:"consumer_id" validate:"required"`
	ConsumerName   string    `json:"consumer_name,omitempty"`
	ConsumerSalary float64   `json:"consumer_salary,omitempty"`
	IsApproved     *bool     `json:"is_approved,omitempty"`
	ContractNumber string    `json:"contract_number"`
	OTR            float64   `json:"otr" validate:"required"`
	AdminFee       float64   `json:"admin_fee" validate:"required"`
	Installment    float64   `json:"installment" validate:"required"`
	Interest       float64   `json:"interest" validate:"required"`
	AssetName      string    `json:"asset_name" validate:"required"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
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
func TransformFromConsumerEntities(entities []entity.ConsumerEntity) (res []RequestLoanDto) {

	for _, cs := range entities {
		for _, req := range cs.ListRequestLoan {
			res = append(res, RequestLoanDto{
				Id:             req.Id,
				ConsumerId:     req.ConsumerId,
				ConsumerName:   cs.Legalname,
				ConsumerSalary: cs.Salary,
				IsApproved:     req.IsApproved,
				ContractNumber: req.ContractNumber,
				OTR:            req.OTR,
				AdminFee:       req.AdminFee,
				Installment:    req.Installment,
				Interest:       req.Interest,
				AssetName:      req.AssetName,
				CreatedAt:      req.CreatedAt,
			})
		}
	}

	return
}
