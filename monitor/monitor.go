// Core logic for pinging

package monitor

import (
	"log"
	"net/http"
	"time"
	"uptime-monitor/config"
)

func StartMonitoring(targets []config.Target) {
	for _, target := range targets {
		go monitorTarget(target)
	}
}

func monitorTarget(target config.Target) {
	interval := time.Duration(target.Interval) * time.Second

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	client := http.Client{
		Timeout: interval,
	}

	for {
		<-ticker.C
		start := time.Now()
		resp, err := client.Get(target.URL)
		duration := time.Since(start)

		if err != nil {
			RecordFailed(target.Name)
			log.Printf("Fail: %s (%v)", target.Name, err)
			continue
		}

		resp.Body.Close()
		RecordSuccess(target.Name, duration.Seconds())
		log.Printf("UP: %s - %s in %v", target.Name, resp.Status, duration)
	}
}
