package projects

import (
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
)

type Project struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	DBClient *postgres.PGClient
}

//Create - create new project and insert data to database
func (p *Project) Create() error {
	return p.DBClient.NewProject(p.Name)
}

//Update -delete project
func (p *Project) Update() error {
	return p.DBClient.UpdateProject(p.ID, p.Name, p.Status)
}

//Delete -delete project
func (p *Project) Delete() error {
	return p.DBClient.DeleteProject(p.ID)
}
