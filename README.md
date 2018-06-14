## Pushweb Utils

This is just a prototype app that help you to send push either by using the GCM provider or the FCM SDK

## Installation

Clone this repo, make sure that you have set the **GOPATH**.
Install the following dependencies using the **go get** command

```
https://github.com/andlabs/ui
https://github.com/firebase/firebase-admin-go
https://github.com/gauntface/web-push-go
```

## Trick

The legacy GCM Sender (no VAPID) use the gautface sender dependencies. (Note that the repo is archive however it does the work...). As the lib is unmaintained we need to update the **tempGcmURL** like below

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

## Note

If executing the **cheesy** executable give you a **Permission denied** error please use the command ```chmod +x cheesy``` and try it again
