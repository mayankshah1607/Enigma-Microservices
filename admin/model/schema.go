package model

import "github.com/dgrijalva/jwt-go"

//Question is the struct of a question
type Question struct {
	Text     string
	Image    string
	Answer   string
	SolvedBy int
}

//Claims is for JWT payload
type Claims struct {
	Email string
	jwt.StandardClaims
}

//User defines the User Schema
type User struct {
	Name        string
	Email       string
	University  string
	Password    string
	CurQuestion int
	Admin       bool
}
