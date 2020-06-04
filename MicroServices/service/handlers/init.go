package handlers

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
	"github.com/matscus/Hamster/Package/Services/service"
)

var (
	pgClient      *postgres.PGClient
	HostsAndUsers sync.Map
	AllService    = make(map[int64]service.Service)
)

func init() {
	pgConf := postgres.Config{Driver: "postgres", User: os.Getenv("POSTGRESUSER"), Password: os.Getenv("POSTGRESPASSWORD"), DataBase: os.Getenv("POSTGRESDB"), SSLMode: "disable"}
	client := client.New("postgres", pgConf).(postgres.PGClient)
	pgClient = &client
	hostsAndUsers, err := pgClient.GetUsersAndHosts()

	if err != nil {
		log.Println("[ERR] Error get users and hosts: ", err)
	}
	for k, v := range hostsAndUsers {
		HostsAndUsers.Store(k, v)
	}
}

//InitGetResponseAllData - init gorutune for get actual status once in the minute
func InitGetResponseAllData() error {
	for {
		s, err := pgClient.GetAllServices()
		if err != nil {
			return err
		}
		l := len(*s)
		for i := 0; i < l; i++ {
			AllService[(*s)[i].ID] = service.Service{
				ID:       (*s)[i].ID,
				Host:     (*s)[i].Host,
				Name:     (*s)[i].Name,
				Port:     (*s)[i].Port,
				Owner:    (*s)[i].Owner,
				Status:   (*s)[i].Status,
				RunSTR:   (*s)[i].RunSTR,
				BinsID:   (*s)[i].BinsID,
				Projects: (*s)[i].Projects,
				Type:     (*s)[i].Type}
		}
		time.Sleep(1 * time.Minute)
	}
}

//OnceInitAllData func for update init data
func OnceInitAllData() error {
	s, err := pgClient.GetAllServices()
	if err != nil {
		return err
	}
	l := len(*s)
	for i := 0; i < l; i++ {
		AllService[(*s)[i].ID] = service.Service{
			ID:       (*s)[i].ID,
			Host:     (*s)[i].Host,
			Name:     (*s)[i].Name,
			Port:     (*s)[i].Port,
			Owner:    (*s)[i].Owner,
			Status:   (*s)[i].Status,
			RunSTR:   (*s)[i].RunSTR,
			BinsID:   (*s)[i].BinsID,
			Projects: (*s)[i].Projects,
			Type:     (*s)[i].Type}
	}
	return err
}

//CheckService - The function to check the status of web services.
//Performs a GET request at the of each service. if available, sets the service in the status Run.
func CheckService() {
	var err error
	var tcpAddr net.TCPAddr
	for k, v := range AllService {
		if strconv.Itoa(v.Port) == "" {
			v.Status = "service does not have listen port"
		}
		tcpAddr = net.TCPAddr{IP: []byte(v.Host), Port: v.Port}
		_, err = net.DialTCP("tcp", nil, &tcpAddr)
		if err != nil {
			v.Status = "stop"
		} else {
			v.Status = "run"
		}
		AllService[k] = v
	}
}

//UpdateState - update service info
func UpdateState(s *service.Service) {
	if val, ok := AllService[s.ID]; ok {
		val.Status = s.Status
		AllService[s.ID] = val
	}
}
