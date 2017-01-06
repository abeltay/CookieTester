package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/set", SetCookieHandler)
	mux.HandleFunc("/get", GetCookieHandler)
	mux.HandleFunc("/delete", DeleteCookieHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("_session_id")
	if err != nil {
		fmt.Fprintln(w, "cookie cannot be found")
	} else {
		fmt.Fprintln(w, sessionID.Value)
	}
}
func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "_session_id",
		Value:    "123",
		MaxAge:   1800,
		HttpOnly: true,
		Domain:   ".lvh.me",
		Path:     "/",
	}

	http.SetCookie(w, &c)
	fmt.Fprintf(w, "<html><head><title>Create Cookie</title></head><body>Cookie Created</body></html>")
}

func DeleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "_session_id",
		MaxAge:   -1,
		HttpOnly: true,
		Domain:   ".lvh.me",
		Path:     "/",
	}

	http.SetCookie(w, &c)
	fmt.Fprintf(w, "<html><head><title>Delete Cookie</title></head><body>Cookie Deleted</body></html>")
}
