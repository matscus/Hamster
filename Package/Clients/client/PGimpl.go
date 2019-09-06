package client

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	//_ mask PG driver
	"github.com/lib/pq"
	"github.com/matscus/Hamster/Package/Clients/subset"
)

var (
	selectAllService    = "select id,name,host,uri,type,projects from service"
	selectAllScenarios  = "select id, name,test_type,last_modified,gun_type,projects from scenarios;"
	selectAllGenerators = "select id, host from generators"
	selectServiceRunSTR = "select runstr from service where id=$1"
	deleteService       = "delete service where id=$1"
	//selectProjectService = "select id,name,host,uri,type,projects from service where '{$1}' <@ projects"
	db *sql.DB
)

func init() {
	db, _ = sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" dbname="+os.Getenv("POSTGRESDB")+" sslmode=disable")
}

//PGClient - default PGClient struct
type PGClient struct{}

//New - return new interface PGClient
func (c PGClient) New() subset.PGClient {
	var client subset.PGClient
	client = PGClient{}
	return client
}

//GetLastRunID - return last run ID
func (c PGClient) GetLastRunID() (runID int64, err error) {
	db.QueryRow("select max(id) from runs").Scan(&runID)
	return runID, err
}

//GetNewRunID - return new run ID
func (c PGClient) GetNewRunID() (runID int64, err error) {
	db.QueryRow("select max(id) from runs").Scan(&runID)
	return runID + 1, err
}

//GetLastServiceID - return last service id
func (c PGClient) GetLastServiceID() (ID int64, err error) {
	db.QueryRow("select max(id) from service").Scan(&ID)
	return ID, err
}

//GetUserHash - return user password hash
func (c PGClient) GetUserHash(user string) (hash string, err error) {
	q, err := db.Query("select password from users where users=$1", user)
	if err != nil {
		return "", err
	}
	for q.Next() {
		err = q.Scan(&hash)
		if err != nil {
			return "", err
		}
	}
	return hash, err
}

//GetAllUserProject - func to return all users project for front project models
func (c PGClient) GetAllUserProject(user string) (projects []string, err error) {
	if err := db.QueryRow("select projects from users where users=$1", user).Scan(pq.Array(&projects)); err != nil {
		return projects, err
	}
	return projects, nil
}

//UpdateServiceWithOutRunSTR - update service values in database wothout string for run service
func (c PGClient) UpdateServiceWithOutRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string) (err error) {
	projectstr := "{" + strings.Join(projects, ",") + "}"
	_, err = db.Exec("UPDATE service SET name = $1, host = $2, uri= $3,type=$4, projects=$5 where id= $6", name, host, uri, typeTest, projectstr, id)
	if err != nil {
		return err
	}
	return err
}

//UpdateServiceWithOutRunSTR - update service values in database wothout string for run service
func (c PGClient) UpdateServiceWithRunSTR(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error) {
	projectstr := "{" + strings.Join(projects, ",") + "}"
	_, err = db.Exec("UPDATE service SET name = $1, host = $2, uri= $3,type=$4, projects=$5,runstr=$6 where id= $7", name, host, uri, typeTest, projectstr, runSTR, id)
	if err != nil {
		return err
	}
	return err
}

//SetStartTest - insert scenario values to table runs at start scenario
func (c PGClient) SetStartTest(runID string, testName string, testType string) error {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	id, _ := strconv.Atoi(runID)
	_, err := db.Exec("insert into runs (id,run_id,test_name,test_type,start_time,stop_time,status,comment,state) values ($1,$2,$3,$4,to_timestamp($5),to_timestamp(0),'','','')", id, runID, testName, testType, timestamp)
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
func (c PGClient) NewScenario(id int64, name string, typeTest string, gun string, projects []string) (err error) {
	t := time.Now().Unix()
	projectstr := "{" + strings.Join(projects, ",") + "}"
	timestamp := strconv.FormatInt(t, 10)
	_, err = db.Exec("INSERT INTO scenarios (id,name,test_type,last_modified,gun_type,projects) VALUES ($1,$2,$3,to_timestamp($4),$5,$6)", id, name, typeTest, timestamp, gun, projectstr)
	if err != nil {
		return err
	}
	return err
}

//NewService - insert new scenario values to table  scenarios
func (c PGClient) NewService(id int64, name string, host string, uri string, typeTest string, projects []string, runSTR string) (err error) {
	t := time.Now().Unix()
	projectstr := "{" + strings.Join(projects, ",") + "}"
	timestamp := strconv.FormatInt(t, 10)
	_, err = db.Exec("INSERT INTO service (id,name,host,uri,type,projects) VALUES ($1,$2,$3,$4,to_timestamp($5),$6,$7", id, name, host, uri, typeTest, timestamp, projectstr, runSTR)
	if err != nil {
		return err
	}
	return err
}

//UpdateScenario - update scenario values to table  scenarios
func (c PGClient) UpdateScenario(id int64, name string, typeTest string, gun string, projects []string) (err error) {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	projectstr := "{" + strings.Join(projects, ",") + "}"
	fmt.Println("test ", projectstr)
	_, err = db.Exec("UPDATE scenarios SET name = $1,test_type  = $2, last_modified = to_timestamp($3), gun_type= $4, projects=$5 where id= $6", name, typeTest, timestamp, gun, projectstr, id)
	if err != nil {
		return err
	}
	return err
}

//GetAllServices - return all services info
func (c PGClient) GetAllServices() (*[]subset.AllService, error) {
	var rows *sql.Rows
	rows, err := db.Query(selectAllService)
	res := make([]subset.AllService, 0, 200)
	for rows.Next() {
		t := subset.AllService{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.URI, &t.Type, pq.Array(&t.Projects)); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}

//GetAllScenarios - return all scenario info
func (c PGClient) GetAllScenarios() (*[]subset.AllScenario, error) {
	var rows *sql.Rows
	rows, err := db.Query(selectAllScenarios)
	res := make([]subset.AllScenario, 0, 100)
	for rows.Next() {
		t := subset.AllScenario{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.LastModified, &t.Gun, pq.Array(&t.Projects)); err != nil {
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
func (c PGClient) GetAllGenerators() ([][]string, error) {
	var ID, host string
	res := make([][]string, 0, 5)
	var rows *sql.Rows
	rows, err := db.Query(selectAllGenerators)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := make([]string, 2, 2)
		err = rows.Scan(&ID, &host)
		if err != nil {
			return nil, err
		}
		data[0] = ID
		data[1] = host
		res = append(res, data)
	}
	return res, err
}

//GetLastGeneratorsID - return last generator id
func (c PGClient) GetLastGeneratorsID() (ID int64, err error) {
	db.QueryRow("select max(id) from generators").Scan(&ID)
	return ID, err
}

//NewGenerator - insert new generators from database
func (c PGClient) NewGenerator(id int64, host string, projects []string) (err error) {
	projectstr := "{" + strings.Join(projects, ",") + "}"
	_, err = db.Exec("insert into generators(id,host,projects)values($1,$2,$3)", id, host, projectstr)
	if err != nil {
		return err
	}
	return nil
}

//UpdateGenerator - update generator values to table  scenarios
func (c PGClient) UpdateGenerator(id int64, host string, projects []string) (err error) {
	projectstr := "{" + strings.Join(projects, ",") + "}"
	_, err = db.Exec("UPDATE generators SET host = $1,SET projects=$2 where id= $3", host, projectstr, id)
	if err != nil {
		return err
	}
	return nil
}

//GetServiceRunSTR - update generator values to table  scenarios
func (c PGClient) GetServiceRunSTR(id int64) (runSTR string, err error) {
	var rows *sql.Rows
	rows, err = db.Query(selectServiceRunSTR, id)
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
	_, err = db.Exec(deleteService, id)
	if err != nil {
		return err
	}
	return nil
}

//GetProjectServices - func return all service for user project
func (c PGClient) GetProjectServices(project string) (*[]subset.AllService, error) {
	var rows *sql.Rows
	str := "select id,name,host,uri,type,projects from service where '{\"" + project + "\"}' <@ projects"
	rows, err := db.Query(str)
	if err != nil {
		return nil, err
	}
	res := make([]subset.AllService, 0, 200)
	for rows.Next() {
		t := subset.AllService{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Host, &t.URI, &t.Type, pq.Array(&t.Projects)); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}
