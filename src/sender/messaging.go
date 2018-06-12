package sender

import (
	"context"
	"fmt"

	messaging "firebase.google.com/go/messaging"
)

type Message struct {
	Title   string
	Message string
	Icon    string
}

func (m Message) preparePayload(tokens []string) {
	length := len(tokens)

	if length == 0 {
		fmt.Println("Tokens array is empty")
		return
	}

	// use slice to create the array at runtime with the desired length
	messages := make([]messaging.Message, length)

	for i := 0; i <= length; i++ {
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
	}
}

func (f FirebaseApp) send(messages []messaging.Message) error {
	ctx := context.Background()
	client, err := f.App.Messaging(ctx)

	if err != nil {
		return fmt.Errorf("An error happened while retrieving the clients: %s", err.Error())
	}

	for i := 0; i <= len(messages); i++ {
		response, err := client.Send(ctx, &messages[i])

		if err != nil {
			fmt.Println("Error occurred while sending message: %s", err.Error)
			fmt.Println(response)
		}
	}

	fmt.Println("Push was send")
	return nil
}
