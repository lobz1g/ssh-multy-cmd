package message

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type (
	Message struct {
		Host    string
		Msg     string
		IsError bool
	}
)

func NewMessage(host string, msg string, isErr bool) *Message {
	return &Message{Host: host, Msg: msg, IsError: isErr}
}

func (m *Message) Save() error {
	fileName := fmt.Sprintf("./log/%s.log", strings.ReplaceAll(strings.Split(m.Host, ":")[0], ".", "_"))
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s: %s\n", time.Now().UTC(), m.Msg))
	if err != nil {
		return err
	}
	return nil
}
