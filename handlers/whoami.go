package handlers

import (
	"encoding/json"
	"log" //"github.com/golang/glog"
	"net"
	"net/http"
	"time"
)

// home returns a simple HTTP handler function which writes a response.
func whoami() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Print("whoami called")
		info := struct {
			CurrentTime string `json:"currentTime"`
			IPaddress   string `json:"ipaddress"`
		}{
			time.Now().String(), GetOutboundIP().String(),
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

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
