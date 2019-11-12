package service

import (
	"fmt"
	"sync"

	"github.com/matscus/Hamster/Package/Clients/client"
)

//Service - service structure, contains the service name, installation host,
//launch status, web address and mutex for blocking parallel changes in the launch status by different threads
type Service struct {
	Mutex    *sync.Mutex `json:",omitempty"`
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	Host     string      `json:"host"`
	Status   string      `json:"status"`
	URI      string      `json:"uri"`
	Type     string      `json:"type"`
	Projects []string    `json:"projects"`
	RunSTR   string      `json:"runstr,omitempty"`
}

//Run  - service run function. performs connection to the host via ssh.
func (s *Service) Run(user string) error {
	client, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	fmt.Println(s.RunSTR)
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

//InsertToDB - insert new generator to database
func (s *Service) Create() error {
	return client.PGClient{}.New().NewService(s.Name, s.Host, s.URI, s.Type, s.RunSTR, s.Projects)
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
	err = pgclient.NewService(s.Name, s.Host, s.URI, s.Type, s.RunSTR, s.Projects)
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
