package generators

import (
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Generators/subset"
)

//Generator - Generators struct
type Generator struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	State    string   `json:"state"`
	Projects []string `json:"projects"`
}

//New - func to return new generators interface
func (c Generator) New() subset.Generator {
	var generator subset.Generator
	generator = Generator{}
	return generator
}

//InsertToDB - func for insert new generators, from database
func (g Generator) InsertToDB() error {
	pgclient := client.PGClient{}.New()
	id, err := pgclient.GetLastGeneratorsID()
	if err != nil {
		return err
	}
	err = pgclient.NewGenerator(id, g.Host, g.Projects)
	if err != nil {
		return err
	}
	return nil
}

//UpdateGenerator - func for udpate generator, from database
func (g Generator) UpdateGenerator() error {
	return client.PGClient{}.New().UpdateGenerator(g.ID, g.Host, g.Projects)
}
