package jwttoken

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/matscus/Hamster/Package/JWTToken/subset"
)

//Token - struct for token
type Token struct {
	Token string `json:"token"`
}

func (t Token) New() subset.Token {
	var token subset.Token
	token = Token{}
	return token
}

func (t Token) Generate(role string, user string, projects []string) (tokenstring string, err error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"user":    user,
		"role":    role,
		"project": projects,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenstring, err = token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		return tokenstring, err
	}
	return tokenstring, err
}

//Check  - func for check validate JWT, result bool
func (t Token) Check() bool {
	token, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	if err == nil && token.Valid {
		return true
	}
	return false
}

//Parse - func to parse token and check to valid
func Parse(t string) bool {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	if err == nil && token.Valid {
		return true
	}
	return false
}

//IsAdmin - func to parse token and check to valid
func IsAdmin(t string) bool {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"]
		if role == "admin" {
			return true
		} else {
			return false
		}
	}
	return false
}
