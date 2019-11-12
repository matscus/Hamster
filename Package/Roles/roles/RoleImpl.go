package roles

import "github.com/matscus/Hamster/Package/Clients/client"

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Create - create new role and insert data to database
func (r *Role) Create() error {
	return client.PGClient{}.New().NewRole(r.Name)
}

//Update - update role data
func (r *Role) Update() error {
	return client.PGClient{}.New().UpdateRole(r.ID, r.Name)
}

//Delete -delete role
func (r *Role) Delete() error {
	return client.PGClient{}.New().DeleteRole(r.ID)
}
