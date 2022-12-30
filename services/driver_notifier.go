package services

import (
	"context"

	"azep/mqtt-publisher/domain"
	"azep/mqtt-publisher/interfaces/mqtt"
	"azep/mqtt-publisher/usecases"

	"github.com/shkshariq/go-util/logger"
)

type MQTTNotificationService interface {
	PublishMQTTNotification(ctx context.Context, req domain.Notifier) (res interface{}, err error)
}

type MQTTNotifier struct{}

func (MQTTNotifier) PublishMQTTNotification(ctx context.Context, req domain.Notifier) (res interface{}, err error) {

	//Driver Notification Interactor
	notifyInteractor := usecases.NotificationInteractor{}

	//get & set topic values
	topic, err := notifyInteractor.GetTopic(ctx, req)
	if err != nil {
		logger.Log().ErrorContext(ctx, err)
		return
	}

	header := notifyInteractor.GetHeader(ctx, req, topic)

	//Initialize event and set values
	event := new(mqtt.EventNotification)
	event.EventHead = header
	event.EventBody = req.Payload

	//Fire event payload
	err = mqtt.FireEvent(topic.Topic, event, topic.Qos, topic.Retain)
	if err != nil {
		logger.Log().ErrorContext(ctx, err)
		return
	}

	return event, err
}
