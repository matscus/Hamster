package subset

//AllService - struct for return all service
type AllService struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Host     string   `json:"host"`
	URI      string   `json:"uri"`
	Type     string   `json:"type"`
	Projects []string `json:"projects"`
}

//AllServiceBinsNoSort - struct for return all servicebins, no sotr
type AllServiceBinsNoSort struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	RunSTR       string   `json:"runstr"`
	Owner        string   `json:"owner"`
	LastModified string   `json:"lastmodified"`
	Projects     []string `json:"projects"`
}

//AllServiceBinsByOwner - struct for return all servicebins, sort by owner
type AllServiceBinsByOwner struct {
	Owner    string        `json:"owner"`
	Services []ServicesBin `json:"services"`
}

//ServicesBin - substruct for AllServiceBinsByOwner
type ServicesBin struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	RunSTR       string   `json:"runstr"`
	LastModified string   `json:"lastmodified"`
	Projects     []string `json:"projects"`
}

//AllServiceBinType - struct for return all type  servicebins
type AllServiceBinType struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

//AllScenario - struct for return all scenario
type AllScenario struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Gun          string `json:"gun"`
	LastModified string `json:"lastmodified"`
	Projects     string `json:"projects"`
	TreadGroups  string `json:"params"`
}

//AllGenerator - struct for response all generators data
type AllGenerator struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	State    string   `json:"state"`
	Projects []string `json:"projects"`
}

//AllUser - struct for response all users data
type AllUser struct {
	ID       int64    `json:id`
	User     string   `json:user`
	Password string   `json:password,omitempty`
	Role     string   `json:role`
	Projects []string `json:"projects"`
}

//AllHost - struct for response all hosts data
type AllHost struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	Type     string   `json:"type`
	User     string   `json:user`
	State    string   `json:"state,omitempty"`
	Projects []string `json:"projects"`
}

//AllProjects - struct for response all Projects data
type AllProjects struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

//AllRoles struct for return all roles
type AllRoles struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
