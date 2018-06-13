package sender

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseApp struct{}

var (
	app *firebase.App
)

// Get the executable Path
func getExecPath() (string, error) {
	ex, err := os.Executable()

	if err != nil {
		return "", fmt.Errorf("error while retrieving the path of the executable")
	}

	exPath := filepath.Dir(ex)
	return exPath, nil
}

// Prepare Gateway return the firebase app
func prepareGateway() (*firebase.App, error) {
	path, pErr := getExecPath()

	if pErr != nil {
		return nil, pErr
	}

	opt := option.WithCredentialsFile(path + "/static/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, fmt.Errorf("error while initializing the project: %v", err)
	}

	return app, nil
}

// Send the message
func (f FirebaseApp) Send(messages []*messaging.Message) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)

	if err != nil {
		fmt.Println("an error happened", err.Error())
		return fmt.Errorf("An error happened while retrieving the clients: %s", err.Error())
	}

	for i := 1; i <= len(messages)-1; i++ {
		_, err := client.Send(ctx, messages[i])

		if err != nil {
			fmt.Println("Error occurred while sending message: ", err.Error())
		}
	}

	fmt.Println("Push was send")
	return nil
}

// Init the firebase App
func (f FirebaseApp) Init() {
	pAapp, err := prepareGateway()

	if err != nil {
		fmt.Printf("FCM App stop here: %s", err.Error())
		// stop the app right away
		return
	}

	app = pAapp
}
