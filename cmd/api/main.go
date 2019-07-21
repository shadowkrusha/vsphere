package main

import (
	// "context"
	// "fmt"
	"github.com/shadowkrusha/vsphere/api"
	// "time"
)

func main() {
	c, _ := api.NewVSphereCollector("https://user:pass@127.0.0.1:8989/sdk")

	c.Collect()
}
