package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/context"
	"github.com/mayankshah1607/Enigma-Microservices/admin/iohandlers"
	"github.com/mayankshah1607/Enigma-Microservices/admin/model"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func jsonBodyParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Parse the incoming request to get the jsonbody
		//Get []byte from r.Body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		//This area needs some improvement
		switch r.RequestURI {
		case "/submit":
			req, err := iohandlers.DecodeSubmitRequest(b)
			if err != nil {
				log.Println("Error while decoding request")
				http.Error(w, err.Error(), 500)
			}
			context.Set(r, "req", req)

		case "/delete":
			req, err := iohandlers.DecodeDeleteRequest(b)
			if err != nil {
				log.Println("Error while decoding request")
				http.Error(w, err.Error(), 500)
			}
			context.Set(r, "req", req)

		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " - ", r.RequestURI)
		w.Header().Add("Content-Type", "application/json")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("enigma_auth")
		if err != nil {
			log.Println("Cookie not found!")
			return
		}
		tknString := c.Value
		claims := &model.Claims{}
		jwtKey, _ := os.LookupEnv("JWT_KEY")

		tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//If token is valid, get Email and check in database
		email := claims.Email
		exists := model.CheckAdmin(email)
		if !exists {
			log.Println("Admin log in attempt from unauthorized user : ", email)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
