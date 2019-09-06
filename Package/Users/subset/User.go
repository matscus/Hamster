package subset

type User interface {
	CheckUser() (res bool, err error)
	NewTokenString() (tokenstring string, err error)
}
