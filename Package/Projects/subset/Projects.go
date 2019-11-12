package subset

//Project - interface for projects
type Project interface {
	Create() error
	Delete() error
}
