// Serves metrics

package api

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthz", healthCheckHandler)

	log.Printf("API server listening on port %s", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	// or
	// var buffer bytes.Buffer
	// buffer.WriteString("OK")
	// w.Write(buffer.Bytes())

	// or
	// fmt.Fprint(w, "OK")

	// or
	// fmt.Fprintf(w, "Status: %s\n", "OK")

	// or
	// io.WriteString(w, "OK")
}
