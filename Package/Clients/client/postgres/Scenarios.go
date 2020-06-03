package postgres

import (
	"database/sql"
	"strconv"
	"time"
)

//Scenario - struct for return all scenario
type Scenario struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Gun          string `json:"gun"`
	LastModified string `json:"lastmodified"`
	Projects     string `json:"projects"`
	TreadGroups  string `json:"params"`
}

//NewScenario - insert new scenario values to table  scenarios
func (c PGClient) NewScenario(name string, typeTest string, gun string, projects string, params string) (err error) {
	_, err = c.DB.Exec("INSERT INTO tScenarios (name,test_type,last_modified,gun_type,project_name,params) VALUES ($1,$2,now(),$3,$4,$5)", name, typeTest, gun, projects, params)
	return err
}

//GetScenarioName - insert new scenario values to table  scenarios
func (c PGClient) GetScenarioName(id int64) (res string, err error) {
	err = c.DB.QueryRow("select name from scenarios where id=$1", id).Scan(&res)
	return res, err
}

//GetAllScenarios - return all scenario info
func (c PGClient) GetAllScenarios() (*[]Scenario, error) {
	var rows *sql.Rows
	rows, err := c.DB.Query("select id, name,test_type,last_modified,gun_type,project_name,params from tScenarios")
	res := make([]Scenario, 0, 100)
	for rows.Next() {
		t := Scenario{}
		if err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.LastModified, &t.Gun, &t.Projects, &t.TreadGroups); err != nil {
			return &res, err
		}
		res = append(res, t)
	}
	return &res, err
}

//GetLastScenarioID - func to return lfst scenario id
func (c PGClient) GetLastScenarioID() (id int64, err error) {
	err = c.DB.QueryRow("select max(id) from scenarios").Scan(&id)
	return id, err
}

//SetStopTest - stop runs scenario. Send kill to parent procces a gatling
func (c PGClient) SetStopTest(runID string) error {
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	id, _ := strconv.Atoi(runID)
	_, err := c.DB.Exec("update truns set stop_time=to_timestamp($1) where id=$2", timestamp, id)
	if err != nil {
		return err
	}
	return nil
}

//GetNewRunID - return new run ID
func (c PGClient) GetNewRunID() (runID int64, err error) {
	rows, err := c.DB.Query("select max(id) from truns")
	if err != nil {
		return runID, err
	}
	for rows.Next() {
		err := rows.Scan(&runID)
		if err != nil {
			runID = 1
			return runID, nil
		}
	}
	return runID + 1, err
}

//CheckScenario - Check scenario, if exist return true, if not exist return fasle
func (c PGClient) CheckScenario(name string, gun string, projects string) (res bool, err error) {
	var tempname string
	err = c.DB.QueryRow("select name from  tScenarios where name=$1 and gun_type=$2 and project_name=$3", name, gun, projects).Scan(&tempname)
	if err != nil {
		return false, err
	}
	if tempname == "" {
		return false, err
	} else {
		return true, err
	}
}

//SetStartTest - insert scenario values to table runs at start scenario
func (c PGClient) SetStartTest(testName string, testType string) (err error) {
	_, err = c.DB.Exec("insert into truns (test_name,test_type,start_time,stop_time,status,comment,state) values ($1,$2,now(),to_timestamp(0),'','','')", testName, testType)
	return err
}

//UpdateScenario - update scenario values to table  scenarios
func (c PGClient) UpdateScenario(id int64, name string, typeTest string, gun string, projects string, params string) (err error) {
	_, err = c.DB.Exec("UPDATE tServices SET name = $1,test_type  = $2, last_modified = now(), gun_type= $3, projects=$4,params=$5 where id= $6", name, typeTest, gun, projects, params, id)
	return err
}

//DeleteScenario - delete scenario from db
func (c PGClient) DeleteScenario(id int64) (err error) {
	_, err = c.DB.Exec("delete from tscenarios where id=$1", id)
	return err
}
