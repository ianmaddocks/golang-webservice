package handlers

import {
  "net/http"
  "time"
}

// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
duration := time.Now().Sub(started)
    if duration.Seconds() > 10 {
        w.WriteHeader(500)
        w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
    } else {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    }

	//w.WriteHeader(http.StatusOK)
}
