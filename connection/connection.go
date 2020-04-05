package connection

import (
	"bytes"
	"log"
	"net"

	"golang.org/x/crypto/ssh"

	"ssh/logger"
)

type Connection struct {
	client *ssh.Client
	cfg    config
	Output struct {
		Host string
	}
}

func newConnection(cfg config) (*Connection, error) {
	sshConfig := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.Password),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	conn, err := ssh.Dial("tcp", cfg.Host, sshConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Connection{client: conn, cfg: cfg}, nil
}

func OpenConnections() ([]*Connection, error) {
	cfg, err := newConfig().getConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ch := make(chan *Connection)
	go func(ch chan *Connection, cfg []config) {
		for _, v := range cfg {
			c, err := newConnection(v)
			if err != nil {
				log.Println(err)
			}
			ch <- c
		}
	}(ch, cfg)

	allConnection := []*Connection{}
	for i := 0; i < len(cfg); i++ {
		allConnection = append(allConnection, <-ch)
	}

	return allConnection, nil
}

func (conn *Connection) SendCommands() {
	session, err := conn.client.NewSession()
	if err != nil {
		logger.SaveToFile([]byte(err.Error()), conn.cfg.Host)
		return
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(conn.cfg.Cmd)
	if err != nil {
		logger.SaveToFile([]byte(err.Error()), conn.cfg.Host)
		return
	}

	err = logger.SaveToFile(stdoutBuf.Bytes(), conn.cfg.Host)
	if err != nil {
		log.Println(err)
	}
}
