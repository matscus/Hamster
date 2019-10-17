package users

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"

	"github.com/matscus/Hamster/Package/Users/subset"
)

//User - struct for user, contains username, passowd (in base64) and JWT
type User struct {
	ID       int64
	User     string
	Password string
	jwt.StandardClaims
}

//New - return inretface user
func (u *User) New() subset.User {
	var user subset.User
	user = u
	return user
}

//NewTokenString - func for generate and response new token
func (u User) NewTokenString() (token string, err error) {
	projects, err := client.PGClient{}.New().GetAllUserProject(u.User)
	if err != nil {
		return "", err
	}
	token, err = jwttoken.Token{}.New().Generate(u.User, projects)
	if err != nil {
		return "", err
	}
	return token, err
}

//CheckUser - func for check users (for )validation of user token)
func (u *User) CheckUser() (res bool, err error) {
	hash, err := client.PGClient{}.New().GetUserHash(u.User)
	password, err := b64.StdEncoding.DecodeString(u.Password)
	ok := compareHash(hash, password)
	if ok {
		return true, nil
	}
	err = errors.New("check user fail")
	return false, err
}

func compareHash(hash string, password []byte) bool {
	h := sha256.New()
	h.Write(password)
	pass := fmt.Sprintf("%x", h.Sum(nil))
	if hash == pass {
		return true
	}
	return false
}
