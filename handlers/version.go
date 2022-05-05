package handlers

import (
	"encoding/json"
	"log" //"github.com/golang/glog"
	"net/http"
)

// home returns a simple HTTP handler function which writes a response.
func version(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Print("home called")
		info := struct {
			Build time string `json:"buildTime"`
			Commit id  string `json:"commit"`
			Release    string `json:"release"`
		}{
			buildTime, commit, release,
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
