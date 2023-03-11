package dto

import (
	"net/http"
	"strings"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
	"github.com/spf13/cast"
)

type GetListRequestLoanRequestDto struct {
	Pagination model.PaginationRequestModel
	Filter     entity.SearchRequestLoanFilterModel
}

func (m *GetListRequestLoanRequestDto) ReadRequest(r *http.Request) {
	m.Pagination.TotalData = cast.ToUint64(r.URL.Query().Get("total_data"))
	m.Pagination.Page = cast.ToUint64(r.URL.Query().Get("current_page"))

	m.Filter.ApprovalStatus = entity.RequestLoanApprovalStatus(strings.ReplaceAll(r.URL.Query().Get("approval_status"), " ", ""))
	m.Filter.AssetName = strings.TrimSpace(r.URL.Query().Get("asset_name"))
	m.Filter.ConsumerId = strings.ReplaceAll(r.URL.Query().Get("consumer_id"), " ", "")
	m.Filter.ConsumerName = strings.TrimSpace(r.URL.Query().Get("consumer_name"))
	m.Filter.ContractNumber = strings.ReplaceAll(r.URL.Query().Get("contract_number"), " ", "")

}
