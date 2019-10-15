package scn

import (
	"strconv"

	"github.com/matscus/Hamster/Package/Clients/client"
)

//StartScenario - func for start test scenario and update state after the finish test
func StartScenario(runid int64, host string, str string) (err error) {
	user, _ := HostsAndUsers.Load(host)
	cl, err := client.SSHClient{}.New(user.(string))
	RunsGenerators.Store(host, host)
	err = cl.Run(host, str)
	err = client.PGClient{}.New().SetStopTest(strconv.FormatInt(runid, 10))
	RunsGenerators.Delete(host)
	SetState(false, runid, "", "", 0, "", nil)
	return err
}

//StopScenario - func for stop test scenario and update state
func StopScenario(runid int64, host string, str string) (err error) {
	user, _ := HostsAndUsers.Load(host)
	client, err := client.SSHClient{}.New(user.(string))
	err = client.Run(host, str)
	RunsGenerators.Delete(host)
	SetState(false, runid, "", "", 0, "", nil)
	return err
}
