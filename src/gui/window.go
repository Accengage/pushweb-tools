package gui

import (
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
	// Firebase
	firebaseApp = sender.FirebaseApp{}
)

type Gui struct{}

// Design the UI programmaticaly
func designUI() {
	// fields
	title = ui.NewEntry()
	body = ui.NewEntry()
	icon = ui.NewEntry()
	token = ui.NewEntry()

	// button
	submit := ui.NewButton("Submit")

	// the window
	box := ui.NewVerticalBox()

	// Append the element on the box
	box.Append(ui.NewLabel("Title"), false)
	box.Append(title, false)

	box.Append(ui.NewLabel("Body"), false)
	box.Append(body, false)

	box.Append(ui.NewLabel("Icon"), false)
	box.Append(icon, false)

	box.Append(ui.NewLabel("Token"), false)
	box.Append(token, false)
	box.Append(submit, true)

	// event on button
	submit.OnClicked(handleClick)

	// Create the window
	createWindow(box)
}

// Handle the click
func handleClick(*ui.Button) {
	// Retrieve the text of the gui
	m := message.Message{
		Title:   title.Text(),
		Message: body.Text(),
		Icon:    icon.Text(),
	}

	userToken := token.Text()

	tokens := [1]string{userToken}
	payloads := m.PreparePayload(tokens[:])

	// send the payload
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
