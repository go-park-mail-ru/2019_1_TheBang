package test 

import (
	"2019_1_TheBang/api"
	"testing"

	jwriter "github.com/mailru/easyjson/jwriter"
)

func TestEasyJson(t *testing.T) {
	r := api.Profile{}

	bytes, _ := r.MarshalJSON()

	w := &jwriter.Writer{}

	r.MarshalEasyJSON(w)

	_ = r.UnmarshalJSON(bytes)
}
