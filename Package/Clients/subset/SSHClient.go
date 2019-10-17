package subset

//SSHClient - ssh package impl
type SSHClient interface {
	Run(target string, str string) error
	RunNoWait(target string, str string) error
	Ping(target string) (res bool, err error)
	SCP(target, filePath, destinationPath string) error
	CombinedOutput(target string, str string) ([]byte, error)
	InstallServiceToRemoteHost(serviceType string, name string, target string) (err error)
	DeleteServiceFromRemoteHost(serviceType string, name string, target string) (err error)
}
