// Copyright 2020 Billy Wooten
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/billykwooten/openweather-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	version := "0.0.10"
	// Setup better logging
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(formatter)

	config := Load()

	// Create a new instance of the weatherCollector and
	// register it with the prometheus client.
	weatherCollector := collector.NewOpenweatherCollector(config.Unit, config.Language, config.ApiKey, config.City)
	prometheus.MustRegister(weatherCollector)

	// This section will start the HTTP server and expose
	// any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Beginning openweather exporter v%s to serve on port %s", version, config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
