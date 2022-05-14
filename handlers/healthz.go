package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
	if timeToDie != nil {
		log.Print("Healthz called and its timeToDie")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("time to die")))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
