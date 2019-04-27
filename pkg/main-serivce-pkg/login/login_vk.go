package login

import (
	"2019_1_TheBang/api"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/config"
	"github.com/gin-gonic/gin"
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/oauth2/vk"
	"golang.org/x/oauth2"
)

const (
	AppId     = "6943828"
	AppKey    = "HdcqD2VBSPFtTx82XiGX"
	AppSecret = "69eb8e0269eb8e0269eb8e022c69827a56669eb69eb8e023559ea1b8cb9e8c7b3ab9f55"
)

var vkConfig = oauth2.Config{
	ClientID:     AppId,
	ClientSecret: AppKey,
	RedirectURL:  "http://127.0.0.1:8001/oauth/vk/authorize",
	Endpoint:     vk.Endpoint,
	Scopes:       []string{"email", "photos"},
}

type Response struct {
	Response []map[string]interface{}
}

func VKAuthConnect(c *gin.Context) {
	url := vkConfig.AuthCodeURL("", oauth2.SetAuthURLParam("display", "popup"))
	// http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GenerateApiUrl(accessToken string, fields ...string) string {
	baseUrl := "https://api.vk.com/method/users.get?fields=%s&access_token=%s&v=5.52"
	fieldsString := strings.Join(fields, ",")
	return fmt.Sprintf(baseUrl, fieldsString, accessToken)
}

func VKAuthAuthorize(c *gin.Context) {
	ctx := c
	code := c.Query("code")

	//retrieving user's access-token
	token, err := vkConfig.Exchange(ctx, code)
	if err != nil {
		log.Println("cannot exchange", err)
		c.JSONP(http.StatusBadRequest, ":(")
		return
	}

	//creating client with user privileges
	client := vkConfig.Client(ctx, token)

	//getting dataec
	ApiUrl := GenerateApiUrl(token.AccessToken, "id", "photo_100")
	resp, err := client.Get(ApiUrl)
	if err != nil {
		log.Println("cannot request data", err)
		c.JSONP(http.StatusBadRequest, ":(")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read buffer", err)
		c.JSONP(http.StatusBadRequest, ":(")
		return
	}

	// email := token.Extra("email").(string)
	fmt.Println("%v", token)
	email := token.Extra("email")
	// user_id := token.Extra("user_id").(float64)
	// photo_url := token.Extra("photo_url").(string)
	// first_name := token.Extra("first_name").(string)
	// last_name := token.Extra("last_name").(string)
	// fmt.Println(user_id, first_name, last_name)

	data := &Response{}
	_ = json.Unmarshal(body, data)
	fmt.Println(data.Response)
	time.Sleep(1000)
	fmt.Println(email)

	first_name := data.Response[0]["first_name"].(string)
	last_name := data.Response[0]["last_name"].(string)
	user_id := data.Response[0]["id"].(float64)
	photo_url := data.Response[0]["photo_100"].(string)
	id_str := fmt.Sprintf("%d", int64(user_id))
	fmt.Println(id_str, first_name, last_name, photo_url)

	// username := data.Response[0]["domain"].(string)
	// password := "nil"

	signup := &api.Signup{}
	signup.Nickname = first_name + id_str
	signup.Name = first_name
	signup.Surname = last_name
	signup.Passwd = "here"
	signup.DOB = "2018-01-01"

	status := user.CreateUser(signup)
	if status != http.StatusCreated && status != http.StatusConflict {
		c.AbortWithStatus(status)

		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ss, status := LoginAcount(signup.Nickname, "here")
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     config.CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	prof, status := user.SelectUser(signup.Nickname)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	c.JSONP(http.StatusOK, prof)

	// c.Redirect(http.StatusTemporaryRedirect, "/user")
}