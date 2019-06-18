package model

import "github.com/dgrijalva/jwt-go"

//User defines the User Schema
type User struct {
	Name        string
	Email       string
	University  string
	Password    string
	CurQuestion int
	Admin       bool
}

//Claims is for JWT payload
type Claims struct {
	Email string
	jwt.StandardClaims
}
