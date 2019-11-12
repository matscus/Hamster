package subset

//Role - interface from roles
type Role interface {
	Create() error
	Update() error
	Delete() error
}
