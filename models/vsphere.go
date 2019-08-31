package models

import "time"

type VSphereDatacenter struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Collected time.Time `json:"collected"`
}

type VSphereHost struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Cluster     string     `json:"cluster"`
	PowerState  string     `json:"power_state"`
	BootTime    *time.Time `json:"boot_time"`
	NCPU        int        `json:"ncpu"`
	Memory      int64      `json:"memory"`
	Collected   time.Time  `json:"collected"`
	Environment string     `json:"environment"`
	VMs         int        `json:"vms"`
	DataCenter  string     `json:"datacenter"`
	Networks    []string   `json:"networks`
}

type VSphereVM struct {
	Id            string     `son:"id"`
	HostId        string     `json:"host_id"`
	HostName      string     `json:"host_name"`
	Cluster       string     `json:"cluster"`
	DatastoreId   string     `json:"datastore_id"`
	DatastoreName string     `json:"datastore_name"`
	Name          string     `json:"name"`
	PowerState    string     `json:"power_state"`
	BootTime      *time.Time `json:"boot_time"`
	NCPU          int        `json:"ncpu"`
	Memory        int64      `json:"memory"`
	Storage       int64      `json:"storage"`
	IP            string     `json:"ip"`
	Collected     time.Time  `json:"collected"`
	Environment   string     `json:"environment"`
	DataCenter    string     `json:"datacenter"`
}

type VSphereDatastore struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Capacity    int64     `json:"capacity"`
	Free        int64     `json:"free"`
	Collected   time.Time `json:"collected"`
	Environment string    `json:"environment"`
	VMs         int       `json:"vms"`
	DataCenter  string    `json:"datacenter"`
}

type VSpherePayload struct {
	Hosts      []VSphereHost      `json:"hosts"`
	DataStores []VSphereDatastore `json:"data_stores"`
	VMs        []VSphereVM        `json:"vms"`
}

type VSphereNetwork struct {
	Name    string `json:"name"`
	Cluster string `json:"cluster"`
}
