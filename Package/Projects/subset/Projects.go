package subset

//Project - interface for projects
type Project interface {
	IfExist() (bool, error)
	Create() error
	Update() error
	Delete() error
}
