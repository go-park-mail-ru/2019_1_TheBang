package chat

import "net/http"

func DummyBeforeChat(w http.ResponseWriter, r *http.Request) {
	w.Write(home)
}
