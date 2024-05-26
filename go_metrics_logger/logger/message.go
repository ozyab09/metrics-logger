package logger

import "time"

// Status - status of the operation transmitted in the message
type Status string

const (
	// StatusSuccess - success status
	StatusSuccess Status = "Success"
	// StatusError - error status
	StatusError Status = "Error"
)

// Message - structure of messages that are sent for logging

type I_Message interface {
	GetOperationName() string
	GetComponentName() string
	GetEventStatus() Status
	GetLatency() time.Duration
}

/*type Message struct {
	// EventID - id of event
	EventID int64 `json:"EVENT_ID"`
	// OperationName - name of logging operation
	OperationName string `json:"OPERATION_NAME"`
	// ComponentName - name of component name
	ComponentName string `json:"COMPONENT_NAME"`
	// EventStatus - status of message event
	EventStatus Status `json:"EVENT_STATUS"`
	// Description - description of event
	Description string `json:"DESCRIPTION"`
	// TS - event timestamp
	TS time.Time `json:"TS"`
	// Latency - latency of operation
	Latency time.Duration `json:"LATENCY"`
	// Headers - event headers
	Headers string `json:"Headers"`
	// Payload - additional message payload
	Payload string `json:"Payload"`
}
*/
