package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	controller "github.com/tschokko/learnk8s/pkg/controller/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	yaml "gopkg.in/yaml.v2"
)

const banner = `OAM Service
-----------

Service-ID:   %s
Registered:   %t
Host:         %s
Request path: %s`

// Config contains the setting of the OAM service daemon
type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	OAM struct {
		ServiceID                    string `yaml:"serviceID"`
		ControllerAddr               string `yaml:"controllerAddr"`
		ControllerSSLCaFile          string `yaml:"controllerSslCaFile"`
		ControllerServerHostOverride string `yaml:"controllerServerHostOverride"`
	} `yaml:"oam"`
}

func load(file string) (cfg Config, err error) {
	cfg.Server.Addr = ":8080"
	cfg.OAM.ServiceID = "unassigned"

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

	// Connect to oam-controller and register the service
	var registered = false

	// SSL
	var opts []grpc.DialOption
	// caFile := "icomcloud-ca.crt"
	// serverHostOverride := "oam-controller-nsys.eu-west-1.icomcloud.net"
	creds, err := credentials.NewClientTLSFromFile(config.OAM.ControllerSSLCaFile, config.OAM.ControllerServerHostOverride)
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(config.OAM.ControllerAddr, grpc.WithInsecure()) // opts...

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := controller.NewServiceControllerClient(conn)
	res, err := c.RegisterService(context.Background(), &controller.RegisterServiceRequest{ServiceID: config.OAM.ServiceID})
	if err != nil {
		log.Fatalf("error when calling RegisterService: %s", err)
	}
	registered = res.Success
	// End of service registration

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serviceID := config.OAM.ServiceID
		name, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, banner, serviceID, registered, name, r.URL.Path)
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
