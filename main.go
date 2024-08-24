package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nartvt/session-management/app/httputils"
)

var session = httputils.NewSessionHttp("nartvtsessionCookie", 30*time.Second)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/set-expiration", setExpiration)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/destroy", destroy)
	fmt.Println("Starting server on: 8088")
	http.ListenAndServe(":9088", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	err := session.Set(w, "user_name", "nartvt")
	if err != nil {
		http.Error(w, "Unable to set session value", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session value set")
}

func get(w http.ResponseWriter, r *http.Request) {
	value, err := session.Get(r, "user_name")
	if err != nil {
		http.Error(w, "Unable to get session value", http.StatusInternalServerError)
		return
	}
	if value == "" {
		fmt.Fprintln(w, "No session value found")
		return
	}
	fmt.Fprintf(w, "Session value: %s\n", value)
}

func setExpiration(w http.ResponseWriter, t *http.Request) {
	session.SetExpiration(30 * time.Minute)
	fmt.Fprintln(w, "Session expiration set to 30 minutes")
}

func remove(w http.ResponseWriter, r *http.Request) {
	err := session.Remove(w, "user_name")
	if err != nil {
		http.Error(w, "Unable to remove session value", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session value removed")
}

func destroy(w http.ResponseWriter, r *http.Request) {
	err := session.Destroy(r, w)
	if err != nil {
		http.Error(w, "Unable to destroy session", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session destroyed")
}
