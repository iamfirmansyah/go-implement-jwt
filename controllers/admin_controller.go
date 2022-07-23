package controllers

import "net/http"

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}

	w.Write([]byte("Welcome, Admin."))
}
