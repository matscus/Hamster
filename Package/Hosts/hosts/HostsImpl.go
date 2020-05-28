package hosts

import (
	"os/exec"

	"github.com/matscus/Hamster/Package/Clients/client/postgres"
	"github.com/matscus/Hamster/Package/Hosts/subset"
)

//Host - Generators struct
type Host struct {
	ID       int64              `json:"id"`
	Host     string             `json:"host"`
	Type     string             `json:"type,omitempty"`
	User     string             `json:"user,omitempty"`
	Password string             `json:"password,omitempty"`
	State    string             `json:"state"`
	Projects []string           `json:"projects"`
	DBClient *postgres.PGClient `json:",omitempty"`
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
	projectIDs, err := h.DBClient.GetProjectsIDtoString(h.Projects)
	if err != nil {
		return err
	}
	err = h.DBClient.NewHost(h.Host, h.Type, h.User, projectIDs)
	if err != nil {
		return err
	}
	return nil
}

//Update - func for udpate generator, from database
func (h Host) Update() error {
	err := h.DBClient.UpdateHost(h.ID, h.Host, h.User, h.Type)
	if err != nil {
		return err
	}
	projectsID, err := h.DBClient.GetProjectsIDtoString(h.Projects)
	if err != nil {
		return err
	}
	return h.DBClient.UpdatetHostProjects(h.ID, projectsID)
}

//Delete - func for udelete host
func (h Host) Delete() error {
	return h.DBClient.DeleteHost(h.ID)
}

//IfExist -
func (h Host) IfExist() (bool, error) {
	return h.DBClient.HostIfExist(string(h.ID))
}
