package subset

//Project - interface for projects
type Project interface {
	Create() error
	Update() error
	Delete() error
}
