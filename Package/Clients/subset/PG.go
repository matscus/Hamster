package subset

//PGClient - postgres client impl
type PGClient interface {
	//runs
	GetLastRunID() (runID int64, err error)
	GetNewRunID() (runID int64, err error)
	SetStartTest(testName string, testType string) error
	SetStopTest(runID string) error
	//services
	NewService(name string, host string, uri string, type_service string, runSTR string, projects []string) (err error)
	GetAllServices() (*[]AllService, error)
	GetLastServiceID() (ID int64, err error)
	GetProjectServices(project string) (*[]AllService, error)
	UpdateServiceWithRunSTR(id int64, name string, host string, uri string, typeTest string, runSTR string) (err error)
	UpdateServiceWithOutRunSTR(id int64, name string, host string, uri string, typeTest string) error
	UpdatetServiceProjects(id int64, projects []string) error
	DeleteService(id int64) (err error)
	//scenario
	NewScenario(name string, typeTest string, gun string, projects string, params string) error
	CheckScenario(name string, gun string, projects string) (res bool, err error)
	GetAllScenarios() (*[]AllScenario, error)
	GetLastScenarioID() (ID int64, err error)
	GetScenarioName(id int64) (res string, err error)
	UpdateScenario(id int64, name string, typeTest string, gun string, projects string, params string) (err error)
	DeleteScenario(id int64) (err error)
	//Generators
	GetAllGenerators() ([]AllHost, error)
	GetLastGeneratorsID() (ID int64, err error)
	//hosts
	GetAllHosts() ([]AllHost, error)
	HostIfExist(ip string) (bool, error)
	NewHost(ip string, user string, host_type string, projects []string) (err error)
	UpdateHost(id int64, ip string, host_type string, user string) (err error)
	DeleteHost(id int64) (err error)
	UpdatetHostProjects(id int64, projects []string) error
	//users
	GetUserIDAndRole(user string) (id int64, role string, err error)
	GetUserProjects(userID int64) (projects []string, err error)
	GetProjectsIDtoString(projects []string) (ids []string, err error)

	GetUsersAndHosts() (map[string]string, error)
	GetUserHash(user string) (hash string, err error)
	GetUserPasswordExp(user string) (exp string, err error)
	GetAllUsers() ([]AllUser, error)
	NewUser(users string, password string, role string, projects []string) error
	GetUserRoleAndProject(user string) (role string, projects []string, err error)
	UpdateUser(id int64, role string) error
	DeleteUser(user int64) (err error)
	ChangeUserPassword(id int64, password string) (err error)
	UpdatetUserProjects(user int64, projects []string) error
	//projects
	GetAllProjects() ([]AllProjects, error)
	NewProject(project string) (err error)
	DeleteProject(id int64) (err error)

	//role
	NewRole(role string) (err error)
	UpdateRole(id int64, role string) (err error)
	DeleteRole(id int64) (err error)
	GetAllRoles() ([]AllRoles, error)
}
