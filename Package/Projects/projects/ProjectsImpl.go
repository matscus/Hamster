package projects

import "github.com/matscus/Hamster/Package/Clients/client"

type Project struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

//Create - create new project and insert data to database
func (p *Project) Create() error {
	return client.PGClient{}.New().NewProject(p.Name)
}

//Update -delete project
func (p *Project) Update() error {
	return client.PGClient{}.New().UpdateProject(p.ID, p.Name, p.Status)
}

//Delete -delete project
func (p *Project) Delete() error {
	return client.PGClient{}.New().DeleteProject(p.ID)
}
