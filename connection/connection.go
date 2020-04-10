package connection

import (
	"bytes"
	"errors"
	"net"
	"sync"

	"golang.org/x/crypto/ssh"

	"ssh/message"
)

type Connection struct {
	client *ssh.Client
	cfg    config
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
		return &Connection{cfg: cfg}, err
	}

	return &Connection{client: conn, cfg: cfg}, nil
}

func OpenConnections() ([]*Connection, error) {
	cfg, err := newConfig().getConfig()
	if err != nil {
		return nil, err
	}

	allConnections := []*Connection{}
	w, mu := new(sync.WaitGroup), new(sync.Mutex)
	for _, v := range cfg {
		w.Add(1)
		go func(cfg config, w *sync.WaitGroup, mu *sync.Mutex) {
			defer w.Done()
			c, err := newConnection(cfg)
			if err != nil {
				message.NewMessage("localhost", err.Error(), true).Save()
			}

			mu.Lock()
			allConnections = append(allConnections, c)
			mu.Unlock()

			return
		}(v, w, mu)
	}
	w.Wait()
	return allConnections, nil
}

func (c *Connection) SendCommands() *message.Message {
	if c.client == nil {
		return message.NewMessage(c.cfg.Host, errors.New("see localhost.log file").Error(), true)
	}

	session, err := c.client.NewSession()
	if err != nil {
		return message.NewMessage(c.cfg.Host, err.Error(), true)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(c.cfg.Cmd)
	if err != nil {
		return message.NewMessage(c.cfg.Host, err.Error(), true)
	}
	return message.NewMessage(c.cfg.Host, string(stdoutBuf.Bytes()), false)
}
