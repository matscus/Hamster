package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/matscus/Hamster/Package/Clients/subset"
	"github.com/tmc/scp"

	"golang.org/x/crypto/ssh"
)

var (
	pemPath string
	keyPath string
)

//SSHClient - ssh client implementation struct, whit ssh config
type SSHClient struct {
	SHHConfig *ssh.ClientConfig
}

//New - return nuw ssh client interface
func (c SSHClient) New(userName string) (subset.SSHClient, error) {
	var client subset.SSHClient
	var err error
	key, err := ioutil.ReadFile(os.Getenv("RSAPATH"))
	if err != nil {
		log.Printf("Unable to read private key:: %s", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Printf("Unable to parse private key: : %s", err)
	}
	client = SSHClient{
		SHHConfig: &ssh.ClientConfig{
			User: userName,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}
	return client, err
}

//Run - run comant and wait os code
func (c SSHClient) Run(target string, str string) error {
	var err error
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Run(str)
	return err
}

//RunNoWait - run comant no wait os code
func (c SSHClient) RunNoWait(target string, str string) error {
	var err error
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Start(str)
	return err
}

//Ping - func ping host, res false is host not avalible
func (c SSHClient) Ping(target string) (res bool, err error) {
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return false, err
	}
	defer client.Close()
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		return false, err
	}
	return true, err
}

func (c SSHClient) SCP(target, filePath, destinationPath string) error {
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	err = scp.CopyPath(filePath, destinationPath, session)
	if err != nil {
		return err
	}
	defer session.Close()
	return nil
}

//CombinedOutput - run comand and wait os code and return combined values(output + error)
func (c SSHClient) CombinedOutput(target string, str string) ([]byte, error) {
	var err error
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session.CombinedOutput(str)
}

//InstallServiceToRemoteHost -
func (c SSHClient) InstallServiceToRemoteHost(serviceType string, name string, target string, archType string) (err error) {
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	err = session.Shell()
	if err != nil {
		return err
	}
	pathBins := filepath.Join("/home", c.SHHConfig.User, "Hamster", "bins", serviceType)
	filePath := filepath.Join(os.Getenv("HOME"), "Hamster", "distr", serviceType, name)
	cmd := strings.Join([]string{"mkdir", pathBins}, " ")
	fmt.Fprintf(stdin, "%s\n", cmd)
	scp := exec.Command("scp", "-r", filePath, strings.Join([]string{c.SHHConfig.User, "@", target, ":", pathBins, name}, ""))
	err = scp.Start()
	if err != nil {
		return err
	}
	var cmdUnArch string
	switch archType {
	case ".tar.gz":
		cmdUnArch = strings.Join([]string{"tar", "-xf", name}, " ")
	case ".zip":
		cmdUnArch = strings.Join([]string{"unzip", name}, " ")
	}
	commands := []string{
		"cd " + pathBins,
		cmdUnArch,
	}
	_, err = fmt.Fprintf(stdin, "%s\n", strings.Join(commands, ";"))
	return err
}

//DeleteServiceFromRemoteHost - func connect to remote host and delete service
func (c SSHClient) DeleteServiceFromRemoteHost(serviceType string, name string, target string) (err error) {
	client, err := ssh.Dial("tcp", target+":22", c.SHHConfig)
	if err != nil {
		return err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	pathBins := filepath.Join("echo $HOME", "Hamster", "bins", serviceType, name)
	err = session.Run(strings.Join([]string{"rm", "-rf", pathBins}, " "))
	return err
}
