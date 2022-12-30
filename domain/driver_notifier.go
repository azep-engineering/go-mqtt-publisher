package domain

import "github.com/shkshariq/go-util/mqtt"

const (
	VERSION = 1
)

const (
	TYPE_ORDER_CREATED       = "new_order"
	TYPE_ORDER_DELIVERED     = "order_delivered"
	TYPE_ORDER_COMPLETED     = "completed_order"
	TYPE_ORDER_CANCELLED     = "cancelled_order"
	TYPE_DRIVER_DISPATCHED   = "driver_arrived"
	TYPE_KITCHEN_PAUSE       = "kitchen_pause"
	TYPE_KITCHEN_UNPAUSE     = "kitchen_unpause"
	TYPE_DELIVERY_NOTE       = "delivery_note"
	TYPE_ORDER_SCHEDULED     = "shedule_order"
	TYPE_KITCHEN_TERMINATION = "kitchen_termination"
	TYPE_KITCHEN_AUDIT       = "kitchen_audit"
	TYPE_KITCHEN_SUSPENDED   = "kitchen_suspended"
	TYPE_KITCHEN_REACTIVATE  = "kitchen_reactivated"
	TYPE_KITCHEN_DN_ONGOING  = "kitchen_dn_ongoing"
)

type Notifier struct {
	Topic   TopicInfo   `json:"topic,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type TopicInfo struct {
	Topic          string   `json:"topic"`
	CreatedAt      int64    `json:"created_at"`
	Qos            mqtt.Qos `json:"qos"`
	ExpiryDuration int      `json:"expiry"` //in Seconds
	Retain         bool     `json:"retain"`
}

type Header struct {
	Type      string `json:"type"`
	Version   int    `json:"version"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int64  `json:"expiry,omitempty"`
}
