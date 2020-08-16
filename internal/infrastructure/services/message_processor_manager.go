package application

import "fmt"

type MessageProcessorManager interface {
	Get(messageType string) (MessageProcessor, error)
}

type messageProcessorManager struct {
	processors []MessageProcessor
}

func (m *messageProcessorManager) Get(messageType string) (MessageProcessor, error) {
	for _, processor := range m.processors {
		if processor.shouldProcess(messageType) {
			return processor, nil
		}
	}

	return nil, fmt.Errorf("processor not found for type %s", messageType)
}

func NewMessageProcessorManager(processors ...MessageProcessor) MessageProcessorManager {
	return &messageProcessorManager{processors: processors}
}
