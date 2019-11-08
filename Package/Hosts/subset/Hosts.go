package subset

type Host interface {
	Create() error
	Update() error
	Delete() error
	IfExist() (bool, error)
}
