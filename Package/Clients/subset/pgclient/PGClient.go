package pgclient

import "github.com/matscus/Hamster/Package/Clients/client/postgres"

//PGClient - postgres client impl
type PGClient interface {
	//USERS
	//////Create or Update or Delele
	NewUser(users string, password string, role string, projects []string) error
	UpdateUser(id string, role string, projects []string) error
	DeleteUser(user int64) (err error)
	//////Auth func
	GetUserHash(user string) (hash string, err error)
	ChangeUserPassword(id int64, password string) (err error)
	GetUserPasswordExp(user string) (exp string, err error)
	//////Get users data
	GetAllUsers() (users []postgres.User, err error)
	GetUserRoleAndProjects(user string) (role string, projects []string, err error)
	GetUserIDAndRole(user string) (id int64, role string, err error)
	GetUserProjects(userID int64) (projects []string, err error)

	//PROJECTS
	NewProject(project string) (err error)
	UpdateProject(id int64, project string, status string) (err error)
	DeleteProject(id int64) (err error)
	//////Get project data
	GetAllProjects() ([]postgres.Project, error)
	GetProjectsIDtoString(projects []string) (ids []string, err error)

	//ROLE
	NewRole(role string) (err error)
	UpdateRole(id int64, role string) (err error)
	DeleteRole(id int64) (err error)
	//////Get Role data
	GetAllRoles() ([]postgres.Role, error)

	//HOST
	NewHost(ip string, user string, hostType string, projects []string) (err error)
	UpdateHost(id int64, ip string, hostType string, user string) (err error)
	DeleteHost(id int64) (err error)
	UpdatetHostProjects(id int64, projects []string) error
	//////Get Host data
	GetAllHosts() ([]postgres.Host, error)
	GetAllHostsWithProject(project string) ([]postgres.Host, error)
	HostIfExist(ip string) (bool, error)
	GetUsersAndHosts() (map[string]string, error)
	GetUserToHost(ip string) (user string, err error)
	//////Get Host data if host is generator
	GetAllGenerators() ([]postgres.Host, error)
	//GetLastGeneratorsID() (ID int64, err error)

	//SERVICEBIN
	NewServiceBin(name string, typeService string, runSTR string, own string, projects []string) (err error)
	UpdateServiceBin(id int64, runSTR string, projects []string) (err error)
	DeleteServiceBin(id int64) (err error)
	//////Get ServiceBin data
	GetServiceBin(id int64) (*postgres.ServiceBin, error)
	GetAllServiceBinsByOwner(projectIDs []string) (*[]postgres.ServiceBinsByOwner, error)
	GetAllServiceBinsType() (*[]postgres.ServiceBin, error)

	//SERVICE
	NewService(name string, binsIB int64, host string, port int, typeService string, runSTR string, projects []string, owner string) (err error)
	UpdateService(id int64, port int, runSTR string) (err error)
	DeleteService(id int64) (err error)
	/////Get service data
	GetAllServices() (*[]postgres.Service, error)
	GetService(id int64) (*postgres.Service, error)

	//SCENARIOS
	NewScenario(name string, typeTest string, gun string, projects string, params string) (err error)
	UpdateScenario(id int64, name string, typeTest string, gun string, projects string, params string) (err error)
	DeleteScenario(id int64) (err error)
	//////Get Scenario Data
	CheckScenario(name string, gun string, projects string) (res bool, err error)
	GetScenarioName(id int64) (res string, err error)
	GetNewRunID() (runID int64, err error)
	GetLastScenarioID() (id int64, err error)
	GetAllScenarios() (*[]postgres.Scenario, error)
	//////Manage scenario
	SetStartTest(testName string, testType string) (err error)
	SetStopTest(runID string) error
}
