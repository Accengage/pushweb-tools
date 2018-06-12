package main

import (
	"fmt"
	"sender"
)

var (
	firebaseApp = sender.FirebaseApp{}
)

func main() {
	// init firebase admin
	firebaseApp.Init()
	fmt.Println("init !")

	// Example of usage of the sender
	// create an array of tokens
	// Create a payload
	// Get the array of payload by passing the array of tokens
	// Send the push using the firebase App Struct instance
}
