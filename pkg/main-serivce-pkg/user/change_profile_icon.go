package user

import (
	"2019_1_TheBang/config"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func ChangeProfileAvatarHandler(c *gin.Context) {
	token := TokenFromCookie(c.Request)
	info, status := InfoFromCookie(token)
	if status == http.StatusInternalServerError {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	profile, status := SelectUser(info.Nickname)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)

		return
	}
	defer file.Close()

	filein, err := header.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)

		return
	}
	defer filein.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}
	filename := hex.EncodeToString(hasher.Sum(nil))

	fileout, err := os.OpenFile("tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}
	defer fileout.Close()

	_, err = io.Copy(fileout, filein)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	ok := UpdateUserPhoto(info.Nickname, filename)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	DeletePhoto(profile.Photo)
	profile.Photo = filename

	ss := GetFreashToken(profile.Nickname)
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     config.CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "",
	}
	http.SetCookie(c.Writer, &cookie)

	c.Status(http.StatusOK)
}
