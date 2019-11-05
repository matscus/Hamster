package hosts

import (
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Hosts/subset"
)

//Host - Generators struct
type Host struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	Type     string   `json:"type`
	User     string   `json:user`
	State    string   `json:"state"`
	Projects []string `json:"projects"`
}

//New - func to return new host interface
func (c Host) New() subset.Host {
	var host subset.Host
	host = Host{}
	return host
}

//InsertToDB - func for insert new generators, from database
func (g Host) InsertToDB() error {
	pgclient := client.PGClient{}.New()
	err := pgclient.NewHost(g.Host, g.Type, g.User, g.Projects)
	if err != nil {
		return err
	}
	return nil
}

//UpdateHost - func for udpate generator, from database
func (g Host) UpdateHost() error {
	return client.PGClient{}.New().UpdateHost(g.ID, g.Host, g.Type, g.User, g.Projects)
}
