package config

import "os"

var (
	SECRET      = getSecret()
	FrontentDst = getFrontDest()

	CookieName       = "bang_token"
	AuthServerAddr   = "127.0.0.1"
	AuthServerPort   = "50051"
	PointsServerAddr = "127.0.0.1"
	PointsServerPort = "50052"
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
		dst = "http://127.0.0.1:8010"
	}

	return dst
}
