package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func die(release string, birth time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Print("die called")

		t := time.Now()
		age := t.Sub(birth)

		info := struct {
			CurrentTime string `json:"currentTime"`
			IPaddress   string `json:"ipaddress"`
			Release     string `json:"release"`
			Age         string `json:"age"`
		}{
			t.Format(time.RubyDate),
			GetOutboundIP().String(),
			release,
			age.Truncate(time.Second).String(),
		}

		body, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		timeToDie = new(time.Time)
		//timeToDie = time.Now()
	}
}
