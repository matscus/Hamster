package sshimpl

import "golang.org/x/crypto/ssh"

//SSHClient - ssh client implementation struct, whit ssh config
type SSHClient struct {
	SHHConfig *ssh.ClientConfig
}
