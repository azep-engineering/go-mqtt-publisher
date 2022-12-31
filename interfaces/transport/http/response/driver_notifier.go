package response

import (
	"azep/mqtt-publisher/interfaces/endpoints"
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponseMQTTNotification(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := response.(endpoints.CreateDriverNotifyResponse)

	apiResponse := getResponse(false, "success", res.Response)

	return json.NewEncoder(w).Encode(apiResponse)
}
