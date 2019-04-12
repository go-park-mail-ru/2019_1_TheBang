package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"2019_1_TheBang/config"
)

func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
	var status = http.StatusOK

	if r.Method == "OPTIONS" {
		return
	}

	token := TokenFromCookie(r)
	info, status := InfoFromCookie(token)
	if status == http.StatusInternalServerError {
		w.WriteHeader(status)

		return
	}

	profile, status := SelectUser(info.Nickname)
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}
	defer file.Close()

	filein, err := header.Open()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

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

	ok := UpdateUserPhoto(info.Nickname, filename)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "can not update photo name with sql!")

		return
	}

	DeletePhoto(profile.Photo)
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
}
