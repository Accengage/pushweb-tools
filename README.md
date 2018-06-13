## Pushweb Utils

This is just a prototype app that help you to send push either by using the GCM provider or the FCM SDK

## Installation

Clone this repo, make sure that your gopath is set.
Install the dependencies by running this command from the root of the project

```
go get ./...
```

## Trick

The legacy GCM Sender (no VAPID) use the archive gauntface sender (webpush package). However as the library use the previous endpoint. Please update the dependency like below 

- old: 
```
tempGcmURL = https://gcm-http.googleapis.com/gcm
```

- update
```
tempGcmURL = "https://fcm.googleapis.com/fcm/send"
```

## Build

Build the command ```go build cheesy.go```
Run the project by using the command ```./cheesy```