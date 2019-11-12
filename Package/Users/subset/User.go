package subset

type User interface {
	NewTokenString(temp bool) (token string, err error)
	CheckUser() (res bool, err error)
	CheckPasswordExp() (res bool, err error)
	ChangePassword() error
	Create() error
	Update() error
	Delete() error
}
