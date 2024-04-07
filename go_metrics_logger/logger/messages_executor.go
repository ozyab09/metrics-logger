package logger

import "context"

type (
	// MessageExecutor - message executor interface
	MessageExecutor interface {
		// ExecuteMessage - process message
		ExecuteMessage(level string, text string)
		// Close - close executor
		Close() error
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

// NewMessageExecutorImpl - return new message executor
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

// ExecuteMessage - process message
func (i *messageExecutorImpl) ExecuteMessage(level string, text string) {
	i.messages <- logEvent{level: level, text: text}
}

// Close - close executor
func (i *messageExecutorImpl) Close() error {
	return i.logger.Close()
}
