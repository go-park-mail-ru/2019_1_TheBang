package config

import (
	"os"

	"google.golang.org/grpc"
)

var (
	SECRET      = getSecret()
	FrontentDst = getFrontDest()

	CookieName = "bang_token"
	DefaultImg = "default_img"

	AuthServerAddr   = "127.0.0.1"
	PointsServerAddr = "127.0.0.1"

	AUTHPORT   = getAuhtPort()
	CHATPORT   = getChatPort()
	MAINPORT   = getMainPort()
	POINTSPORT = getPointsPort()
	GAMEPORT   = getGamePort()

	AuthConn = getAuthConn()
)

func getAuthConn() *grpc.ClientConn {
	conn, err := grpc.Dial(AuthServerAddr+":"+AUTHPORT, grpc.WithInsecure())
	if err != nil {
		Logger.Fatal("Can not get grpc connection to auth server")

		return nil
	}

	return conn
}

func getAuhtPort() string {
	port := os.Getenv("AUTHPORT")
	if port == "" {
		Logger.Warn("There is no AUTH_PORT!")
		port = "50051"
	}

	return port
}

func getChatPort() string {
	port := os.Getenv("CHATPORT")
	if port == "" {
		Logger.Warn("There is no CHATPORT!")
		return "8003"
	}

	return port
}

func getGamePort() string {
	port := os.Getenv("GAMEPORT")
	if port == "" {
		Logger.Warn("There is no GAMEPORT!")
		port = "8002"
	}

	return port
}

func getMainPort() string {
	port := os.Getenv("MAINPORT")
	if port == "" {
		Logger.Warn("There is no MAINPORT!")
		port = "8001"
	}

	return port
}

func getPointsPort() string {
	port := os.Getenv("POINTSPORT")
	if port == "" {
		Logger.Warn("There is no POINTSPORT!")
		port = "50052"
	}

	return port
}

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
		dst = "http://95.163.212.32:8009"
	}

	return dst
}
