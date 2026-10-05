package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var notes []string

func NotesHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		newNotes := strings.Join(notes, ",")
		fmt.Fprintln(w, newNotes)
	} else if r.Method == http.MethodPost {
		defer r.Body.Close()
		bodyText, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("failed read body request:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if string(bodyText) == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Note content cannot be empty")
			return
		}

		notes = append(notes, string(bodyText))
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Note added successfully")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
