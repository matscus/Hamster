package subset

type Service interface {
	Run(user string) error
	Stop(user string) error
	Update() error
	InsertToDB() (err error)
	InstallServiceToRemoteHost(user string) error
}
