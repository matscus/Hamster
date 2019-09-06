package scenario

import (
	"github.com/matscus/Hamster/Package/Clients/client"
)

//Scenario - struct for scenario
type Scenario struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Gun          string   `json:"gun"`
	LastModified string   `json:"lastmodified"`
	Projects     []string `json:"projects"`
}

//Update - func for update scenario values in table
func (s *Scenario) Update() (err error) {
	return client.PGClient{}.New().UpdateScenario(s.ID, s.Name, s.Type, s.Gun, s.Projects)
}

//InsertToDB - func for insert new scenario values in table
func (s *Scenario) InsertToDB() (err error) {
	pgclient := client.PGClient{}.New()
	id, err := pgclient.GetLastScenarioID()
	if err != nil {
		return err
	}
	err = pgclient.NewScenario(id, s.Name, s.Type, s.Gun, s.Projects)
	return err
}
