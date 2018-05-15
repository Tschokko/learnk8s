package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/tschokko/learnk8s/pkg/controller"
	yaml "gopkg.in/yaml.v2"
)

// Config contains the setting of the OAM service daemon
type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
}

func load(file string) (cfg Config, err error) {
	cfg.Server.Addr = ":7777"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return
	}

	return
}

func main() {
	filename := flag.String("config", "config.yml", "Configuration file")
	flag.Parse()

	config, err := load(*filename)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	lis, err := net.Listen("tcp", config.Server.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s, _ := controller.NewServer()
	log.Println("start listening on port 7777")
	if err := s.Run(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
