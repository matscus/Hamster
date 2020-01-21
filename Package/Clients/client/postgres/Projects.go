package postgres

import (
	"database/sql"
	"strconv"

	pg "github.com/lib/pq"
)

//Project - temp struct for response all Projects data
type Project struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

//NewProject - insert new projects from database
func (c PGClient) NewProject(project string) (err error) {
	_, err = c.DB.Exec("insert into tProjects (name,status)values($1,'active')", project)
	return err
}

//UpdateProject - insert new projects from database
func (c PGClient) UpdateProject(id int64, project string, status string) (err error) {
	_, err = c.DB.Exec("update tProjects SET name=$2,status=$3 where id=$1", id, project, status)
	return err
}

//DeleteProject - update projects values to table  scenarios
func (c PGClient) DeleteProject(id int64) (err error) {
	_, err = c.DB.Exec("delete from tProjects where id= $1", id)
	return err
}

//GetAllProjects - func return all projects
func (c PGClient) GetAllProjects() ([]Project, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id,name,status from tProjects")
	if err != nil {
		return nil, err
	}
	res := make([]Project, 0, 20)
	for rows.Next() {
		p := Project{}
		rows.Scan(&p.ID, &p.Name, &p.Status)
		res = append(res, p)
	}
	return res, nil
}

//GetProjectsIDtoString - func to return all users role and  project for front project and role models
func (c PGClient) GetProjectsIDtoString(projects []string) (ids []string, err error) {
	rows, err := c.DB.Query("select id from tprojects where name = any($1)", pg.Array(projects))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, strconv.Itoa(id))
	}
	return ids, nil
}
