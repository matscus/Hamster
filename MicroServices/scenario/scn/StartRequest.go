package scn

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/matscus/Hamster/Package/Scenario/scenario"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Hosts/hosts"
)

//StartRequest - struct request for start scenario
type StartRequest struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Gun        string `json:"gun"`
	Projects   string `json:"projects"`
	Generators []hosts.Host
	Params     []scenario.ThreadGroup
}

//Start - func for the run scenario
func (s *StartRequest) Start() error {
	var err error
	var u, g, mod, userForGen float64
	gencount := len(s.Generators)
	pgclient := client.PGClient{}.New()
	runid, err := pgclient.GetNewRunID()
	if err != nil {
		return err
	}
	switch s.Gun {
	case "gatling":
		if gencount == 1 {
			str := `cd ` + os.Getenv("MAVENPATH") + ` && mvn clean gatling:execute -Dgatling.simulationClass=com.testingexcellence.simulations.` + s.Name
			for _, v := range s.Params {
				for _, v1 := range v.ThreadGroupParams {
					str = str + "-D" + v1.Name + "=" + v1.Value
				}
			}
			pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/"
			err = pgclient.SetStartTest(s.Name, s.Type)
			if err != nil {
				return err
			}
			go StartScenario(runid, s.Generators[0].Host, pathScript, s.Name+".zip", str)
		} else {
			for i := 0; i < gencount; i++ {
				str := `cd ` + os.Getenv("MAVENPATH") + ` && mvn clean gatling:execute -Dgatling.simulationClass=com.testingexcellence.simulations.` + s.Name
				for _, v := range s.Params {
					for _, v1 := range v.ThreadGroupParams {
						if v1.Type == "Threads" || v1.Type == "TargetLevel" {
							u, _ = strconv.ParseFloat(v1.Value, 64)
							g = float64(gencount)
							mod = math.Mod(u, g)
							userForGen = math.RoundToEven(mod)
							str = str + "-D" + v1.Name + "=" + fmt.Sprint(userForGen)
						} else {
							str = str + "-D" + v1.Name + "=" + v1.Value
						}
					}
				}
				err = pgclient.SetStartTest(s.Name, s.Type)
				pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/"
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, pathScript, s.Name+".zip", str)
			}
		}
	case "jmeter":
		if gencount == 1 {
			str := "nohup " + os.Getenv("JMETERPATH") + "./jmeter.sh -n -t scripts/$$$"
			for _, v := range s.Params {
				for _, v1 := range v.ThreadGroupParams {
					str = str + "-J" + v1.Name + "=" + v1.Value + " "
				}
			}
			str = str + " -JRunID=" + strconv.FormatInt(runid, 10) + " &> /dev/null"
			pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/"
			err = pgclient.SetStartTest(s.Name, s.Type)
			if err != nil {
				return err
			}
			go StartScenario(runid, s.Generators[0].Host, pathScript, s.Name+".zip", str)
		} else {
			for i := 0; i < gencount; i++ {
				str := "nohup " + os.Getenv("JMETERPATH") + "./jmeter.sh -n -t scripts/$$$"
				for _, v := range s.Params {
					for _, v1 := range v.ThreadGroupParams {
						if v1.Type == "Threads" || v1.Type == "TargetLevel" {
							u, _ = strconv.ParseFloat(v1.Value, 64)
							g = float64(gencount)
							mod = math.Mod(u, g)
							userForGen = math.RoundToEven(mod)
							str = str + "-J" + v1.Name + "=" + fmt.Sprint(userForGen)
						} else {
							str = str + "-J" + v1.Name + "=" + v1.Value
						}
					}
				}
				str = str + " -JRunID=" + strconv.FormatInt(runid, 10) + " &> /dev/null"
				pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/"
				go StartScenario(runid, s.Generators[i].Host, pathScript, s.Name+".zip", str)
			}
			err = pgclient.SetStartTest(s.Name, s.Type)
			if err != nil {
				return err
			}
		}
	}
	allDuration := make([]int64, 0, len(s.Params))
	for _, v := range s.Params {
		switch v.ThreadGroupType {
		case "DefaultThreadGroup":
			for _, v1 := range v.ThreadGroupParams {
				if v1.Type == "Hold" || v1.Type == "Duration" {
					d, _ := strconv.Atoi(v1.Value)
					allDuration = append(allDuration, int64(d*60))
				}
			}
		case "BlazemeterConcurrencyThreadGroup":
			var steps int64
			var rampup int64
			var duration int64
			for _, v1 := range v.ThreadGroupParams {
				switch v1.Type {
				case "RampUp":
					d, _ := strconv.Atoi(v1.Value)
					rampup = int64(d)
				case "Steps":
					d, _ := strconv.Atoi(v1.Value)
					steps = int64(d)
				case "Duration":
					d, _ := strconv.Atoi(v1.Value)
					duration = int64(d)
				}
			}
			allDuration = append(allDuration, rampup*steps+duration)
		}
	}
	sort.Slice(allDuration, func(i, j int) bool { return allDuration[i] > allDuration[j] })
	SetState(true, runid, s.Name, s.Type, allDuration[0], s.Gun, s.Generators)
	return err
}
