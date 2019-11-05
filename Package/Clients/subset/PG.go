package subset

//PGClient - postgres client impl
type PGClient interface {
	//runs
	GetLastRunID() (runID int64, err error)
	GetNewRunID() (runID int64, err error)
	SetStartTest(testName string, testType string) error
	SetStopTest(runID string) error
	//services
	NewService(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error)
	GetAllServices() (*[]AllService, error)
	GetLastServiceID() (ID int64, err error)
	GetProjectServices(project string) (*[]AllService, error)
	UpdateServiceWithRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error)
	UpdateServiceWithOutRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string) (err error)
	DeleteService(id int64) (err error)
	//scenario
	NewScenario(name string, typeTest string, gun string, projects []string, params string) (err error)
	CheckScenario(name string, gun string, projects []string) (res bool, err error)
	GetAllScenarios() (*[]AllScenario, error)
	GetLastScenarioID() (ID int64, err error)
	GetScenarioName(id int64) (res string, err error)
	UpdateScenario(id int64, name string, typeTest string, gun string, projects []string, params string) (err error)
	DeleteScenario(id int64) (err error)
	//Generators
	GetAllGenerators() (generators [][]string, err error)
	GetAllUserProject(user string) (projects []string, err error)
	GetLastGeneratorsID() (ID int64, err error)
	NewGenerator(id int64, host string, projects []string) (err error)
	UpdateGenerator(id int64, host string, projects []string) (err error)
	GetUserHash(user string) (hash string, err error)
	//users and host
	GetUsersAndHosts() (map[string]string, error)
}
