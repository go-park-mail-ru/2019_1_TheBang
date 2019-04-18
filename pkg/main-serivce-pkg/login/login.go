package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/pkg/public/auth"

	"github.com/dgrijalva/jwt-go"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	login := api.Login{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("LogoutHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}

	err = json.Unmarshal(body, &login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("LogoutHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}

	ss, status := LoginAcount(login.Nickname, login.Passwd)
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     config.CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	prof, status := user.SelectUser(login.Nickname)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("LogoutHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "Can not find valid user")

		return
	}

	err = json.NewEncoder(w).Encode(prof)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("LogoutHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}

}

func LoginAcount(username, passwd string) (ss string, status int) {
	ok := user.CheckUser(username, passwd)
	if !ok {
		return ss, http.StatusUnauthorized
	}

	prof, _ := user.SelectUser(username)

	// логирование
	fmt.Println(prof)

	claims := auth.CustomClaims{
		prof.Id,
		prof.Nickname,
		prof.Photo,
		jwt.StandardClaims{
			Issuer: mainconfig.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SECRET)
	if err != nil {
		config.Logger.Warnw("LoginAcount",
			"Error with JWT tocken generation:", err.Error())

		return ss, http.StatusInternalServerError
	}

	return ss, http.StatusOK
}
