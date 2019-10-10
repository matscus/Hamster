package subset

type PGClient interface {
	GetLastRunID() (runId int64, err error)
	GetNewRunID() (runID int64, err error)
	GetLastServiceID() (ID int64, err error)
	GetLastScenarioID() (ID int64, err error)
	GetUserHash(user string) (hash string, err error)
	GetAllServices() (*[]AllService, error)
	UpdateServiceWithRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error)
	UpdateServiceWithOutRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string) (err error)
	SetStartTest(runID string, testName string, testType string) error
	SetStopTest(runID string) error
	GetAllScenarios() (*[]AllScenario, error)
	NewScenario(name string, typeTest string, gun string, projects []string) error
	CheckScenario(name string, gun string, projects []string) (res bool, err error)
	GetScenarioName(id int64) (res string, err error)
	DeleteScenario(id int64) (err error)
	NewService(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error)
	UpdateScenario(id int64, name string, typeTest string, gun string, projects []string) error
	GetAllGenerators() (generators [][]string, err error)
	GetAllUserProject(user string) (projects []string, err error)
	GetLastGeneratorsID() (ID int64, err error)
	NewGenerator(id int64, host string, projects []string) (err error)
	UpdateGenerator(id int64, host string, projects []string) (err error)
	DeleteService(id int64) (err error)
	GetProjectServices(project string) (*[]AllService, error)
}
