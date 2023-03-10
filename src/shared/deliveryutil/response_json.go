package deliveryutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedError "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
)

func Log(data interface{}, r *http.Request) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Printf("response %s : %s\n", r.URL.Path, string(dataBytes))
}

func ResponseJson[T any](w http.ResponseWriter, r *http.Request, model model.BaseResponseModel[T]) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(model.HttpStatusCode)
	Log(model, r)
	json.NewEncoder(w).Encode(model)
}

func ResponseErrorJson(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	model := model.BaseResponseModel[bool]{
		ErrorMessage:   err.Error(),
		Message:        "internal server error",
		Success:        false,
		HttpStatusCode: sharedError.INTERNAL_SERVER_ERROR.ToHttpStatus(),
		StatusCode:     string(sharedError.INTERNAL_SERVER_ERROR),
	}
	Log(model, r)
	json.NewEncoder(w).Encode(model)
}

func ReadRequestBody(w http.ResponseWriter, r *http.Request, dst interface{}) (err error) {
	err = json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		ResponseErrorJson(w, r, err)
		return
	}
	return
}
