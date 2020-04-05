package logger

import (
	"os"
	"strings"
	"time"
)

func SaveToFile(data []byte, filename string) error {
	fileName := "./log/" + strings.ReplaceAll(strings.Split(filename, ":")[0], ".", "_") + " " + time.Now().Format("15_04_05") + ".log"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(data))
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
