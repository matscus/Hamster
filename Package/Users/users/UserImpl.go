package users

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"

	"github.com/matscus/Hamster/Package/Users/subset"
)

//User - struct for user, contains username, passowd (in base64) and JWT
type User struct {
	ID       int64    `json:id`
	User     string   `json:user`
	Password string   `json:password,omitempty`
	Role     string   `json:"role"`
	Projects []string `json:"projects"`
	jwt.StandardClaims
}

//New - return inretface user
func (u *User) New() subset.User {
	var user subset.User
	user = u
	return user
}

//NewTokenString - func for generate and response new token
func (u User) NewTokenString(temp bool) (token string, err error) {
	client := client.PGClient{}.New()
	userID, role, err := client.GetUserIDAndRole(u.User)
	if err != nil {
		return "", err
	} else {
		projects, err := client.GetUserProjects(userID)
		if err != nil {
			return "", err
		} else {
			if temp {
				token, err = jwttoken.Token{}.New().GenerateTemp(role, u.User, projects)
				if err != nil {
					return "", err
				}
			} else {
				token, err = jwttoken.Token{}.New().Generate(role, u.User, projects)
				if err != nil {
					return "", err
				}
			}
		}
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

//CheckPasswordExp - func for check PasswordExp (for )validation of user token)
func (u *User) CheckPasswordExp() (res bool, err error) {
	exp, err := client.PGClient{}.New().GetUserPasswordExp(u.User)
	t, _ := time.Parse(time.RFC3339, exp)
	if t.Unix() == 0 {
		return true, nil
	} else {
		if time.Now().Unix() > t.Unix() {
			return false, err
		} else {
			return true, err
		}
	}
}

//Create - create new user and insert data to database
func (u *User) Create() error {
	h := sha256.New()
	pass, err := b64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}
	h.Write([]byte(pass))
	client := client.PGClient{}.New()
	projectsID, err := client.GetProjectsIDtoString(u.Projects)
	if err != nil {
		return err
	}
	return client.NewUser(u.User, fmt.Sprintf("%x", h.Sum(nil)), u.Role, projectsID)
}

//Update - update user data
func (u *User) Update() error {
	client := client.PGClient{}.New()
	err := client.UpdateUser(u.ID, u.Role)
	if err != nil {
		return err
	}
	projectsID, err := client.GetProjectsIDtoString(u.Projects)
	if err != nil {
		return err
	}
	return client.UpdatetUserProjects(u.ID, projectsID)
}

//Delete -delete user
func (u *User) Delete() error {
	return client.PGClient{}.New().DeleteUser(u.ID)
}

//ChangePassword - change user password
func (u *User) ChangePassword() error {
	client := client.PGClient{}.New()
	h := sha256.New()
	pass, err := b64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}
	h.Write([]byte(pass))
	return client.ChangeUserPassword(u.ID, fmt.Sprintf("%x", h.Sum(nil)))
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
