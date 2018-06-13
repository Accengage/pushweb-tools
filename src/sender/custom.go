package sender

// Because we don't use the VAPID...
func sendWebPush() error {

	encrypt, err := webpushencrypto.GetEncryptoMessage()
}
