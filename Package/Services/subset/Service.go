package subset

type Service interface {
	Run() error
	Stop() error
	Update() error
	InsertToDB() (err error)
	InstallServiceToRemoteHost() error
}
