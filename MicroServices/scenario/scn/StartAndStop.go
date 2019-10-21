package scn

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/matscus/Hamster/Package/Clients/client"
)

//StartScenario - func for start test scenario and update state after the finish test
func StartScenario(runid int64, host string, pathScript string, fileName string, str string) (err error) {
	user, _ := HostsAndUsers.Load(host)
	cl, err := client.SSHClient{}.New(user.(string))
	RunsGenerators.Store(host, host)
	err = cl.Run(host, "mkdir /home/"+user.(string)+"/scripts")
	if err != nil {
		return err
	}
	copyPath := pathScript + fileName
	cmd := exec.Command("scp", copyPath, user.(string)+"@"+host+":/home/"+user.(string)+"/scripts/")
	err = cmd.Run()
	if err != nil {
		return err
	}
	unzipSTR := "unzip /home/" + user.(string) + "/scripts/" + fileName + " -d " + "/home/" + user.(string) + "/scripts/"
	err = cl.Run(host, unzipSTR)
	if err != nil {
		return err
	}
	filesInfo, err := ioutil.ReadDir("/home/" + user.(string) + "/scripts/")
	if err != nil {
		return err
	}
	for i := 0; i < len(filesInfo); i++ {
		name := filesInfo[i].Name()
		if strings.Contains(name, ".jmx") {
			text := strings.Replace(str, "$$$", name+" ", 1)
			log.Println("text  ", text)
			err = cl.Run(host, text)
			if err != nil {
				return err
			}
			break
		}
	}
	err = cl.Run(host, "rm -rf /home/"+user.(string)+"/scripts/")
	if err != nil {
		return err
	}
	err = client.PGClient{}.New().SetStopTest(strconv.FormatInt(runid, 10))
	if err != nil {
		return err
	}
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
