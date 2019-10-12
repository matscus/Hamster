package scenario

import (
	"fmt"

	"github.com/matscus/Hamster/Package/Clients/client"
)

//Scenario - struct for scenario
type Scenario struct {
	ID                int64               `json:"id"`
	Name              string              `json:"name"`
	Type              string              `json:"type"`
	Gun               string              `json:"gun"`
	LastModified      string              `json:"lastmodified"`
	Projects          []string            `json:"projects"`
	TreadGroupsParams []TreadGroupsParams `json:"TreadGroupsParams"`
}

//Update - func for update scenario values in table
func (s *Scenario) Update() (err error) {
	return client.PGClient{}.New().UpdateScenario(s.ID, s.Name, s.Type, s.Gun, s.Projects, treadGroupsParamsToSliceString(s.TreadGroupsParams))
}

//InsertToDB - func for insert new scenario values in table
func (s *Scenario) InsertToDB() (err error) {
	pgclient := client.PGClient{}.New()
	err = pgclient.NewScenario(s.Name, s.Type, s.Gun, s.Projects, treadGroupsParamsToSliceString(s.TreadGroupsParams))
	return err
}

//GetNameForID - func for insert new scenario values in table
func (s *Scenario) GetNameForID() (res string, err error) {
	pgclient := client.PGClient{}.New()
	res, err = pgclient.GetScenarioName(s.ID)
	if err != nil {
		return "", err
	}
	return res, err
}

//DeleteScenario - func for delete secenario(rows db and files)
func (s *Scenario) DeleteScenario() (err error) {
	pgclient := client.PGClient{}.New()
	err = pgclient.DeleteScenario(s.ID)
	if err != nil {
		return err
	}
	return nil
}

//CheckScenario - func for delete secenario(rows db and files)
func (s *Scenario) CheckScenario() (res bool, err error) {
	res, err = client.PGClient{}.New().CheckScenario(s.Name, s.Gun, s.Projects)
	if err != nil {
		return res, err
	}
	return res, nil
}

func treadGroupsParamsToSliceString(params []TreadGroupsParams) []string {
	l := len(params)
	res := make([]string, 0, l)
	for i := 0; i < l; i++ {
		res = append(res, fmt.Sprint(params[i]))
	}
	return res
}
