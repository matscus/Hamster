package subset

type Service interface {
	Create(user string) error
	Run(user string) error
	Stop(user string) error
	Update() error
	InstallServiceToRemoteHost(user string) (err error)
	DeleteServiceToRemoteHost(user string) (err error)
}
