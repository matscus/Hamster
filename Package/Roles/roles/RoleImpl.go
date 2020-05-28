package roles

import (
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
)

type Role struct {
	ID       int64              `json:"id"`
	Name     string             `json:"name"`
	DBClient *postgres.PGClient `json:",omitempty"`
}

//Create - create new role and insert data to database
func (r *Role) Create() error {
	return r.DBClient.NewRole(r.Name)
}

//Update - update role data
func (r *Role) Update() error {
	return r.DBClient.UpdateRole(r.ID, r.Name)
}

//Delete -delete role
func (r *Role) Delete() error {
	return r.DBClient.DeleteRole(r.ID)
}
