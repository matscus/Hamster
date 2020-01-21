package users

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/matscus/Hamster/Package/Clients/client/postgres"

	"github.com/dgrijalva/jwt-go"
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
	DBClient *postgres.PGClient
	jwt.StandardClaims
}

//New - return inretface user
func New(client *postgres.PGClient) subset.User {
	var user subset.User
	user = User{DBClient: client}
	return user
}

//NewTokenString - func for generate and response new token
func (u User) NewTokenString(temp bool) (token string, err error) {
	userID, role, err := u.DBClient.GetUserIDAndRole(u.User)
	if err != nil {
		return "", err
	} else {
		projects, err := u.DBClient.GetUserProjects(userID)
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
func (u User) CheckUser() (res bool, err error) {
	hash, err := u.DBClient.GetUserHash(u.User)
	password, err := b64.StdEncoding.DecodeString(u.Password)
	ok := compareHash(hash, password)
	if ok {
		return true, nil
	}
	err = errors.New("check user fail")
	return false, err
}

//CheckPasswordExp - func for check PasswordExp (for )validation of user token)
func (u User) CheckPasswordExp() (res bool, err error) {
	exp, err := u.DBClient.GetUserPasswordExp(u.User)
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
func (u User) Create() error {
	h := sha256.New()
	pass, err := b64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}
	h.Write([]byte(pass))
	projectsID, err := u.DBClient.GetProjectsIDtoString(u.Projects)
	if err != nil {
		return err
	}
	return u.DBClient.NewUser(u.User, fmt.Sprintf("%x", h.Sum(nil)), u.Role, projectsID)
}

//Update - update user data
func (u User) Update() error {
	err := u.DBClient.UpdateUser(string(u.ID), u.Role, u.Projects)
	if err != nil {
		return err
	}
	projectsID, err := u.DBClient.GetProjectsIDtoString(u.Projects)
	if err != nil {
		return err
	}
	return u.DBClient.UpdatetUserProjects(u.ID, projectsID)
}

//Delete -delete user
func (u User) Delete() error {
	return u.DBClient.DeleteUser(u.ID)
}

//ChangePassword - change user password
func (u User) ChangePassword() error {
	h := sha256.New()
	pass, err := b64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}
	h.Write([]byte(pass))
	return u.DBClient.ChangeUserPassword(u.ID, fmt.Sprintf("%x", h.Sum(nil)))
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
