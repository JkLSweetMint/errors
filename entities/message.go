package entities

import (
	"fmt"
	"strings"
)

// Message - сообщение.
type Message struct {
	content string
}

// MessageOption - опции сообщения.
type MessageOption struct {
	Key   string
	Value any
}

// Set - установить значение сообщения.
func (message *Message) Set(content string, options ...MessageOption) *Message {
	message.content = content
	message.format(options)

	return message
}

// String - получение строкового представления сообщения.
func (message *Message) String(options ...MessageOption) (val string) {
	message.format(options)
	val = message.content

	return
}

// useOptions - применение опций форматирования к тексту.
func (message *Message) format(options []MessageOption) {
	if options == nil {
		return
	}

	for _, opt := range options {
		message.content = strings.Replace(message.content, fmt.Sprintf("{{%s}}", opt.Key), fmt.Sprintf("%+v", opt.Value), -1)
	}
}
