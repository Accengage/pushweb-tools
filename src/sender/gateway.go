package sender

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	App *firebase.App
}

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

// Init the firebase App
func (f FirebaseApp) Init() {
	app, err := prepareGateway()

	if err != nil {
		fmt.Printf("FCM App stop here: %s", err.Error())
		// stop the app right away
		return
	}

	f.App = app
}
