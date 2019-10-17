package subset

type User interface {
	NewTokenString() (tokenstring string, err error)
	CheckUser() (res bool, err error)
}
