package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/jenningsloy318/netapp_exporter/collector"
)

var (
	configFile = kingpin.Flag(
		"config.file",
		"Path to configuration file.",
	).String()
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9109").String()

	reloadCh chan chan error
)

// define new http handleer
func metricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var scrapers =  []collector.Scraper {
			//collector.system,
		}
		registry := prometheus.NewRegistry()

		filers := loadFilerFromFile(*configFile)
		for  _, f := range filers {
			netappClient := newNetappClient(f)
			collector := collector.New(netappClient, scrapers)
			registry.MustRegister(collector)
		}
		gatherers := prometheus.Gatherers{
			prometheus.DefaultGatherer,
			registry,
		}
		// Delegate http serving to Prometheus client library, which will call collector.Collect.
		h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	}
}





func main() {
	log.AddFlags(kingpin.CommandLine)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log.Infoln("Starting netapp_exporter")

	http.Handle("/metrics", metricsHandler())       // Regular metrics endpoint for local IPMI metrics.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head>
            <title>NetApp Exporter</title>
            <style>
            label{
            display:inline-block;
            width:75px;
            }
            form label {
            margin: 10px;
            }
            form input {
            margin: 10px;
            }
            </style>
            </head>
            <body>
            <h1>IPMI Exporter</h1>
            <form action="/ipmi">
            <label>Target:</label> <input type="text" name="target" placeholder="X.X.X.X" value="1.2.3.4"><br>
            <input type="submit" value="Submit">
			</form>
			<p><a href="/metrics">Local metrics</a></p>
			<p><a href="/config">Config</a></p>
            </body>
            </html>`))
	})

	log.Infof("Listening on %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
