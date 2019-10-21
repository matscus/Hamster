package scn

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/matscus/Hamster/Package/Scenario/scenario"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Generators/generators"
)

//StartRequest - struct request for start scenario
type StartRequest struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Gun        string   `json:"gun"`
	Projects   []string `json:"projects"`
	Generators []generators.Generator
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
			pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects[0] + "/" + s.Gun + "/"
			err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
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
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects[0] + "/" + s.Gun + "/"
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
			str = str + "&> /dev/null"
			pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects[0] + "/" + s.Gun + "/"
			err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
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
				str = str + "&> /dev/null"
				pathScript := os.Getenv("DIRPROJECTS") + "/" + s.Projects[0] + "/" + s.Gun + "/"
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, pathScript, s.Name+".zip", str)
			}
		}
	}
	var duration int64
	for _, v := range s.Params {
		for _, v1 := range v.ThreadGroupParams {
			if v1.Type == "Hold" || v1.Type == "Duration" {
				d, _ := strconv.Atoi(v1.Value)
				duration = int64(d)
			}
		}
	}
	SetState(true, runid, s.Name, s.Type, duration, s.Gun, s.Generators)
	return err
}
