package postgres

import (
	"database/sql"

	pg "github.com/lib/pq"
)

//ServiceBinsByOwner - struct for return all servicebins, sort by owner
type ServiceBinsByOwner struct {
	Owner    string       `json:"owner"`
	Services []ServiceBin `json:"services"`
}

//ServiceBin - substruct for AllServiceBinsByOwner
type ServiceBin struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name,omitempty"`
	Type         string   `json:"type,omitempty"`
	RunSTR       string   `json:"runstr,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	LastModified string   `json:"lastmodified,omitempty"`
	Projects     []string `json:"projects,omitempty"`
}

//NewServiceBin - insert new scenario values to table  scenarios
func (c PGClient) NewServiceBin(name string, typeService string, runSTR string, own string, projects []string) (err error) {
	_, err = c.DB.Query("select * from new_bins_function($1,$2,$3,$4,$5)", name, typeService, runSTR, own, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//UpdateServiceBin - func for update bins info
func (c PGClient) UpdateServiceBin(id int64, runSTR string, projects []string) (err error) {
	_, err = c.DB.Exec("update tBins set runstr=$1,  last_modified = now() where id=$2", runSTR, id)
	if err != nil {
		return err
	}
	_, err = c.DB.Exec("delete from tServiceBinProjects where bin_id=$1", id)
	if err != nil {
		return err
	}
	_, err = c.DB.Query("select * from tServiceBinProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//DeleteServiceBin - delete row from table tBins
func (c PGClient) DeleteServiceBin(id int64) (err error) {
	_, err = c.DB.Exec("delete from tBins where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//GetServiceBin return serviceBin info(ONLY ID, Name, Type and runSTR)
func (c PGClient) GetServiceBin(id int64) (*ServiceBin, error) {
	var err error
	res := ServiceBin{}
	err = c.DB.QueryRow("select name,type,runstr from tBins where id=$1", id).Scan(&res.Name, &res.Type, &res.RunSTR)
	return &res, err
}

//GetAllServiceBinsByOwner - return all servicebins info
func (c PGClient) GetAllServiceBinsByOwner(projectIDs []string) (*[]ServiceBinsByOwner, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select distinct tbins.id,tbins.name,tbins.type,tbins.runstr,tbins.own,tbins.last_modified from tbins  left join tServiceBinProjects on tbins.id=tServiceBinProjects.bin_id where tServiceBinProjects.id is null or tServiceBinProjects.project_id = any($1)", pg.Array(projectIDs))
	tempRes := make([]ServiceBin, 0, 20)
	res := make([]ServiceBinsByOwner, 0, 20)
	if rows == nil {
		return &res, nil
	}
	for rows.Next() {
		t := ServiceBin{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.RunSTR, &t.Owner, &t.LastModified); err != nil {
			return &res, err
		}
		rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from tServiceBinProjects where bin_id=$1)", t.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var project string
			err = rowsProjects.Scan(&project)
			if err != nil {
				return nil, err
			}
			t.Projects = append(t.Projects, project)
		}
		tempRes = append(tempRes, t)
	}
	len := len(tempRes)
	allOwner := make([]string, 0, len)
	for i := 0; i < len; i++ {
		allOwner = append(allOwner, tempRes[i].Owner)
	}
	uniqueOwner := make([]string, 0, len)
	for _, v := range allOwner {
		skip := false
		for _, u := range uniqueOwner {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueOwner = append(uniqueOwner, v)
		}
	}
	for _, v := range uniqueOwner {
		t := ServiceBinsByOwner{}
		t.Owner = v
		for i := 0; i < len; i++ {
			if v == tempRes[i].Owner {
				t.Services = append(t.Services, ServiceBin{ID: tempRes[i].ID, Name: tempRes[i].Name, Type: tempRes[i].Type, RunSTR: tempRes[i].RunSTR, LastModified: tempRes[i].LastModified, Projects: tempRes[i].Projects})
			}
		}
		res = append(res, t)
	}
	return &res, err
}

//GetAllServiceBinsType - return all servicebins type
func (c PGClient) GetAllServiceBinsType() (*[]ServiceBin, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id,type_name from tBinsType")
	res := make([]ServiceBin, 0, 10)
	if rows == nil {
		return &res, nil
	}
	for rows.Next() {
		t := ServiceBin{}
		if err = rows.Scan(&t.ID, &t.Type); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}
