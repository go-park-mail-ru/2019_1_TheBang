package test

import (
	"2019_1_TheBang/pkg/auth-service-pkg/authchecker"
	"testing"
)

func TestAuthSUCCESS(t *testing.T) {
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJpZCI6MTQ0LCJuaWNrbmFtZSI6IjMyIiwicGhvdG9fdXJsIjo" +
		"iODI3MzU3Mzk4NTMzNWYyMDJjZTQ3ZTdlOGIyZjU5MzMiLCJpc3M" +
		"iOiJUaGVCYW5nIHNlcnZlciJ9.Std_bFydKdPRvr5f3iKSRevQOzdKS1EshcsTasvJkTM"

	info, err := authchecker.InfoFromCookie(token)
	if err != nil {
		t.Errorf("TestAuthSUCCESS: valid cookie was failed! Error: %v", err.Error())
	}
	_ = info
}

func TestAuthFAIL(t *testing.T) {
	const token = "token"

	info, err := authchecker.InfoFromCookie(token)
	if err == nil {
		t.Error("TestAuthFAIL: valid cookie was accepted!")
	}
	_ = info
}
