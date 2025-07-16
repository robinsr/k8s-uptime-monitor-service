package main

import (
	"log"

	"uptime-monitor/api"
	"uptime-monitor/config"
	"uptime-monitor/monitor"
)

// Loads config - located at ./config/service-targets.yaml
// Starts both processes
// Handles basic logging
func main() {

	// Load configuration
	targets, err := config.LoadConfig("config/service-targets.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
		return
	}

	// Start monitoring
	go monitor.StartMonitoring(targets.Targets)

	// Start API server
	go api.StartServer(":8080")

  // Log the number of targets loaded
  log.Printf("Uptime monitoring service started with %d targets", len(targets.Targets))

	// Block forever
	select {}
}
