package consumer

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/deliveryutil"
)

func NewConsumerHandler(consumerUsecase consumer.ConsumerUsecase) ConsumerHandler {
	return ConsumerHandler{
		consumerUsecase: consumerUsecase,
	}
}

type ConsumerHandler struct {
	consumerUsecase consumer.ConsumerUsecase
}

func (h ConsumerHandler) Register(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /consumer/register consumer registerConsumer
	// Register consumer
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/RequestCreateNewConsumerDto"
	// responses:
	//   '200':
	//     description: register response

	var requestBody dto.RequestCreateNewConsumerDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		return
	}

	ctx := r.Context()
	result, err := h.consumerUsecase.Register(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ConsumerHandler) RequestLoan(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /consumer/request-loan consumer requestLoan
	// request loan
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/RequestLoanDto"
	// responses:
	//   '200':
	//     description: request loan

	var requestBody dto.RequestLoanDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		return
	}

	ctx := r.Context()
	result, err := h.consumerUsecase.RequestLoan(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ConsumerHandler) AddTenorLimit(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /consumer/tenor-limit admin addTenorLimit
	// Add tenor limit for consumer
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/AddTenorLmitRequestDto"
	// responses:
	//   '200':
	//     description: add consumer tenor limit

	var requestBody dto.AddTenorLmitRequestDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		return
	}

	ctx := r.Context()
	result, err := h.consumerUsecase.AddTenorLimit(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ConsumerHandler) SearchRequestLoan(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /consumer/search-request-loan consumer searchRequestLoan
	// Add tenor limit for consumer
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: current_page
	//     in: query
	//   - name: total_data
	//     in: query
	//   - name: consumer_id
	//     in: query
	//   - name: asset_name
	//     in: query
	//   - name: contract_number
	//     in: query
	//   - name: consumer_name
	//     in: query
	//   - name: approval_status
	//     in: query
	// responses:
	//   '200':
	//     description: get list consumer's request loan

	filter := dto.GetListRequestLoanRequestDto{}
	filter.ReadRequest(r)
	ctx := r.Context()
	result, err := h.consumerUsecase.GetListRequestLoan(ctx, filter)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ConsumerHandler) ApproveRequestLoan(w http.ResponseWriter, r *http.Request) {

	// swagger:operation PUT /consumer/approve-request-loan admin approveRequestLoan
	// Add tenor limit for consumer
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     schema:
	//       "$ref": "#/definitions/ApprovalResponseDataDto"
	// responses:
	//   '200':
	//     description: set is_approve request loan

	var requestBody dto.ApprovalResponseDataDto
	if err := deliveryutil.ReadRequestBody(w, r, &requestBody); err != nil {
		return
	}

	ctx := r.Context()
	result, err := h.consumerUsecase.ApproveRequestLoan(ctx, requestBody)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}

func (h ConsumerHandler) GetDetailConsumer(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /consumer/{id} consumer admin GetConsumer
	// get consumer data
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: path
	//   description: Consumer ID
	//   type: string
	//   required: true
	// responses:
	//   '200':
	//     description: get consumer data

	consumerId := mux.Vars(r)["id"]

	ctx := r.Context()
	result, err := h.consumerUsecase.GetConsumer(ctx, consumerId)
	if err != nil {
		deliveryutil.ResponseErrorJson(w, r, err)
		return
	}

	deliveryutil.ResponseJson(w, r, result)
}
