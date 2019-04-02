// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	http.ServeFile(w, r, "home.html")
}
