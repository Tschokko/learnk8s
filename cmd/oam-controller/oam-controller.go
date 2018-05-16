package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/tschokko/learnk8s/pkg/controller"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"
)

// Config contains the setting of the OAM service daemon
type Config struct {
	Server struct {
		Addr        string `yaml:"addr"`
		SSLCertFile string `yaml:"sslCertFile"`
		SSLKeyFile  string `yaml:"sslKeyFile"`
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

	// SSL
	var opts []grpc.ServerOption

	// certFile := "./oam-controller-nsys.crt"
	//keyFile := "./oam-controller-nsys.key"
	/*creds, err := credentials.NewServerTLSFromFile(config.Server.SSLCertFile, config.Server.SSLKeyFile)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}*/

	s, _ := controller.NewServer()
	log.Println("start listening on port 7777")
	if err := s.Run(lis, opts...); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
