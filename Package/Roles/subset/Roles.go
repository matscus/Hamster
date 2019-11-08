package subset

//Role - interface from roles
type Role interface {
	IfExist() (bool, error)
	Create() error
	Update() error
	Delete() error
}
