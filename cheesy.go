package main

import (
	"fmt"
	"gui"
	"sender"
)

var (
	firebaseApp = sender.FirebaseApp{}
	appGui      = gui.Gui{}
)

func main() {
	// init firebase admin
	firebaseApp.Init()
	appGui.MakeUI()
	fmt.Println("init !")
	// Example of usage of the sender
	// create an array of tokens
	// Create a payload
	// Get the array of payload by passing the array of tokens
	// Send the push using the firebase App Struct instance
}
