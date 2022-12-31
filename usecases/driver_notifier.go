package usecases

import (
	"context"
	"time"

	"azep/mqtt-publisher/domain"

	"github.com/shkshariq/go-util/log"
	"github.com/shkshariq/go-util/mqtt"
)

// Driver Notify Interactor
type NotificationInteractor struct {
}

func (interactor NotificationInteractor) GetTopic(ctx context.Context, req domain.Notifier) (topic domain.TopicInfo, err error) {
	topic.Topic = req.Topic.Topic
	topic.Qos = req.Topic.Qos
	topic.ExpiryDuration = req.Topic.ExpiryDuration
	topic.Retain = req.Topic.Retain

	return topic, err
}

func (interactor NotificationInteractor) GetHeader(ctx context.Context, req domain.Notifier, topic domain.TopicInfo) (header mqtt.Header) {
	header.Topic = req.Topic.Topic
	header.Version = domain.VERSION
	header.CreatedAt = req.Topic.CreatedAt
	if topic.ExpiryDuration != 0 {
		header.Expiry = GetExpiryTime(req.Topic.CreatedAt, topic.ExpiryDuration)
	}
	header.MessageID = GetNanoTime()
	return header
}

func GetExpiryTime(createdAt int64, duration int) (expiry int64) {
	cTime := time.Unix(createdAt, 0)
	eTime := cTime.Local().Add(time.Second * time.Duration(duration))
	log.InfoContext(context.Background(), "cTime.Unix():", cTime.Unix())
	log.InfoContext(context.Background(), "eTime.Unix():", eTime.Unix())

	return eTime.Unix()
}

func GetNanoTime() (nano int64) {
	time := time.Now()
	nano = time.UnixNano()
	return nano
}
