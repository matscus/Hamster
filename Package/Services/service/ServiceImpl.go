package service

import (
	"io"
	"os"
	"sync"

	"github.com/matscus/Hamster/Package/Clients/client"
)

//Service - service structure, contains the service name, installation host,
//launch status, web address and mutex for blocking parallel changes in the launch status by different threads
type Service struct {
	Mutex    *sync.Mutex `json:",omitempty"`
	ID       int64       `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Host     string      `json:"host,omitempty"`
	Status   string      `json:"status,omitempty"`
	URI      string      `json:"uri,omitempty"`
	Type     string      `json:"type,omitempty"`
	Projects []string    `json:"projects,omitempty"`
	RunSTR   string      `json:"runstr,omitempty"`
	Own      string      `json:"owner,omitempty"`
}

//Run  - service run function. performs connection to the host via ssh.
func (s *Service) Run(user string) error {
	client, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	return client.RunNoWait(s.Host, s.RunSTR)
}

//Stop  - service stop function. performs connection to the host via ssh.
//executes command pkill for service
func (s *Service) Stop(user string) error {
	client, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	str := "pkill " + s.Name
	err = client.Run(s.Host, str)
	if err != nil {
	}
	return err
}

//CreateBin - create bin from bins dir and insert data from tBins
func (s *Service) CreateBin(f io.Reader, own string) error {
	newFile := os.Getenv("BINSDIR") + "/" + s.Type + "/" + s.Name + ".zip"
	file, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		return err
	}
	pgclient := client.PGClient{}.New()
	projectsID, err := pgclient.GetProjectsIDtoString(s.Projects)
	if err != nil {
		return err
	}
	return client.PGClient{}.New().NewServiceBin(s.Name, s.Type, s.RunSTR, own, projectsID)
}

//UpdateBin - update bins and data from tBins
func (s *Service) UpdateBin(own string) error {
	return client.PGClient{}.New().UpdateServiceBin(s.ID, s.Name, s.Type, s.RunSTR, own)
}

//DeleteBin - delete bins and data from tBins
func (s *Service) DeleteBin() error {
	err := os.Remove(os.Getenv("BINSDIR") + "/" + s.Type + "/" + s.Name + ".zip")
	if err != nil {
		return err
	}
	return client.PGClient{}.New().DeleteServiceBin(s.ID)
}

//Create - insert new service to database
func (s *Service) Create(owner string) error {
	return client.PGClient{}.New().NewService(s.Name, s.Host, s.URI, s.Type, s.RunSTR, s.Projects, owner)
}

//Update - update service info from database
func (s *Service) Update() error {
	client := client.PGClient{}.New()
	if s.RunSTR == "" {
		client.UpdateServiceWithOutRunSTR(s.ID, s.Name, s.Host, s.URI, s.Type)
		return client.UpdatetServiceProjects(s.ID, s.Projects)
	} else {
		client.UpdateServiceWithRunSTR(s.ID, s.Name, s.Host, s.URI, s.Type, s.RunSTR)
		return client.UpdatetServiceProjects(s.ID, s.Projects)
	}

}

//InstallServiceToRemoteHost - install new service from remote host
func (s *Service) InstallServiceToRemoteHost(user string) (err error) {
	sshclient, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	err = sshclient.InstallServiceToRemoteHost(s.Type, s.Name, s.Host)
	if err != nil {
		return err
	}
	pgclient := client.PGClient{}.New()
	err = pgclient.NewService(s.Name, s.Host, s.URI, s.Type, s.RunSTR, s.Projects, user)
	if err != nil {
		return err
	}
	return nil
}

//DeleteServiceToRemoteHost - install new service from remote host
func (s *Service) DeleteServiceToRemoteHost(user string) (err error) {
	sshclient, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	err = sshclient.DeleteServiceFromRemoteHost(s.Type, s.Name, s.Host)
	if err != nil {
		return err
	}
	err = client.PGClient{}.New().DeleteService(s.ID)
	if err != nil {
		return err
	}
	return nil
}
