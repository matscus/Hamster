package hosts

import (
	"os/exec"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Hosts/subset"
)

//Host - Generators struct
type Host struct {
	ID       int64    `json:"id"`
	Host     string   `json:"host"`
	Type     string   `json:"type,omitempty"`
	User     string   `json:"user,omitempty"`
	Password string   `json:"password,omitempty"`
	State    string   `json:"state"`
	Projects []string `json:"projects"`
}

//New - func to return new host interface
func (h Host) New() subset.Host {
	var host subset.Host
	host = Host{}
	return host
}

//Create - func for insert new host, from database
func (h Host) Create() error {
	cmd := exec.Command("sshpass", "-p", h.Password, "ssh-copy-id", h.User+"@"+h.Host)
	err := cmd.Run()
	if err != nil {
		return err
	}
	pgclient := client.PGClient{}.New()
	projectIDs, err := pgclient.GetProjectsIDtoString(h.Projects)
	if err != nil {
		return err
	}
	err = pgclient.NewHost(h.Host, h.Type, h.User, projectIDs)
	if err != nil {
		return err
	}
	return nil
}

//Update - func for udpate generator, from database
func (h Host) Update() error {
	client := client.PGClient{}.New()
	err := client.UpdateHost(h.ID, h.Host, h.User, h.Type)
	if err != nil {
		return err
	}
	projectsID, err := client.GetProjectsIDtoString(h.Projects)
	if err != nil {
		return err
	}
	return client.UpdatetHostProjects(h.ID, projectsID)
}

//Delete - func for udelete host
func (h Host) Delete() error {
	return client.PGClient{}.New().DeleteHost(h.ID)
}

//IfExist -
func (h Host) IfExist() (bool, error) {
	return client.PGClient{}.New().HostIfExist(string(h.ID))
}
