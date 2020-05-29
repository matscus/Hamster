package service

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
)

//Service - service structure, contains the service name, installation host,
//launch status, web address and mutex for blocking parallel changes in the launch status by different threads
type Service struct {
	Mutex    *sync.Mutex        `json:",omitempty"`
	ID       int64              `json:"id,omitempty"`
	BinsID   int64              `json:"binid,omitempty"`
	Name     string             `json:"name"`
	Host     string             `json:"host,omitempty"`
	Status   string             `json:"status,omitempty"`
	Port     int                `json:"port,omitempty"`
	Type     string             `json:"type,omitempty"`
	Projects []string           `json:"projects,omitempty"`
	RunSTR   string             `json:"runstr,omitempty"`
	Owner    string             `json:"owner,omitempty"`
	DBClient *postgres.PGClient `json:",omitempty"`
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
	err = client.Run(s.Host, "pkill "+s.Name)
	if err != nil {
	}
	return err
}

//CreateBin - create bin from bins dir and insert data from tBins
func (s *Service) CreateBin(f io.Reader, own string) error {
	newFile := os.Getenv("BINSDIR") + "/" + s.Type + "/" + s.Name
	file, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		return err
	}
	projectsID, err := s.DBClient.GetProjectsIDtoString(s.Projects)
	if err != nil {
		return err
	}
	return s.DBClient.NewServiceBin(s.Name, s.Type, s.RunSTR, own, projectsID)
}

//UpdateBin - update bins and data from tBins
func (s *Service) UpdateBin() error {
	projectsID, err := s.DBClient.GetProjectsIDtoString(s.Projects)
	if err != nil {
		return err
	}
	return s.DBClient.UpdateServiceBin(s.ID, s.RunSTR, projectsID)
}

//DeleteBin - delete bins and data from tBins
func (s *Service) DeleteBin() error {
	err := os.Remove(os.Getenv("BINSDIR") + "/" + s.Type + "/" + s.Name)
	if err != nil {
		return err
	}
	return s.DBClient.DeleteServiceBin(s.ID)
}

//Create - install new service from remote host
func (s *Service) Create(user string, owner string) (err error) {
	sshclient, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	ext := filepath.Ext(s.Name)
	var fileName, archType string
	switch ext {
	case ".gz":
		fileName = strings.TrimSuffix(s.Name, ".tar.gz")
		archType = ".tar.gz"

	case ".zip":
		fileName = strings.TrimSuffix(s.Name, ".zip")
		archType = ".zip"
	}
	err = sshclient.InstallServiceToRemoteHost(s.Type, s.Name, s.Host, archType)
	if err != nil {
		return err
	}
	serviceIDs, err := s.DBClient.GetProjectsIDtoString(s.Projects)
	if err != nil {
		return err
	}
	return s.DBClient.NewService(fileName, s.BinsID, s.Host, s.Port, s.Type, s.RunSTR, serviceIDs, owner)
}

//Update - update service info from database
func (s *Service) Update() error {
	return s.DBClient.UpdateService(s.ID, s.Port, s.RunSTR)
}

//Delete - install new service from remote host
func (s *Service) Delete(user string) (err error) {
	sshclient, err := client.SSHClient{}.New(user)
	if err != nil {
		return err
	}
	err = sshclient.DeleteServiceFromRemoteHost(s.Type, s.Name, s.Host)
	if err != nil {
		return err
	}
	err = s.DBClient.DeleteService(s.ID)
	if err != nil {
		return err
	}
	return nil
}
