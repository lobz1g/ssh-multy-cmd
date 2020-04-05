package connection

import (
	"encoding/json"
	"log"
	"os"
)

type (
	config struct {
		User     string
		Password string
		Host     string
		Cmd      string
	}

	cfg struct {
		Path string
	}
)

func newConfig(p ...string) *cfg {
	c := &cfg{}
	if len(p) != 0 {
		c.Path = p[0]
	} else {
		c.Path = "./config/config.json"
	}
	return c
}

func (c *cfg) getConfig() ([]config, error) {
	file, err := os.Open(c.Path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	cfg := []config{}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cfg, err
}
