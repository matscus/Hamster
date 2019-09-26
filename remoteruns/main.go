package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

var (
	scriptName string
	action     string
	shhcnfg    *ssh.ClientConfig
)

func init() {
	var err error
	key, err := ioutil.ReadFile(os.Getenv("RSAPATH"))
	if err != nil {
		log.Printf("Unable to read private key:: %s", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("Unable to parse private key: : %s", err)
	}
	shhcnfg = &ssh.ClientConfig{
		User: os.Getenv("SSHUSER"),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type RemoteAddr []struct {
	Host string `json:"Host"`
}

func main() {
	flag.StringVar(&scriptName, "script", "", "script name without .jmx")
	flag.StringVar(&action, "action", "start", "value for action, default start")
	flag.Parse()
	switch action {
	case "start":
		if scriptName == "" {
			log.Panic("[ERROR] %s", "param str is nil")
		}
		file, err := os.Open("config.json")
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		cnfg := RemoteAddr{}
		err = json.NewDecoder(file).Decode(&cnfg)
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		l := len(cnfg)
		var wg sync.WaitGroup
		for i := 0; i < l; i++ {
			wg.Add(1)
			go start(cnfg[i].Host, scriptName, &wg)
		}
		wg.Wait()
		log.Printf("[INFO] %s", "start is complited")
	case "stop":
		cnfg := RemoteAddr{}
		file, err := os.Open("config.json")
		err = json.NewDecoder(file).Decode(&cnfg)
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		l := len(cnfg)
		var wg sync.WaitGroup
		for i := 0; i < l; i++ {
			wg.Add(1)
			go stop(cnfg[i].Host, &wg)
		}
		wg.Wait()
		log.Printf("[INFO] %s", "stop is complited")

	default:
		log.Panic("[ERROR] %s", "params action is not a validate")
	}

}

func start(host string, scriptname string, wg *sync.WaitGroup) {
	client, err := ssh.Dial("tcp", host+":22", shhcnfg)
	session, err := client.NewSession()
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	defer session.Close()
	filePath := "~/scripts_jmeter/" + scriptname + ".jmx"
	destinationPath := "~/scripts_jmeter/" + scriptname + ".jmx"
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	err = scp.CopyPath(filePath, destinationPath, session)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	str := "nohup ~/apache-jmeter-4.0/bin/jmeter -n -t ~/scripts_jmeter/" + scriptname + ".jmx &> /dev/null"
	log.Printf("[INFO] Str for run %s", str)
	err = session.Start(str)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	wg.Done()
}
func stop(host string, wg *sync.WaitGroup) {
	client, err := ssh.Dial("tcp", host+":22", shhcnfg)
	session, err := client.NewSession()
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	defer session.Close()
	err = session.Start("~/apache-jmeter-4.0/bin/./shutdown.sh")
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	wg.Done()
}
