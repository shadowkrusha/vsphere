package api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	unrender "github.com/unrolled/render"
	"net/http"
	"runtime"
	"strconv"
)

type HttpServer struct {
	Config *Config
}

type Results struct {
	Data interface{} `json:"data"`
}

func (s *HttpServer) Start() {

	render := unrender.New(unrender.Options{
		IndentJSON: true,
		Layout:     "layout",
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		render.Text(w, http.StatusOK, "pong")
	})
	http.HandleFunc("/config", func(w http.ResponseWriter, req *http.Request) {
		render.JSON(w, http.StatusOK, s.Config)
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		render.Text(w, http.StatusOK, "OK")
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		info := map[string]string{
			// "syros_version": version,
			"os":         runtime.GOOS,
			"arch":       runtime.GOARCH,
			"golang":     runtime.Version(),
			"max_procs":  strconv.FormatInt(int64(runtime.GOMAXPROCS(0)), 10),
			"goroutines": strconv.FormatInt(int64(runtime.NumGoroutine()), 10),
			"cpu_count":  strconv.FormatInt(int64(runtime.NumCPU()), 10),
		}
		render.JSON(w, http.StatusOK, info)
	})
	http.HandleFunc("/data/pvc", func(w http.ResponseWriter, req *http.Request) {
		c, err := NewVSphereCollector("https://svc_datacollector:OinlMqkyfSxfuC67HGa0@pvc.ads.westernpower.com.au/sdk")
		if err != nil {
			fmt.Println("Error", err)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		data, err := c.Collect()
		if err != nil {
			fmt.Println("Error", err)
			render.JSON(w, http.StatusBadRequest, err)
		}

		results := &Results{}
		results.Data = data.VMs

		render.JSON(w, http.StatusOK, results)
	})

	http.HandleFunc("/data/svc", func(w http.ResponseWriter, req *http.Request) {
		c, err := NewVSphereCollector("https://svc_datacollector:OinlMqkyfSxfuC67HGa0@svc.ads.westernpower.com.au/sdk")
		if err != nil {
			fmt.Println("Error", err)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		data, err := c.Collect()
		if err != nil {
			fmt.Println("Error", err)
			render.JSON(w, http.StatusBadRequest, err)
		}

		results := &Results{}
		results.Data = data.VMs

		render.JSON(w, http.StatusOK, results)
	})

	log.Error(http.ListenAndServe(fmt.Sprintf(":%v", s.Config.Port), http.DefaultServeMux))
}
