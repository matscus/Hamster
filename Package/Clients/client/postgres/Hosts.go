package postgres

import (
	"database/sql"

	pg "github.com/lib/pq"
)

//Host - struct for response all hosts data
type Host struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	Type     string   `json:"type`
	User     string   `json:user`
	State    string   `json:"state,omitempty"`
	Projects []string `json:"projects"`
}

//NewHost - insert new host from database
func (c PGClient) NewHost(ip string, user string, hostType string, projects []string) (err error) {
	_, err = c.DB.Query("select * from new_host_function($1,$2,$3,$4)", ip, hostType, user, pg.Array(projects))
	return err
}

//UpdateHost - update host values to table  scenarios
func (c PGClient) UpdateHost(id int64, ip string, hostType string, user string) (err error) {
	_, err = c.DB.Exec("UPDATE hosts SET ip = $1, host_type=$2, Users=$3  where id= $4", ip, hostType, user, id)
	return err
}

//DeleteHost - delete host
func (c PGClient) DeleteHost(id int64) (err error) {
	_, err = c.DB.Exec("delete from hosts where id=$1", id)
	return err
}

//GetAllHosts - func return all hosts
func (c PGClient) GetAllHosts() ([]Host, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id,ip,host_type,users from tHosts")
	if err != nil {
		return nil, err
	}
	res := make([]Host, 0, 20)
	for rows.Next() {
		h := Host{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type, &h.User)
		if err != nil {
			return nil, err
		}
		rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var project string
			err = rowsProjects.Scan(&project)
			if err != nil {
				return nil, err
			}
			h.Projects = append(h.Projects, project)
		}

		res = append(res, h)
	}
	return res, nil
}

//GetAllHostsWithProject - func return all hosts
func (c PGClient) GetAllHostsWithProject(project string) ([]Host, error) {
	hostIDs := make([]int64, 0, 0)
	rows, err := c.DB.Query("select host_id from tHostProjects where project_id in(select id from tProjects where name=$1)", project)
	for rows.Next() {
		var hostID int64
		err = rows.Scan(&hostID)
		if err != nil {
			return nil, err
		}
		hostIDs = append(hostIDs, hostID)
	}
	res := make([]Host, 0, 20)
	rows, err = c.DB.Query("select id,ip,host_type,users from tHosts where id = any($1)", pg.Array(hostIDs))
	for rows.Next() {
		h := Host{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type, &h.User)
		if err != nil {
			return nil, err
		}
		rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var project string
			err = rowsProjects.Scan(&project)
			if err != nil {
				return nil, err
			}
			h.Projects = append(h.Projects, project)
		}
		res = append(res, h)
	}
	return res, nil
}

//HostIfExist - chacke host, is exist return true
func (c PGClient) HostIfExist(ip string) (bool, error) {
	var tempHost string
	err := c.DB.QueryRow("select ip from hosts where ip=$1", ip).Scan(&tempHost)
	if err != nil {
		return false, err
	}
	return true, nil
}

//GetUsersAndHosts - func return ipp host and user for him
func (c PGClient) GetUsersAndHosts() (map[string]string, error) {
	res := make(map[string]string)
	var ip, users string
	rows, err := c.DB.Query("select ip,users from tHosts")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&ip, &users); err != nil {
			return nil, err
		}
		res[ip] = users
	}
	return res, nil
}

//GetUserToHost - func return user to host
func (c PGClient) GetUserToHost(ip string) (user string, err error) {
	row := c.DB.QueryRow("select users from tHosts where ip=$1", ip)
	if err = row.Scan(&user); err != nil {
		return "", err
	}
	return user, nil
}

//UpdatetHostProjects -
func (c PGClient) UpdatetHostProjects(id int64, projects []string) error {
	_, err := c.DB.Exec("delete from tHostProjects where user_id=$1", id)
	if err != nil {
		return err
	}
	_, err = c.DB.Query("select tHostProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//GetAllGenerators - return all generators info.
func (c PGClient) GetAllGenerators() ([]Host, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id, ip, host_type from tHosts where host_type='generator'")
	if err != nil {
		return nil, err
	}
	res := make([]Host, 0, 20)
	for rows.Next() {
		h := Host{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type)
		if err != nil {
			return nil, err
		} else {
			rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
			if err != nil {
				return nil, err
			}
			for rowsProjects.Next() {
				var project string
				err = rowsProjects.Scan(&project)
				if err != nil {
					return nil, err
				}
				h.Projects = append(h.Projects, project)
			}
		}
		res = append(res, h)
	}
	return res, nil
}
