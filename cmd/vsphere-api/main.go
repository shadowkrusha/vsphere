package main

import (
	// "context"
	// "fmt"
	"github.com/shadowkrusha/vsphere/api"
	log "github.com/sirupsen/logrus"
	// "time"
	// "flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var config = &api.Config{}

	// c, err := api.NewVSphereCollector("https://user:pass@127.0.0.1:8989/sdk")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }

	// data, err := c.Collect()
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	// fmt.Printf("%+v", data)

	config.Port = 8886

	server := &api.HttpServer{
		Config: config,
	}
	log.Infof("Starting HTTP server on port %v", config.Port)
	go server.Start()

	//wait for SIGINT (Ctrl+C) or SIGTERM (docker stop)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	sig := <-sigChan
	log.Infof("Shutting down %v signal received", sig)
}
