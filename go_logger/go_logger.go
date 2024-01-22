package go_logger

import (
	"time"

	"github.com/google/uuid"
)

type logger struct {
	id             string        `json:"ESBEVENT_ID,omitempty"`
	operationName  string        `json:"ESBEVENT_OPRATIONNAME,omitempty"`
	componentName  string        `json:"ESBEVENT_COMPONENTNAME,omitempty"`
	status         string        `json:"ESBEVENT_STATUS,omitempty"`
	description    string        `json:"ESBEVENT_DESCRIPTION,omitempty"`
	eventStartTime time.Time     `json:"ESBEVENT_EVNTTS,omitempty"`
	timeDiff       time.Duration `json:"ESBEVENT_TIMEDIFF,omitempty"`
	comment        string        `json:"ESBEVENT_COMMENT,omitempty"`
}

func Init() (log *logger) {
	log = &logger{
		id:             uuid.NewString(),
		eventStartTime: time.Now(),
	}
	return
}

func (this *logger) SetId(id string) {
	if id != "" {
		this.id = id
	}
}

func (this *logger) SetOperationName(operationName string) {
	if operationName != "" {
		this.operationName = operationName
	}
}

func (this *logger) SetComponentName(componentName string) {
	if componentName != "" {
		this.componentName = componentName
	}
}

func (this *logger) SetStatus(status string) {
	if status != "" {
		this.status = status
	}
}

func (this *logger) SetDescription(description string) {
	if description != "" {
		this.description = description
	}
}

func (this *logger) SetComment(comment string) {
	if comment != "" {
		this.comment = comment
	}
}

func (this *logger) End() {
	this.timeDiff = time.Now().Sub(this.eventStartTime)
}
