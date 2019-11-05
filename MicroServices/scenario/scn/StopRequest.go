package scn

import (
	"os"
	"strconv"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Hosts/hosts"
)

//StopRequest - struct request for stop scenario
type StopRequest struct {
	RunID      int64  `json:"runid"`
	Gun        string `json:"gun"`
	Generators []hosts.Host
}

//Stop - func for the stop scenario
func (s *StopRequest) Stop() error {
	var err error
	gencount := len(s.Generators)
	switch s.Gun {
	case "gatling":
		str := "kill `jps | grep \"Launcher\" | cut -d \" \" -f 1`"
		for i := 0; i < gencount; i++ {
			go StopScenario(s.RunID, s.Generators[i].Host, str)
		}
	case "jmeter":
		str := os.Getenv("JMETERPATH") + "./shutdown.sh"
		for i := 0; i < gencount; i++ {
			go StopScenario(s.RunID, s.Generators[i].Host, str)
		}
	}
	err = client.PGClient{}.New().SetStopTest(strconv.FormatInt(s.RunID, 10))
	return err
}
