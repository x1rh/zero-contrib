package jwtx

import "github.com/golang-jwt/jwt/v4"

// type User struct {
// 	Username string `json:"username"`
// 	Uid      uint64 `json:"uid"`
// }

type Claims struct {
	Payload any `json:"payload"` 
	jwt.RegisteredClaims
}
