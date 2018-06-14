package message

import (
	"encoding/json"
	"fmt"

	messaging "firebase.google.com/go/messaging"
)

type Key struct {
	UserAuth   string `json:"auth"`
	UserPubKey string `json:"p256dh"`
}

type User struct {
	Token string `json:"endpoint"`
	Keys  Key    `json:"keys"`
}

type A4Sb struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	ID    int    `json:"id"`
}

type Message struct {
	Title    string `json:"title"`
	Message  string `json:"body"`
	Icon     string `json:"icon"`
	Button   []A4Sb `json:"a4sb"`
	Pictures string `json:"a4sbigpicture"`
	UserInfo User   `json:"-"`
}

// MarshalPayload convert the Message struct to a JSON
func (m Message) MarshalPayload() ([]byte, error) {
	json, err := json.Marshal(m)

	if err != nil {
		return nil, fmt.Errorf("An error occured while marshaling the Message struct %v", err)
	}

	return json, nil
}

// PreparePayload for fcm
func (m Message) PreparePayload(tokens []string) []*messaging.Message {
	length := len(tokens)

	if length == 0 {
		fmt.Println("Tokens array is empty")
		return nil
	}

	// use slice to create the array at runtime with the desired length
	messages := make([]*messaging.Message, length)
	for i := 0; i < length; i++ {
		messages = append(messages, &messaging.Message{
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
