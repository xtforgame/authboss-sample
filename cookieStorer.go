package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"gopkg.in/authboss.v0"
)

var cookieStore *securecookie.SecureCookie

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStorer{w, r}
}

func (s CookieStorer) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = cookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	encoded, err := cookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Name:  key,
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
}