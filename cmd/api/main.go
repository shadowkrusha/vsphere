package main

import (
	// "context"
	"fmt"
	"github.com/shadowkrusha/vsphere/api"
	// "time"
)

func main() {
	c, err := api.NewVSphereCollector("https://user:pass@127.0.0.1:8989/sdk")
	if err != nil {
		fmt.Println("Error", err)
	}

	data, err := c.Collect()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("%+v", data)
}
