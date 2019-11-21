package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"github.com/kballard/go-shellquote"
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
			log.Panicf("[ERROR] %s", "param str is nil")
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
	//session.Start("nohup " + os.Getenv("PROMETHEUSDIR") + "./prometheus --web.listen-address=" + os.Getenv("PROMETHEUSPORT") + " --config.file=" + os.Getenv("PROMETHEUSDIR") + "prometheus.yml &> /dev/null")
	filePath := os.Getenv("HOME") + "/scripts_jmeter/" + scriptname + ".jmx"
	destinationPath := os.Getenv("HOME") + "/scripts_jmeter/" + scriptname + ".jmx"
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	err = copyPath(filePath, destinationPath, session)
	if err != nil {
		log.Printf("[ERROR] copy %s", err)
	}
	str := "nohup " + os.Getenv("HOME") + "/apache-jmeter-4.0/bin/jmeter -n -t " + os.Getenv("HOME") + "/scripts_jmeter/" + scriptname + ".jmx &> /dev/null"
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
	err = session.Start(os.Getenv("HOME") + "/apache-jmeter-4.0/bin/./shutdown.sh")
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	wg.Done()
}

func copyPath(filePath, destinationPath string, session *ssh.Session) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := f.Stat()
	if err != nil {
		return err
	}
	return copy(s.Size(), s.Mode().Perm(), path.Base(filePath), f, destinationPath, session)
}

func copy(size int64, mode os.FileMode, fileName string, contents io.Reader, destination string, session *ssh.Session) error {
	defer session.Close()
	w, err := session.StdinPipe()
	if err != nil {
		return err
	}
	cmd := shellquote.Join("scp", "-t", destination)
	if err := session.Start(cmd); err != nil {
		w.Close()
		return err
	}
	errors := make(chan error)
	go func() {
		errors <- session.Wait()
	}()
	fmt.Fprintf(w, "C%#o %d %s\n", mode, size, fileName)
	io.Copy(w, contents)
	fmt.Fprint(w, "\x00")
	w.Close()
	return <-errors
}
