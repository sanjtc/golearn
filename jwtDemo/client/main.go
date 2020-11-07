package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	username := "username=xinwang"

	resp, err := http.Post("http://127.0.0.1:2333", "text/plain", bytes.NewReader([]byte(username)))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	received, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(received))

	// get token
	jwtCookie := getTokenStringFromCookies(resp.Cookies())
	if jwtCookie == nil {
		log.Println("no jwt cookie")
		return
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	}

	token, err := jwt.ParseWithClaims(
		jwtCookie.Value,
		jwt.MapClaims{},
		keyFunc,
	)
	if err != nil {
		log.Println(err)
		return
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("cannot convert claims to mapclaims")
		return
	}

	log.Println(mapClaims["username"])
}

func getTokenStringFromCookies(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == "tokenString" {
			return cookie
		}
	}

	return nil
}
