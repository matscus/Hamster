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

//AllScenario - struct for return all scenario
type AllScenario struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Gun          string   `json:"gun"`
	LastModified string   `json:"lastmodified"`
	Projects     []string `json:"projects"`
	TreadGroups  string   `json:"params"`
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
