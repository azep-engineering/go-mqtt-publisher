package endpoints

import (
	"azep/mqtt-publisher/domain"
	"azep/mqtt-publisher/services"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type CreateDriverNotifyResponse struct {
	Request  domain.Notifier
	Response interface{}
}

func MakeEndpointMQTTNotification(svc services.MQTTNotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		requestStatus, _ := request.(domain.Notifier)
		responseStatus, err := svc.PublishMQTTNotification(ctx, requestStatus)

		return CreateDriverNotifyResponse{
			requestStatus,
			responseStatus,
		}, err
	}
}
