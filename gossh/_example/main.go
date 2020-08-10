package main

import (
	"golang.org/x/crypto/ssh"
	"net"
	"study_go/gossh"
)

func main() {

	// local loop login
	var config = ssh.ClientConfig{
		User: "mm2",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	var client = gossh.NewSSHClient("127.0.0.1", 22, &config)
	_ = client.TerminalConnect()
}
