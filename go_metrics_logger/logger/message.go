package logger

import "time"

type Status string

const (
	StatusSuccess Status = "Success"
	StatusError   Status = "Error"
)

type Message struct {
	EventID       int64         `json:"EVENT_ID"`
	OperationName string        `json:"OPERATION_NAME"`
	ComponentName string        `json:"COMPONENT_NAME"`
	EventStatus   Status        `json:"EVENT_STATUS"`
	Description   string        `json:"DESCRIPTION"`
	TS            time.Time     `json:"TS"`
	Latency       time.Duration `json:"LATENCY"`
	Headers       string        `json:"Headers"`
	Payload       string        `json:"Payload"`
}
