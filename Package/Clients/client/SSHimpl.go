package client

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"

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
func (c SSHClient) InstallServiceToRemoteHost(serviceType string, name string, target string) (err error) {
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
	path := "~/Hamster/bins/" + serviceType + "/" + name + "/"
	test := regexp.MustCompile("([A-Za-z0-9]+)")
	dirSlice := test.FindAllStringSubmatch(path, -1)
	l := len(dirSlice)
	//values IsNotExist in the byte slice
	isNotExist := []byte{73, 115, 78, 111, 116, 69, 120, 105, 115, 116, 10}
	tempPath := "~/"
	for i := 0; i < l; i++ {
		res, err := session.CombinedOutput("[ -d ~/" + dirSlice[i][0] + "/ ] && echo 'ok' || echo 'IsNotExist'")
		if err != nil {
			return err
		}
		ok := bytes.Equal(res, isNotExist)
		if ok {
			continue
		} else {
			err = session.Run("mkdir " + tempPath + dirSlice[i][0])
			tempPath = tempPath + dirSlice[i][0] + "/"
		}
	}
	err = scp.CopyPath(path+"compress.tar.gzip", path+"compress.tar.gzip", session)
	if err != nil {
		return err
	}
	err = session.Run("tar -xf " + path + "compress.tar.gzip")
	if err != nil {
		return err
	}
	return nil
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
	err = session.Run("rm -rf ~/Hamster/bins/" + serviceType + "/" + name + "/")
	if err != nil {
		return err
	}
	return nil
}
