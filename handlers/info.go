package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time" 
)

func info(release string, birth time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Print("info called")

		t := time.Now()
		age := t.Sub(birth)

		info := struct {
			CurrentTime string `json:"currentTime"`
			IPaddress   string `json:"ipaddress"`
			Release     string `json:"release"`
			Age         string `json:"age"`
		}{
			t.String(), release, GetOutboundIP().String(), age.String(),
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
