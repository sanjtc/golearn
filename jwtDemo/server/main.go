package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", rootHandler)
	log.Println("start server")

	if err := http.ListenAndServe("127.0.0.1:2333", nil); err != nil {
		log.Println(err)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	username := query.Get("username")

	log.Println("username:", username)

	// get jwt
	hmacSampleSecret := []byte("test")

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
		},
	)

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("tokenString: ", tokenString)

	jwtCookie := &http.Cookie{
		Name:  "tokenString",
		Value: tokenString,
	}

	http.SetCookie(w, jwtCookie)

	_, err = w.Write([]byte("received"))
	if err != nil {
		log.Println(err)
		return
	}
}
