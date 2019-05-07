package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"2019_1_TheBang/config"
	pb "2019_1_TheBang/pkg/public/protobuf"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

type CustomClaims struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	PhotoURL string `json:"photo_url"`

	jwt.StandardClaims
}

func GetUserInfo(token string) (user *pb.UserInfo, myerr error) {
	conn, err := grpc.Dial(config.AuthServerAddr+":"+config.AUTHPORT, grpc.WithInsecure())
	if err != nil {
		myerr = fmt.Errorf("did not connect: %v", err.Error())

		return
	}
	defer conn.Close()

	client := pb.NewCookieCheckerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.CheckCookie(ctx, &pb.CookieRequest{JwtToken: token})
	if err != nil {
		myerr = fmt.Errorf("could not check: %v", err.Error())

		return
	}

	if !r.Valid {
		myerr = errors.New("invalid cookie")

		return
	}

	user = r.User
	return
}

func CheckTocken(r *http.Request) (info *pb.UserInfo, ok bool) {
	if r == nil {
		config.Logger.Warnw("CheckTocken -> get cookie:",
			"warn", "request is nil")

		return
	}

	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		config.Logger.Warnw("CheckTocken -> get cookie:",
			"warn", err.Error())

		return
	}

	tokenStr := cookie.Value
	info, err = GetUserInfo(tokenStr)
	if err != nil {
		config.Logger.Warnw("CheckTocken -> GetUserInfo:",
			"warn", err.Error())
	}

	return info, true
}
