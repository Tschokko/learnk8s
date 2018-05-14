package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const banner = `OAM Service
-----------

Account-ID:   %s
Host:         %s
Request path: %s`

// Config contains the setting of the OAM service daemon
type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	OAM struct {
		AccountID string `yaml:"accountID"`
	} `yaml:"oam"`
}

func load(file string) (cfg Config, err error) {
	cfg.Server.Addr = ":8080"
	cfg.OAM.AccountID = "unassigned"

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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		account := config.OAM.AccountID
		name, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, banner, account, name, r.URL.Path)
	})

	srv := &http.Server{
		Addr:    config.Server.Addr,
		Handler: handler,
	}

	log.Println("start listening on 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
