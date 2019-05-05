package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"

	"github.com/gin-gonic/gin"
)

func TestChangeProfileAvatarHandler(t *testing.T) {
	cookie, _ := GetTESTAdminCookie()
	defer DeleteTESTAdmin()

	path := "/user/avatar"

	pathOS := "tmp/" + mainconfig.DefaultImg
	req, err := newfileUploadRequest(path, map[string]string{}, "photo", pathOS)
	if err != nil {
		t.Fatal(err.Error())
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST(path, user.ChangeProfileAvatarHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestChangeProfileAvatarHandler, have not cookie: expected %v, have %v!\n",
			http.StatusOK, rr.Code)
	}
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
