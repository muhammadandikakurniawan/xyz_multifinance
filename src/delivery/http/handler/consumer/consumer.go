package consumer

import (
	"net/http"

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
	// swagger:operation POST /consumer/tenor-limit consumer addTenorLimit
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
