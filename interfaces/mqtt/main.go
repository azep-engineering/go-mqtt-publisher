package mqtt

import (
	"context"

	"github.com/shkshariq/go-util/log"
	"github.com/shkshariq/go-util/mqtt"
)

var Handler *mqtt.MqttEventHandler

func Init() {
	Handler = mqtt.NewMqttHandler("")
}

type EventNotification struct {
	EventType   mqtt.EventType `json:"type,omitempty"`
	EventExpiry bool           `json:"expiry,omitempty"`
	EventHead   mqtt.Header    `json:"header,omitempty"`
	EventBody   interface{}    `json:"body,omitempty"`
}

func (e *EventNotification) Type() mqtt.EventType {
	return e.EventType
}

func (e *EventNotification) Expired() bool {
	return e.EventExpiry
}

func (e *EventNotification) Header() mqtt.Header {
	return e.EventHead
}

func (e *EventNotification) Body() interface{} {
	return e.EventBody
}

func FireEvent(path string, ev mqtt.Event, qos mqtt.Qos, retained bool) (err error) {
	log.DebugContext(context.Background(), "FireEvent:", ev, path, qos, retained)
	err = Handler.Publish(ev, path, qos, retained)
	return err
}
