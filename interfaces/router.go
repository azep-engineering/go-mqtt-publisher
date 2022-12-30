package interfaces

import (
	"fmt"
	"net/http"
	"os"

	"azep/mqtt-publisher/interfaces/endpoints"
	httpRequest "azep/mqtt-publisher/interfaces/transport/http/request"
	httpResponse "azep/mqtt-publisher/interfaces/transport/http/response"
	"azep/mqtt-publisher/services"

	kitLogger "github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/shkshariq/go-util/config"
	"github.com/shkshariq/go-util/response"
)

func Run() {

	r := mux.NewRouter()

	logger := kitLogger.NewLogfmtLogger(os.Stderr)
	opts := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(response.HandleError),
	}

	var notifySvc services.MQTTNotificationService
	{
		notifySvc = services.MQTTNotifier{}
	}

	r.Handle(`/v1/publish`, httpTransport.NewServer(
		endpoints.MakeEndpointMQTTNotification(notifySvc),
		httpRequest.DecodeRequestMQTTNotification,
		httpResponse.EncodeResponseMQTTNotification,
		opts...,
	)).Methods(http.MethodPost)

	port := config.AppConf.Port
	err := http.ListenAndServe(fmt.Sprintf(`:%d`, port), r)
	if err != nil {
		log.Fatalln(`Cannot start web server : `, err)
	}
}
