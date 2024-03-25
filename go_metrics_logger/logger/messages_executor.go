package logger

import "context"

type (
	MessageExecutor interface {
		ExecuteMessage(level string, text string)
	}

	messageExecutorImpl struct {
		ctx        context.Context
		logger     Logger
		messages   chan logEvent
		workersNum uint32
	}

	logEvent struct {
		level string
		text  string
	}
)

func NewMessageExecutorImpl(ctx context.Context, logger Logger, workersNum uint32, bufferSize uint32) MessageExecutor {
	exec := &messageExecutorImpl{
		ctx:        ctx,
		logger:     logger,
		messages:   make(chan logEvent, bufferSize),
		workersNum: workersNum,
	}
	exec.init()

	return exec
}

func (i *messageExecutorImpl) init() {
	for k := 0; uint32(k) < i.workersNum; k++ {
		go func() {
			for {
				select {
				case msg := <-i.messages:
					i.logger.Log(msg.level, msg.text)
				case <-i.ctx.Done():
					return
				}
			}
		}()
	}
}

func (i *messageExecutorImpl) ExecuteMessage(level string, text string) {
	i.messages <- logEvent{level: level, text: text}
}
