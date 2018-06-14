package sender

import (
	"encoding/json"
	"fmt"
	"message"

	"github.com/GoogleChrome/push-encryption-go/webpush"
)

// GCMSender struct
type GCMSender struct{}

const (
	gcmKey string = "YOUR_ANDROID_SENDER_ID_KEY"
)

// SendWebPush Because we don't use the VAPID...
func (GCMSender) SendWebPush(m message.Message) error {
	fmt.Println("Send using legacy sender")
	message, mErr := m.MarshalPayload()

	if mErr != nil {
		return fmt.Errorf("Unable to send push, Message is badly format")
	}

	json, err := json.Marshal(m.UserInfo)

	if err != nil {
		return fmt.Errorf("Unable to convert struct to JSON %v", err)
	}

	fmt.Println(json)

	sub, sErr := webpush.SubscriptionFromJSON(json)

	if sErr != nil {
		return fmt.Errorf("Unable to get the subscription %v", err)
	}

	str := string(message[:])
	res, sendErr := webpush.Send(nil, sub, str, gcmKey)

	if sendErr != nil {
		return fmt.Errorf("Error while sending push notification %v", sendErr)
	} else {
		fmt.Println("Send with code ", res.StatusCode)
	}

	return nil
}
