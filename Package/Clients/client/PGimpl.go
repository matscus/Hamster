package client

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	//_ mask PG driver
	"github.com/lib/pq"
	pg "github.com/lib/pq"
	"github.com/matscus/Hamster/Package/Clients/subset"
)

var (
	db *sql.DB
)

func init() {

	db, _ = sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" dbname="+os.Getenv("POSTGRESDB")+" sslmode=disable")
}

//PGClient - default PGClient struct
//type PGClient struct{}

//New - return new interface PGClient
func (c PGClient) New() subset.PGClient {
	var client subset.PGClient
	client = PGClient{}
	return client
}

//GetLastRunID - return last run ID
func (c PGClient) GetLastRunID() (runID int64, err error) {
	err = db.QueryRow("select max(id) from runs").Scan(&runID)
	if err != nil {
		return runID, err
	}
	return runID, nil
}

//GetNewRunID - return new run ID
func (c PGClient) GetNewRunID() (runID int64, err error) {
	err = db.QueryRow("select max(id) from runs").Scan(&runID)
	if err != nil {
		return runID, err
	}
	return runID + 1, err
}

//GetLastServiceID - return last service id
func (c PGClient) GetLastServiceID() (ID int64, err error) {
	err = db.QueryRow("select max(id) from service").Scan(&ID)
	if err != nil {
		return ID, err
	}
	return ID, nil
}

//GetUserHash - return user password hash
func (c PGClient) GetUserHash(user string) (hash string, err error) {
	q, err := db.Query("select password from tUsers where users=$1", user)
	if err != nil {
		return "", err
	}
	q.Close()
	for q.Next() {
		err = q.Scan(&hash)
		if err != nil {
			return "", err
		}
	}
	return hash, err
}

//GetUserPasswordExp - return user password expiration
func (c PGClient) GetUserPasswordExp(user string) (exp string, err error) {
	err = db.QueryRow("select password_expiration from tUsers where users=$1", user).Scan(&exp)
	if err != nil {
		return exp, err
	}
	return exp, nil
}

//UpdateServiceWithOutRunSTR - update service values in database wothout string for run service
func (c PGClient) UpdateServiceWithOutRunSTR(id int64, name string, host string, uri string, typeTest string) (err error) {
	_, err = db.Exec("UPDATE tServices SET name = $1, host = $2, uri= $3,type=$4 where id= $5", name, host, uri, typeTest, id)
	if err != nil {
		return err
	}
	return err
}

//UpdateServiceWithRunSTR - update service values in database wothout string for run service
func (c PGClient) UpdateServiceWithRunSTR(id int64, name string, host string, uri string, typeTest string, runSTR string) (err error) {
	_, err = db.Exec("UPDATE service SET name = $1, host = $2, uri= $3,type=$4, runstr=$5 where id= $6", name, host, uri, typeTest, runSTR, id)
	if err != nil {
		return err
	}
	return err
}

//UpdatetServiceProjects -
func (c PGClient) UpdatetServiceProjects(id int64, projects []string) error {
	_, err := db.Exec("delete from tServiceProjects where user_id=$1", id)
	if err != nil {
		return err
	}
	_, err = db.Query("select tServiceProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//SetStartTest - insert scenario values to table runs at start scenario
func (c PGClient) SetStartTest(testName string, testType string) error {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	_, err := db.Exec("insert into runs (test_name,test_type,start_time,stop_time,status,comment,state) values ($1,$2,$3,$4,to_timestamp($5),to_timestamp(0),'','','')", testName, testType, timestamp)
	if err != err {
		return err
	}
	return nil
}

//SetStopTest - stop runs scenario. Send kill to parent procces a gatling
func (c PGClient) SetStopTest(runID string) error {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	id, _ := strconv.Atoi(runID)
	_, err := db.Exec("update runs set stop_time=to_timestamp($1) where id=$2", timestamp, id)
	if err != nil {
		return err
	}
	return nil
}

//NewScenario - insert new scenario values to table  scenarios
func (c PGClient) NewScenario(name string, typeTest string, gun string, projects string, params string) (err error) {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	_, err = db.Exec("INSERT INTO tScenarios (name,test_type,last_modified,gun_type,project_name,params) VALUES ($1,$2,to_timestamp($3),$4,$5,$6)", name, typeTest, timestamp, gun, projects, params)
	if err != nil {
		return err
	}
	return err
}

//GetScenarioName - insert new scenario values to table  scenarios
func (c PGClient) GetScenarioName(id int64) (res string, err error) {
	if err := db.QueryRow("select name from scenarios where id=$1", id).Scan(&res); err != nil {
		return res, err
	}
	return res, err
}

//CheckScenario - Check scenario, if exist return true, if not exist return fasle
func (c PGClient) CheckScenario(name string, gun string, projects string) (res bool, err error) {
	var tempname string
	err = db.QueryRow("select name from  tScenarios where name=$1 and gun_type=$2 and project_name=$3", name, gun, projects).Scan(&tempname)
	if err != nil {
		return false, err
	}
	if tempname == "" {
		return false, err
	} else {
		return true, err
	}
}

//DeleteScenario - delete scenario from db
func (c PGClient) DeleteScenario(id int64) (err error) {
	_, err = db.Exec("delete from scenarios where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//NewServiceBin - insert new scenario values to table  scenarios
func (c PGClient) NewServiceBin(name string, typeService string, runSTR string, own string, projects []string) (err error) {
	_, err = db.Query("select * from new_bins_function($1,$2,$3,$4,$5)", name, typeService, runSTR, own, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//UpdateServiceBin - func for update bins info
func (c PGClient) UpdateServiceBin(id int64, runSTR string, projects []string) (err error) {
	_, err = db.Exec("update tBins set runstr=$1,  last_modified = now() where id=$2", runSTR, id)
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from tServiceBinProjects where bin_id=$1", id)
	if err != nil {
		return err
	}
	_, err = db.Query("select * from tServiceBinProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//DeleteServiceBin - delete row from table tBins
func (c PGClient) DeleteServiceBin(id int64) (err error) {
	_, err = db.Exec("delete from tBins where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//GetAllServiceBinsNoSort - return all servicebins info
func (c PGClient) GetAllServiceBinsNoSort(projectIDs []string) (*[]subset.AllServiceBinsNoSort, error) {
	var rows *sql.Rows
	rows, err := db.Query("select distinct tbins.id,tbins.name,tbins.type,tbins.runstr,tbins.own,tbins.last_modified from tbins  left join tServiceBinProjects on tbins.id=tServiceBinProjects.bin_id where tServiceBinProjects.id is null or tServiceBinProjects.project_id = any($1)", pg.Array(projectIDs))
	res := make([]subset.AllServiceBinsNoSort, 0, 20)
	if rows == nil {
		return &res, nil
	}
	for rows.Next() {
		t := subset.AllServiceBinsNoSort{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.RunSTR, &t.Owner, &t.LastModified); err != nil {
			return &res, err
		}
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from tServiceBinProjects where bin_id=$1)", t.ID)
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
		res = append(res, t)
	}
	return &res, err
}

//GetAllServiceBinsByOwner - return all servicebins info
func (c PGClient) GetAllServiceBinsByOwner(projectIDs []string) (*[]subset.AllServiceBinsByOwner, error) {
	var rows *sql.Rows
	rows, err := db.Query("select distinct tbins.id,tbins.name,tbins.type,tbins.runstr,tbins.own,tbins.last_modified from tbins  left join tServiceBinProjects on tbins.id=tServiceBinProjects.bin_id where tServiceBinProjects.id is null or tServiceBinProjects.project_id = any($1)", pg.Array(projectIDs))
	tempRes := make([]subset.AllServiceBinsNoSort, 0, 20)
	res := make([]subset.AllServiceBinsByOwner, 0, 20)
	if rows == nil {
		return &res, nil
	}
	for rows.Next() {
		t := subset.AllServiceBinsNoSort{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.RunSTR, &t.Owner, &t.LastModified); err != nil {
			return &res, err
		}
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from tServiceBinProjects where bin_id=$1)", t.ID)
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
		t := subset.AllServiceBinsByOwner{}
		t.Owner = v
		for i := 0; i < len; i++ {
			if v == tempRes[i].Owner {
				t.Services = append(t.Services, subset.ServicesBin{ID: tempRes[i].ID, Name: tempRes[i].Name, Type: tempRes[i].Type, RunSTR: tempRes[i].RunSTR, LastModified: tempRes[i].LastModified, Projects: tempRes[i].Projects})
			}
		}
		res = append(res, t)
	}
	return &res, err
}

//GetAllServiceBinsType - return all servicebins type
func (c PGClient) GetAllServiceBinsType() (*[]subset.AllServiceBinType, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,type_name from tBinsType")
	res := make([]subset.AllServiceBinType, 0, 10)
	if rows == nil {
		return &res, nil
	}
	for rows.Next() {
		t := subset.AllServiceBinType{}
		if err = rows.Scan(&t.ID, &t.Type); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}

//NewService - insert new scenario values to table  scenarios
func (c PGClient) NewService(name string, host string, uri string, typeService string, runSTR string, projects []string, owner string) (err error) {
	_, err = db.Query("select new_service_function", name, host, uri, typeService, runSTR, pg.Array(projects), owner)
	if err != nil {
		return err
	}
	return err
}

//UpdateScenario - update scenario values to table  scenarios
func (c PGClient) UpdateScenario(id int64, name string, typeTest string, gun string, projects string, params string) (err error) {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	_, err = db.Exec("UPDATE tServices SET name = $1,test_type  = $2, last_modified = to_timestamp($3), gun_type= $4, projects=$5,params=$6 where id= $7", name, typeTest, timestamp, gun, projects, params, id)
	if err != nil {
		return err
	}
	return err
}

//GetAllServices - return all services info
func (c PGClient) GetAllServices() (*[]subset.AllService, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,name,host,uri,type from tServices")
	res := make([]subset.AllService, 0, 200)
	for rows.Next() {
		t := subset.AllService{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.Port, &t.Type); err != nil {
			return &res, err
		}
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from tServiceProjects where service_id=$1)", t.ID)
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

//GetAllScenarios - return all scenario info
func (c PGClient) GetAllScenarios() (*[]subset.AllScenario, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id, name,test_type,last_modified,gun_type,project_name,params from tScenarios")
	res := make([]subset.AllScenario, 0, 100)
	for rows.Next() {
		t := subset.AllScenario{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.LastModified, &t.Gun, &t.Projects, &t.TreadGroups); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}

//GetLastScenarioID - func to return lfst scenario id
func (c PGClient) GetLastScenarioID() (id int64, err error) {
	db.QueryRow("select max(id) from scenarios").Scan(&id)
	return id, err
}

//GetAllGenerators - return all generators info.
func (c PGClient) GetAllGenerators() ([]subset.AllHost, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id, ip, host_type from tHosts where host_type='generator'")
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllHost, 0, 20)
	for rows.Next() {
		h := subset.AllHost{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type)
		if err != nil {
			return nil, err
		} else {
			rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
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

//GetLastGeneratorsID - return last generator id
func (c PGClient) GetLastGeneratorsID() (ID int64, err error) {
	db.QueryRow("select max(id) from thosts where host_type=generator").Scan(&ID)
	return ID, err
}

//GetServiceRunSTR - update generator values to table  scenarios
func (c PGClient) GetServiceRunSTR(id int64) (runSTR string, err error) {
	var rows *sql.Rows
	rows, err = db.Query("select runstr from tServices where id=$1", id)
	if err != nil {
		return runSTR, err
	}
	for rows.Next() {
		rows.Scan(&runSTR)
	}
	return runSTR, nil
}

//DeleteService - func for delete row from db
func (c PGClient) DeleteService(id int64) (err error) {
	_, err = db.Exec("delete tServices where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//GetProjectServices - func return all service for user project
func (c PGClient) GetProjectServices(project string) (*[]subset.AllService, error) {
	var projectID int64
	err := db.QueryRow("select id from tProjects where name=$1", project).Scan(&projectID)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select service_id from tServiceProjects where project_id=$1", projectID)
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
	rows, err = db.Query("select id,name,host,port,type from tServices where id =any($1)", pg.Array(serviceIDs))
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllService, 0, 200)
	for rows.Next() {
		t := subset.AllService{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.Port, &t.Type); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}

//GetUserToHost - func return user to host
func (c PGClient) GetUserToHost(ip string) (user string, err error) {
	row := db.QueryRow("select users from tHosts where ip=$1", ip)
	if err = row.Scan(&user); err != nil {
		return "", err
	}
	return user, nil
}

//GetUsersAndHosts - func return ipp host and user for him
func (c PGClient) GetUsersAndHosts() (map[string]string, error) {
	res := make(map[string]string)
	var ip, users string
	rows, err := db.Query("select ip,users from tHosts")
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

//GetAllUsers - func return all users
func (c PGClient) GetAllUsers() ([]subset.AllUser, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,users,role from tUsers")
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllUser, 0, 20)
	for rows.Next() {
		u := subset.AllUser{}
		rows.Scan(&u.ID, &u.User, &u.Role)
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from tuserprojects where user_id=$1)", u.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var project string
			err = rowsProjects.Scan(&project)
			if err != nil {
				return nil, err
			}
			u.Projects = append(u.Projects, project)
		}
		u.Password = "so that the password is not stored in clear text, use https(с) Programmer of Mail.ru Group"
		res = append(res, u)
	}
	return res, nil
}

//NewUser - func create new users
func (c PGClient) NewUser(users string, password string, role string, projects []string) error {
	_, err := db.Query("select new_user_function($1,$2,$3,$4) ", users, password, role, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//UpdateUser - func update  users
func (c PGClient) UpdateUser(id int64, role string) error {
	t := time.Now().Add(2880 * time.Hour).Unix()
	_, err := db.Exec("update tUsers set password_expiration=to_timestamp($1), role=$2 where id=$3", t, role, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser - delete user
func (c PGClient) DeleteUser(id int64) (err error) {
	_, err = db.Exec("delete from tUsers where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//ChangeUserPassword - delete user
func (c PGClient) ChangeUserPassword(id int64, password string) (err error) {
	t := time.Now().Unix()
	_, err = db.Exec("update tUsers set password=$1, password_expiration=to_timestamp($2) where id=$3", password, t, id)
	if err != nil {
		return err
	}
	return nil
}

//UpdatetUserProjects -
func (c PGClient) UpdatetUserProjects(id int64, projects []string) error {
	_, err := db.Exec("delete from tUserProjects where user_id=$1", id)
	if err != nil {
		return err
	}
	_, err = db.Query("select tUserProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//UpdatetHostProjects -
func (c PGClient) UpdatetHostProjects(id int64, projects []string) error {
	_, err := db.Exec("delete from tHostProjects where user_id=$1", id)
	if err != nil {
		return err
	}
	_, err = db.Query("select tHostProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//GetUserRoleAndProject - func to return all users role and  project for front project and role models
func (c PGClient) GetUserRoleAndProject(user string) (role string, projects []string, err error) {
	if err := db.QueryRow("select role,projects from tUsers where users=$1", user).Scan(&role, pq.Array(&projects)); err != nil {
		return role, projects, err
	}
	return role, projects, nil
}

//GetUserIDAndRole - func to return all users role and  project for front project and role models
func (c PGClient) GetUserIDAndRole(user string) (id int64, role string, err error) {
	if err := db.QueryRow("select id,role from tUsers where users=$1", user).Scan(&id, &role); err != nil {
		return 0, role, err
	}
	return id, role, nil
}

//GetUserProjects - func to return all users role and  project for front project and role models
func (c PGClient) GetUserProjects(userID int64) (projects []string, err error) {
	rows, err := db.Query("select name from tProjects where id in(select project_id from tuserprojects where user_id=$1)", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var project string
		err = rows.Scan(&project)
		if err != nil {
			return nil, err
		} else {
			projects = append(projects, project)
		}
	}
	return projects, nil
}

//GetProjectsIDtoString - func to return all users role and  project for front project and role models
func (c PGClient) GetProjectsIDtoString(projects []string) (ids []string, err error) {
	rows, err := db.Query("select id from tprojects where name = any($1)", pg.Array(projects))
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

//GetProjectsIDtoInt - func to return all users role and  project for front project and role models
func (c PGClient) GetProjectsIDtoInt(projects []string) (ids []int, err error) {
	rows, err := db.Query("select id from tprojects where name = any($1)", pg.Array(projects))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

//GetAllHosts - func return all hosts
func (c PGClient) GetAllHosts() ([]subset.AllHost, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,ip,host_type,users from tHosts")
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllHost, 0, 20)
	for rows.Next() {
		h := subset.AllHost{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type, &h.User)
		if err != nil {
			return nil, err
		}
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
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
func (c PGClient) GetAllHostsWithProject(project string) ([]subset.AllHost, error) {
	hostIDs := make([]int64, 0, 0)
	rows, err := db.Query("select host_id from tHostProjects where project_id in(select id from tProjects where name=$1)", project)
	for rows.Next() {
		var hostID int64
		err = rows.Scan(&hostID)
		if err != nil {
			return nil, err
		}
		hostIDs = append(hostIDs, hostID)
	}
	res := make([]subset.AllHost, 0, 20)
	rows, err = db.Query("select id,ip,host_type,users from tHosts where id = any($1)", pg.Array(hostIDs))
	for rows.Next() {
		h := subset.AllHost{}
		err = rows.Scan(&h.ID, &h.Host, &h.Type, &h.User)
		if err != nil {
			return nil, err
		}
		rowsProjects, err := db.Query("select name from tProjects where id in(select project_id from thostprojects where host_id=$1)", h.ID)
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

//NewHost - insert new host from database
func (c PGClient) NewHost(ip string, user string, hostType string, projects []string) (err error) {
	_, err = db.Query("select * from new_host_function($1,$2,$3,$4)", ip, hostType, user, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}

//UpdateHost - update host values to table  scenarios
func (c PGClient) UpdateHost(id int64, ip string, hostType string, user string) (err error) {
	_, err = db.Exec("UPDATE hosts SET ip = $1, host_type=$2, Users=$3  where id= $4", ip, hostType, user, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteHost - delete host
func (c PGClient) DeleteHost(id int64) (err error) {
	_, err = db.Exec("delete from hosts where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

//HostIfExist - chacke host, is exist return true
func (c PGClient) HostIfExist(ip string) (bool, error) {
	var tempHost string
	err := db.QueryRow("select ip from hosts where ip=$1", ip).Scan(&tempHost)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//GetAllProjects - func return all projects
func (c PGClient) GetAllProjects() ([]subset.AllProjects, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,name,status from tProjects")
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllProjects, 0, 20)
	for rows.Next() {
		p := subset.AllProjects{}
		rows.Scan(&p.ID, &p.Name, &p.Status)
		res = append(res, p)
	}
	return res, nil
}

//NewProject - insert new projects from database
func (c PGClient) NewProject(project string) (err error) {
	_, err = db.Exec("insert into tProjects (name,status)values($1,'active')", project)
	if err != nil {
		return err
	}
	return nil
}

//UpdateProject - insert new projects from database
func (c PGClient) UpdateProject(id int64, project string, status string) (err error) {
	_, err = db.Exec("update tProjects SET name=$2,status=$3 where id=$1", id, project, status)
	if err != nil {
		return err
	}
	return nil
}

//DeleteProject - update projects values to table  scenarios
func (c PGClient) DeleteProject(id int64) (err error) {
	_, err = db.Exec("delete from tProjects where id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

//NewRole - insert new role from database
func (c PGClient) NewRole(role string) (err error) {
	_, err = db.Exec("insert into tRole (name)values($1)", role)
	if err != nil {
		return err
	}
	return nil
}

//UpdateRole - update role values
func (c PGClient) UpdateRole(id int64, role string) (err error) {
	_, err = db.Exec("UPDATE tRole SET name = $1  where id= $2", role, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteRole - update role values to table  scenarios
func (c PGClient) DeleteRole(id int64) (err error) {
	_, err = db.Exec("delete from tRole where id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

//GetAllRoles - func return all projects
func (c PGClient) GetAllRoles() ([]subset.AllRoles, error) {
	var rows *sql.Rows
	rows, err := db.Query("select id,name from tRoles")
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllRoles, 0, 20)
	for rows.Next() {
		p := subset.AllRoles{}
		rows.Scan(&p.ID, &p.Name)
		res = append(res, p)
	}
	return res, nil
}
