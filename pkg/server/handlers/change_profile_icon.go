package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/auth"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
)

func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusOK

	if r.Method == "OPTIONS" {
		return
	}

	token, ok := auth.CheckTocken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)

		return
	}
	nickname, status := NicknameFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)

		return
	}

	profile, status := models.SelectUser(nickname)
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		config.Logger.Infow("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusUnprocessableEntity)

		return
	}
	defer file.Close()

	filein, err := header.Open()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		config.Logger.Infow("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusUnprocessableEntity)

		return
	}
	defer filein.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}
	filename := hex.EncodeToString(hasher.Sum(nil))

	fileout, err := os.OpenFile("tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "file for img was not created!")

		return
	}
	defer fileout.Close()

	_, err = io.Copy(fileout, filein)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "img was not saved on disk!")

		return
	}

	ok = models.UpdateUserPhoto(nickname, filename)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "can not update photo name with sql!")

		return
	}

	deletePhoto(profile.Photo)
	profile.Photo = filename

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", err.Error())

		return
	}
	config.Logger.Infow("GetIconHandler",
		"RemoteAddr", r.RemoteAddr,
		"status", http.StatusOK)
}

func deletePhoto(filename string) {
	if filename == config.DefaultImg {
		return
	}

	err := os.Remove("tmp/" + filename)
	if err != nil {
		config.Logger.Warnw("GetIconHandler",
			"filename", filename,
			"warn", err.Error())

		return
	}
}
