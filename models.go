package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Login struct {
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
}

type Update struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	DOB     string `json:"dob"`
}

func CheckTocken(r *http.Request) (token *jwt.Token, ok bool) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil, false
	}

	tokenStr := cookie.Value

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})
	if err != nil {
		log.Printf("Error with check tocken: %v", err.Error())

		return nil, false
	}

	if !token.Valid {
		log.Printf("%v use faked cookie: %v\n", r.RemoteAddr, err.Error())

		return nil, false
	}

	return token, true
}

func CreateAccount(w http.ResponseWriter, r *http.Request) (prof Profile, err error) {
	signup := Signup{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return prof, err
	}

	err = json.Unmarshal(body, &signup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())

		return prof, err
	}

	prof = Profile{
		Nickname: signup.Nickname,
		Name:     signup.Name,
		Surname:  signup.Surname,
		DOB:      signup.DOB,
	}
	passwd := signup.Passwd

	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	storageProf.mu.Lock()
	defer storageProf.mu.Unlock()

	if _, ok := storageAcc.data[prof.Nickname]; ok {
		w.WriteHeader(http.StatusConflict)
		err := errors.New("This user already exists!")

		return prof, err
	}

	prof.Photo = DefaultImg

	storageAcc.data[prof.Nickname] = passwd
	storageProf.data[prof.Nickname] = prof

	return prof, nil
}

func NicknameFromCookie(w http.ResponseWriter, r *http.Request) (nickname string, err error) {
	token, ok := CheckTocken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		err := errors.New("You can not get profile info!")

		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		nickname = claims["nickname"].(string)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("MyProfileInfoHandler: Error with parsing token's claims")

		return "", err
	}

	return nickname, err
}

func LoginAcount(username, passwd string) (string, error) {
	storageAcc.mu.Lock()
	defer storageAcc.mu.Unlock()

	if pw, ok := storageAcc.data[username]; !ok || pw != passwd {
		err := errors.New("Wrong answer or password!")
		return "", err
	}

	claims := customClaims{
		username,
		jwt.StandardClaims{
			Issuer: ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET)
	if err != nil {
		log.Printf("Error with JWT tocken generation: %v\n", err.Error())
	}

	return ss, nil
}

func deletePhoto(filename string) {
	if filename == DefaultImg {
		return
	}

	err := os.Remove("tmp/" + filename)
	if err != nil {
		log.Printf("Can not remove file tmp/%v\n", filename)
	}
}
