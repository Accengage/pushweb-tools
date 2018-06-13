package gui

import (
	"fmt"
	"message"
	"sender"

	"github.com/andlabs/ui"
)

var (
	window *ui.Window
	// textFields of the UI
	title *ui.Entry
	body  *ui.Entry
	icon  *ui.Entry
	token *ui.Entry
	auth  *ui.Entry
	key   *ui.Entry
	// Checkbox
	checkbox *ui.Checkbox
	// Button
	submit *ui.Button
	// Firebase
	firebaseApp = sender.FirebaseApp{}
	// Old GCM Sender
	gcmSender = sender.GCMSender{}
)

type Gui struct{}

// Design the UI programmaticaly
func designUI() {
	// fields
	title = ui.NewEntry()
	body = ui.NewEntry()
	icon = ui.NewEntry()
	token = ui.NewEntry()
	auth = ui.NewEntry()
	key = ui.NewEntry()
	// button
	submit = ui.NewButton("Submit")
	// Checkbox
	checkbox = ui.NewCheckbox("Legacy GCM provider")
	// Append
	box := appendBox()
	// event on button
	submit.OnClicked(handleClick)
	// Create the window
	createWindow(box)
}

func appendBox() *ui.Box {
	// the window
	box := ui.NewVerticalBox()

	// Append the element on the box
	box.Append(ui.NewLabel("Title"), false)
	box.Append(title, false)

	box.Append(ui.NewLabel("Body"), false)
	box.Append(body, false)

	box.Append(ui.NewLabel("Icon"), false)
	box.Append(icon, false)

	// stack for the token part
	tokenBox := ui.NewVerticalBox()

	tokenBox.Append(ui.NewLabel("Token"), false)
	tokenBox.Append(token, false)

	tokenBox.Append(ui.NewLabel("UserAuth"), false)
	tokenBox.Append(auth, false)

	tokenBox.Append(ui.NewLabel("UserPublicKey"), false)
	tokenBox.Append(key, false)

	box.Append(tokenBox, false)
	box.Append(checkbox, false)
	box.Append(submit, true)

	return box
}

// Handle the click
func handleClick(*ui.Button) {
	// Get the value of the checkbox
	senderPreference := checkbox.Checked()

	// Retrieve the text of the gui
	m := message.Message{
		Title:   title.Text(),
		Message: body.Text(),
		Icon:    icon.Text(),
		UserInfo: message.User{
			Token: token.Text(),
			Keys: message.Key{
				UserAuth:   auth.Text(),
				UserPubKey: key.Text(),
			},
		},
	}

	if senderPreference {
		go func() {
			err := gcmSender.SendWebPush(m)

			if err != nil {
				fmt.Println("error ", err.Error())
			}
		}()

	} else {
		wrapCallFcmSDK(m)
	}
}

// Call the FCM SDK
func wrapCallFcmSDK(m message.Message) {

	userToken := token.Text()

	tokens := [1]string{userToken}
	payloads := m.PreparePayload(tokens[:])

	// send the payload using goroutines
	go firebaseApp.Send(payloads)
}

// Create the window
func createWindow(b *ui.Box) {
	window = ui.NewWindow("Push push", 500, 400, false)
	window.SetMargined(true)
	window.SetChild(b)
	window.Show()
}

// Create the UI
func (Gui) MakeUI() {
	// Init the firebase app
	firebaseApp.Init()
	err := ui.Main(designUI)

	if err != nil {
		panic(err)
	}
}
