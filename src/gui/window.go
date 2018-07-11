package gui

import (
	"fmt"
	"math/rand"
	"message"
	"sender"

	"github.com/andlabs/ui"
)

var (
	window *ui.Window
	// textFields of the UI
	title   *ui.Entry
	body    *ui.Entry
	image   *ui.Entry
	token   *ui.Entry
	auth    *ui.Entry
	key     *ui.Entry
	pic     *ui.Entry
	a4sId   *ui.Entry
	a4sicon *ui.Entry
	a4surl  *ui.Entry
	// Checkbox
	checkbox *ui.Checkbox
	a4sb     *ui.Checkbox
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
	image = ui.NewEntry()
	token = ui.NewEntry()
	auth = ui.NewEntry()
	key = ui.NewEntry()
	pic = ui.NewEntry()
	a4sId = ui.NewEntry()
	a4sicon = ui.NewEntry()
	a4surl = ui.NewEntry()
	// button
	submit = ui.NewButton("Submit")
	// Checkbox
	checkbox = ui.NewCheckbox("Legacy GCM provider")
	a4sb = ui.NewCheckbox("Button")
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

	box.Append(ui.NewLabel("Image"), false)
	box.Append(image, false)

	box.Append(ui.NewLabel("Buttons Push"), false)
	box.Append(a4sb, false)

	box.Append(ui.NewLabel("Big pictures"), false)
	box.Append(pic, false)

	box.Append(ui.NewLabel("a4sid"), false)
	box.Append(a4sId, false)

	box.Append(ui.NewLabel("a4sicon"), false)
	box.Append(a4sicon, false)

	box.Append(ui.NewLabel("a4surl"), false)
	box.Append(a4surl, false)

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

// build message
func buildMessage() message.Message {
	m := message.Message{
		Title:    title.Text(),
		Message:  body.Text(),
		Icon:     image.Text(),
		Pictures: pic.Text(),
		UserInfo: message.User{
			Token: token.Text(),
			Keys: message.Key{
				UserAuth:   auth.Text(),
				UserPubKey: key.Text(),
			},
		},
		Cparams: message.CustomParams{
			A4Sicon:   a4sicon.Text(),
			A4Sid:     a4sId.Text(),
			A4Surl:    a4surl.Text(),
			A4Sparams: []string{"|pid|", "|lat|", "|lon|"},
		},
		A4Sid: a4sId.Text(),
	}

	a4sbchecked := a4sb.Checked()
	if a4sbchecked {
		arr := make([]message.A4Sb, 2)
		arr[0] = message.A4Sb{
			Title: "ok",
			Icon:  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTa3zmeMGSXFgQwfJ5T9Vix-qq1Mdeq1TIjmXfoOrAhdx9M9QD3PQ",
			ID:    rand.Intn(100),
		}

		arr[1] = message.A4Sb{
			Title: "no",
			Icon:  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSYfX1gX8_u70NS95b1BqoyT6mlWdcHttEtlm0UNx856GZGtJCB",
			ID:    rand.Intn(100),
		}

		m.Button = arr
	}

	return m
}

// Handle the click
func handleClick(*ui.Button) {
	// Get the value of the checkbox
	senderPreference := checkbox.Checked()
	// Retrieve the text of the gui
	m := buildMessage()
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
