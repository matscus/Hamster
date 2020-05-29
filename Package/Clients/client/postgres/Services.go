package postgres

import (
	"database/sql"

	pg "github.com/lib/pq"
)

//Service - struct for return all service
type Service struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Owner    string   `json:"owner"`
	Type     string   `json:"type"`
	RunSTR   string   `json:"runstr"`
	BinsID   int64    `json:"binsid"`
	Status   string   `json:"status,omitempty"`
	Projects []string `json:"projects"`
}

//NewService - insert new scenario values to table  scenarios
func (c PGClient) NewService(name string, binsIB int64, host string, port int, typeService string, runSTR string, projects []string, owner string) (err error) {
	_, err = c.DB.Query("select new_service_function($1,$2,$3,$4,$5,$6,$7,$8)", name, binsIB, host, port, typeService, runSTR, pg.Array(projects), owner)
	return err
}

//UpdateService - update service values in database wothout string for run service
func (c PGClient) UpdateService(id int64, port int, runSTR string) (err error) {
	_, err = c.DB.Exec("UPDATE service SET port= $1,runstr=$2,last_modified=now() where id= $3", port, runSTR, id)
	return err
}

//DeleteService - func for delete row from db
func (c PGClient) DeleteService(id int64) (err error) {
	_, err = c.DB.Exec("delete tServices where id=$1", id)
	return nil
}

//GetAllServices - return all services info
func (c PGClient) GetAllServices() (*[]Service, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id,name,host,port,type,runstr,owner,binsid from tServices")
	res := make([]Service, 0, 200)
	for rows.Next() {
		t := Service{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.Port, &t.Type, &t.RunSTR, &t.Owner, &t.BinsID); err != nil {
			return &res, err
		}
		rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from tServiceProjects where service_id=$1)", t.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var tempProject string
			err := rowsProjects.Scan(&tempProject)
			if err != nil {
				return nil, err
			}
			t.Projects = append(t.Projects, tempProject)
		}
		res = append(res, t)
	}
	return &res, err
}

//GetService - return ONLY type,name and runSTR for service
func (c PGClient) GetService(id int64) (res *Service, err error) {
	err = c.DB.QueryRow("select type,name,runstr from tServices where id=$1", id).Scan(&res.Type, &res.Name, &res.RunSTR)
	return res, err
}

//GetServicesByProject - func return all service for user project
func (c PGClient) GetServicesByProject(projects []string) (*[]Service, error) {
	var projectsID []int64
	rows, err := c.DB.Query("select id from tProjects where name = any($1)", pg.Array(projects))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tempID int64
		err := rows.Scan(&tempID)
		if err != nil {
			return nil, err
		}
		projectsID = append(projectsID, tempID)
	}

	rows, err = c.DB.Query("select service_id from tServiceProjects where project_id=any($1)", pg.Array(projectsID))
	if err != nil {
		return nil, err
	}
	serviceIDs := make([]int64, 0, 10)
	for rows.Next() {
		var tempID int64
		err := rows.Scan(&tempID)
		if err != nil {
			return nil, err
		}
		serviceIDs = append(serviceIDs, tempID)
	}
	rows, err = c.DB.Query("select id,name,host,port,type from tServices where id =any($1)", pg.Array(serviceIDs))
	if err != nil {
		return nil, err
	}
	res := make([]Service, 0, 200)
	for rows.Next() {
		t := Service{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.Port, &t.Type); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}
