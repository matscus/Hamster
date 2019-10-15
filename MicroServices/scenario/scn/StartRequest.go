package scn

import (
	"math"
	"os"
	"strconv"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Generators/generators"
)

//StartRequest - struct request for start scenario
type StartRequest struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Users      int    `json:"users"`
	Rampup     int    `json:"rampup"`
	Duration   int64  `json:"duration"`
	Throughput int    `json:"throughput"`
	Gun        string `json:"gun"`
	Generators []generators.Generator
}

//Start - func for the run scenario
func (s *StartRequest) Start() error {
	var err error
	var u, g, mod, userForGen float64
	var str string
	gencount := len(s.Generators)
	u = float64(s.Users)
	g = float64(gencount)
	mod = math.Mod(u, g)
	userForGen = math.RoundToEven(u)
	pgclient := client.PGClient{}.New()
	runid, err := pgclient.GetNewRunID()
	if err != nil {
		return err
	}
	switch s.Gun {
	case "gatling":
		for i := 0; i < gencount; i++ {
			if i == 0 {
				users := FloatToString(userForGen + mod)
				str = `cd ` + os.Getenv("MAVENPATH") + ` && mvn clean gatling:execute -Dgatling.simulationClass=com.testingexcellence.simulations.` + s.Name + `
				 -Dusers=` + users + ` -Drampup=` + strconv.Itoa(s.Rampup) + ` -Dduration=` + strconv.FormatInt(s.Duration, 10) + ` -Dthroughput=` + strconv.Itoa(s.Throughput) + ``
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, str)
			} else {
				str = `cd ` + os.Getenv("MAVENPATH") + ` && mvn clean gatling:execute -Dgatling.simulationClass=com.testingexcellence.simulations.` + s.Name + `
				 -Dusers=` + FloatToString(userForGen) + ` -Drampup=` + strconv.Itoa(s.Rampup) + ` -Dduration=` + strconv.FormatInt(s.Duration, 10) + `
				  -Dthroughput=` + strconv.Itoa(s.Throughput) + ``
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, str)
			}
		}
	case "jmeter":
		for i := 0; i < gencount; i++ {
			if i == 0 {
				users := FloatToString(userForGen + mod)
				str = "nohup " + os.Getenv("JMETERPATH") + "./jmeter.sh -Jtps=" + users + " -Jrtime=" + "5" + " -Jrstep=" + strconv.Itoa(s.Rampup) + " -Jduration=" + strconv.FormatInt(s.Duration, 10) + " -n -t " + os.Getenv("JMETERSCRIPTPATH") + s.Name + ".jmx" + " &> /dev/null"
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, str)
			} else {
				str = "nohup " + os.Getenv("JMETERPATH") + "./jmeter.sh -Jtps=" + FloatToString(userForGen) + " -Jrtime=" + "5" + " -Jrstep=" + strconv.Itoa(s.Rampup) + " -Jduration=" + strconv.FormatInt(s.Duration, 10) + " -n -t " + os.Getenv("JMETERSCRIPTPATH") + s.Name + ".jmx" + " &> /dev/null"
				err = pgclient.SetStartTest(strconv.FormatInt(runid, 10), s.Name, s.Type)
				if err != nil {
					return err
				}
				go StartScenario(runid, s.Generators[i].Host, str)
			}
		}
	}
	SetState(true, runid, s.Name, s.Type, s.Duration, s.Gun, s.Generators)
	return err
}

//CheckParams - func check request params? if once param in nil, return false
func (s *StartRequest) CheckParams() bool {
	if s.Name == "" || s.Users == 0 || s.Rampup == 0 || s.Duration == 0 || s.Throughput == 0 || len(s.Generators) == 0 || s.Gun == "" {
		return false
	}
	return true
}
