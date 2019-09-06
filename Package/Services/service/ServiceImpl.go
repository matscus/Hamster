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
func (s *Service) Run() error {
	client, err := client.SSHClient{}.New()
	if err != nil {
		return err
	}
	fmt.Println(s.RunSTR)
	return client.RunNoWait(s.Host, s.RunSTR)
}

//Stop  - service stop function. performs connection to the host via ssh.
//executes command pkill for service
func (s *Service) Stop() error {
	client, err := client.SSHClient{}.New()
	if err != nil {
		return err
	}
	str := "pkill " + s.Name
	err = client.Run(s.Host, str)
	if err != nil {
	}
	return err
}

//Update - update service info from database
func (s *Service) Update() error {
	if s.RunSTR == "" {
		return client.PGClient{}.New().UpdateServiceWithOutRunSTR(s.ID, s.Name, s.Host, s.URI, s.Type, s.Projects)
	} else {
		return client.PGClient{}.New().UpdateServiceWithRunSTR(s.ID, s.Name, s.Host, s.URI, s.Type, s.Projects, s.RunSTR)
	}

}

//InsertToDB - insert new generator to database
func (s *Service) InsertToDB() error {
	pgclient := client.PGClient{}.New()
	id, err := pgclient.GetLastServiceID()
	if err != nil {
		return err
	}
	return pgclient.NewService(id, s.Name, s.Host, s.URI, s.Type, s.Projects, s.RunSTR)
}

//InstallServiceToRemoteHost - install new service from remote host
func (s *Service) InstallServiceToRemoteHost() (err error) {
	sshclient, err := client.SSHClient{}.New()
	if err != nil {
		return err
	}
	err = sshclient.InstallServiceToRemoteHost(s.Type, s.Name, s.Host)
	if err != nil {
		return err
	}
	pgclient := client.PGClient{}.New()
	id, err := pgclient.GetLastServiceID()
	if err != nil {
		return err
	}
	return pgclient.NewService(id, s.Name, s.Host, s.URI, s.Type, s.Projects, s.RunSTR)
}

//DeleteServiceToRemoteHost - install new service from remote host
func (s *Service) DeleteServiceToRemoteHost() (err error) {
	sshclient, err := client.SSHClient{}.New()
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
