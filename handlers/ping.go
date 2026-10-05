package handlers

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}
