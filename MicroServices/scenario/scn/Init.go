package scn

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Clients/client/postgres"
	"github.com/matscus/Hamster/Package/Hosts/hosts"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
)

var (
	//GetResponseAllData - struct for return information if scenarios and hosts.
	GetResponseAllData = GetResponse{}
	//GetState - struct for return state of scenario
	GetState = []scenario.State{}
	//LastRunsParams - sync map for last runs param from scenario
	LastRunsParams sync.Map
	//RunsGenerators - sync map for used generator
	RunsGenerators sync.Map
	//HostsAndUsers - sync map users from remote host
	HostsAndUsers sync.Map
	//PgClient - postgres client from service
	PgClient *postgres.PGClient
)

//GetResponse -  struct for response
type GetResponse struct {
	Generators []hosts.Host
	Scenarios  []scenario.Scenario
}

//GeneratorState -  struct for generator state
type GeneratorState struct {
	Host  string `json:"name"`
	State string `json:"state"`
}

//InitData - func to init GetRes
func InitData() (err error) {
	scenarios, err := PgClient.GetAllScenarios()
	if err != nil {
		return err
	}
	l := len(*scenarios)
	GetResponseAllData.Scenarios = make([]scenario.Scenario, 0, l)
	for i := 0; i < l; i++ {
		var tgp []scenario.ThreadGroup
		err := json.Unmarshal([]byte((*scenarios)[i].TreadGroups), &tgp)
		if err != nil {
			return (err)
		}
		GetResponseAllData.Scenarios = append(GetResponseAllData.Scenarios, scenario.Scenario{
			ID:           (*scenarios)[i].ID,
			Name:         (*scenarios)[i].Name,
			Type:         (*scenarios)[i].Type,
			LastModified: (*scenarios)[i].LastModified,
			Gun:          (*scenarios)[i].Gun,
			Projects:     (*scenarios)[i].Projects,
			ThreadGroups: tgp,
		})
	}
	gen, err := PgClient.GetAllGenerators()
	if err != nil {
		return err
	}
	l = len(gen)
	GetResponseAllData.Generators = make([]hosts.Host, 0, l)
	for i := 0; i < l; i++ {
		_, ok := RunsGenerators.Load(gen[i].Host)
		if ok {
			gen[i].State = "IsBusy"
		} else {
			gen[i].State = "Free"
		}
		GetResponseAllData.Generators = append(GetResponseAllData.Generators, hosts.Host{
			ID:       gen[i].ID,
			Host:     gen[i].Host,
			Type:     gen[i].Type,
			Projects: gen[i].Projects,
			State:    gen[i].State,
		})
	}
	hostsAndUsers, err := PgClient.GetUsersAndHosts()
	for k, v := range hostsAndUsers {
		HostsAndUsers.Store(k, v)
	}
	return nil
}

//SetState -  init state struct for ws
func SetState(s bool, id int64, name string, scenarioType string, d int64, gun string, generators []hosts.Host) {
	if s {
		GetState = append(GetState, scenario.State{
			RunID:      id,
			Name:       name,
			Type:       scenarioType,
			StartTime:  (time.Now().Unix() - time.Unix(10800, 0).Unix()),
			EndTime:    ((time.Now().Unix() - time.Unix(10800, 0).Unix()) + time.Unix(d, 0).Unix()),
			Gun:        gun,
			Generators: generators,
		})
		return
	}
	for i := 0; i < len(GetState); i++ {
		if id == GetState[i].RunID {
			removeState(i)
		}
	}
}

func removeState(s int) {
	GetState = append(GetState[:s], GetState[s+1:]...)
}

//FloatToString - convert type float  to type string
func FloatToString(i float64) string {
	return strconv.FormatFloat(i, 'f', 2, 64)
}

//СheckRun - fucn for check state scenario
func СheckRun() (res bool) {
	if len(GetState) > 0 {
		res = true
	}
	return res
}

//CheckGen - func fo check state generators
func CheckGen(g []hosts.Host) (res []GeneratorState, err error) {

	l := len(g)
	for i := 0; i < l; i++ {
		user, _ := HostsAndUsers.Load(g[i].Host)
		client, err := client.SSHClient{}.New(user.(string))
		if err != nil {
			return res, err
		}
		_, err = client.Ping(g[i].Host)
		if err != nil {
			res = append(res, GeneratorState{
				Host:  g[i].Host,
				State: "NotAvailable",
			})
			return res, err
		}
	}
	for i := 0; i < l; i++ {
		host, ok := RunsGenerators.Load(g[i].Host)
		if ok {
			res = append(res, GeneratorState{
				Host:  host.(string),
				State: "IsBusy",
			})
		}
	}
	return res, err
}

//CheckRunsGen - func for chack state generators and change state
func CheckRunsGen() {
	l := len(GetResponseAllData.Generators)
	for i := 0; i < l; i++ {
		_, ok := RunsGenerators.Load(GetResponseAllData.Generators[i].Host)
		if ok {
			GetResponseAllData.Generators[i].State = "IsBusy"
		} else {
			GetResponseAllData.Generators[i].State = "Free"
		}
	}
}
