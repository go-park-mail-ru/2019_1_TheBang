package config

import "os"

var (
	SECRET       = getSecret()
	FrontentDst  = getFrontDest()
	CookieName   = "bang_token"
	AuthServer   = "127.0.0.1:50051" // убрать хардкод, заменить на переменные окружения
	PointsServer = "127.0.0.1:50052" // убрать хардкод, заменить на переменные окружения
)

func getSecret() []byte {
	secret := []byte(os.Getenv("SECRET"))
	if string(secret) == "" {
		Logger.Warn("There is no SECRET!")
		secret = []byte("SECRET")
	}

	return secret
}

func getFrontDest() string {
	dst := os.Getenv("FrontentDst")
	if dst == "" {
		Logger.Warn("There is no FrontentDst!")
		dst = "http://127.0.0.1:3000"
	}

	return dst
}
