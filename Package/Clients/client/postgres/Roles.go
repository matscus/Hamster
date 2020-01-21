package postgres

import (
	"database/sql"
)

//Role struct for return all roles
type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//NewRole - insert new role from database
func (c PGClient) NewRole(role string) (err error) {
	_, err = c.DB.Exec("insert into tRole (name)values($1)", role)
	if err != nil {
		return err
	}
	return nil
}

//UpdateRole - update role values
func (c PGClient) UpdateRole(id int64, role string) (err error) {
	_, err = c.DB.Exec("UPDATE tRole SET name = $1  where id= $2", role, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteRole - update role values to table  scenarios
func (c PGClient) DeleteRole(id int64) (err error) {
	_, err = c.DB.Exec("delete from tRole where id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

//GetAllRoles - func return all projects
func (c PGClient) GetAllRoles() ([]Role, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id,name from tRoles")
	if err != nil {
		return nil, err
	}
	res := make([]Role, 0, 20)
	for rows.Next() {
		p := Role{}
		rows.Scan(&p.ID, &p.Name)
		res = append(res, p)
	}
	return res, nil
}
