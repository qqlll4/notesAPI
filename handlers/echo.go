package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		role := r.Header.Get("X-User-Role")
		if role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Access denied: admin role required")
			return
		}

		text, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("the text was not read:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err = w.Write([]byte(text))
		if err != nil {
			fmt.Println("failed to send response:", err)
			return
		}
	}
}
