package message

import (
	"fmt"

	messaging "firebase.google.com/go/messaging"
)

type Message struct {
	Title   string
	Message string
	Icon    string
}

func (m Message) PreparePayload(tokens []string) []messaging.Message {
	length := len(tokens)

	if length == 0 {
		fmt.Println("Tokens array is empty")
		return nil
	}

	// use slice to create the array at runtime with the desired length
	messages := make([]messaging.Message, length)
	for i := 0; i < length; i++ {
		messages = append(messages, messaging.Message{
			Webpush: &messaging.WebpushConfig{
				Notification: &messaging.WebpushNotification{
					Title: m.Title,
					Body:  m.Message,
					Icon:  m.Icon,
				},
			},
			Token: tokens[i],
		})

		fmt.Println(messages)
	}

	return messages
}
