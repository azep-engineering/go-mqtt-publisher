package request

import (
	"azep/mqtt-publisher/domain"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shkshariq/go-util/log"
	"github.com/shkshariq/go-util/mqtt"
)

type Req struct {
	Topic          string      `json:"topic"`
	CreatedAt      int64       `json:"created_at"`
	Qos            int         `json:"qos"`
	Retain         bool        `json:"retain"`
	Expiry         int         `json:"expiry"`
	PublishMessage interface{} `json:"publish_message"`
}

func DecodeRequestMQTTNotification(ctx context.Context, req *http.Request) (request interface{}, err error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.DebugContext(context.Background(), "Error in reading req body:", err)
		return request, err
	}

	var clientRequest Req
	err = json.Unmarshal(body, &clientRequest)
	if err != nil {
		log.DebugContext(context.Background(), "Error in Unmarshaling clientRequest:", err)
		return request, err
	}

	var makeReq domain.Notifier
	makeReq.Topic.Topic = clientRequest.Topic
	makeReq.Topic.Qos = mqtt.Qos(clientRequest.Qos)
	makeReq.Topic.Retain = clientRequest.Retain
	makeReq.Topic.ExpiryDuration = clientRequest.Expiry
	makeReq.Topic.CreatedAt = clientRequest.CreatedAt
	makeReq.Payload = clientRequest.PublishMessage

	return makeReq, err
}
