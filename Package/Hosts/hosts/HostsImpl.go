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

//Create - func for insert new generators, from database
func (h Host) Create() error {
	pgclient := client.PGClient{}.New()
	err := pgclient.NewHost(h.Host, h.Type, h.User, h.Projects)
	if err != nil {
		return err
	}
	return nil
}

//Update - func for udpate generator, from database
func (h Host) Update() error {
	return client.PGClient{}.New().UpdateHost(h.ID, h.Host, h.User, h.Type, h.Projects)
}

//DeleteHost - func for udelete host
func (h Host) Delete() error {
	return client.PGClient{}.New().DeleteHost(h.ID)
}

//IfExist
func (h Host) IfExist() (bool, error) {
	return client.PGClient{}.New().HostIfExist(string(h.ID))
}
