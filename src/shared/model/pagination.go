package model

import (
	"net/http"

	"github.com/spf13/cast"
)

type PaginationResponseModel struct {
	TotalData   uint64 `json:"total_data,omitempty"`
	TotalPage   uint64 `json:"total_page,omitempty"`
	CurrentPage uint64 `json:"current_page,omitempty"`
}

// swagger:parameters listBars addBars
type PaginationRequestModel struct {
	TotalData uint64 `json:"total_data"`
	Page      uint64 `json:"current_page"`
}

func ReadPaginationRequestFromHttp(r *http.Request) (res PaginationRequestModel) {
	res.TotalData = cast.ToUint64(r.URL.Query().Get("total_data"))
	res.Page = cast.ToUint64(r.URL.Query().Get("page"))
	return
}
