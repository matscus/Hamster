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
	ID                int64    `json:"id"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Gun               string   `json:"gun"`
	LastModified      string   `json:"lastmodified"`
	Projects          []string `json:"projects"`
	TreadGroupsParams string   `json:"params"`
}

type AllGenerator struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	State    string   `json:"state"`
	Projects []string `json:"projects"`
}
