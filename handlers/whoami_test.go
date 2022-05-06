package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhoami(t *testing.T) {
	w := httptest.NewRecorder()
	ipaddress := GetOutboundIP().String()
	h := whoami()
	h(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	info := struct {
		CurrentTime string `json:"currentTime"`
		IPaddress   string `json:"ipaddress"`
	}{}
	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.IPaddress != ipaddress {
		t.Errorf("Release IPaddress is wrong. Have: %s, want: %s", info.IPaddress, ipaddress)
	}
}
