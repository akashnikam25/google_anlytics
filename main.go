package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mileusna/useragent"
)

type TrackingData struct {
	Type          string `json:"type"`
	Identity      string `json:"identity"`
	UserAgent     string `json:"ua"`
	Event         string `json:"event"`
	Category      string `json:"category"`
	Referrer      string `json:"referrer"`
	IsTouchDevice bool   `json:"isTouchDevice"`
	OccuredAt     uint32
}

type Tracking struct {
	SiteID string       `json:"site_id"`
	Action TrackingData `json:"tracking"`
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	defer w.WriteHeader(http.StatusOK)

	data := r.URL.Query().Get("data")
	fmt.Println("Base Encoded data ", data)
	trk, err := decodeData(data)
	if err != nil {
		fmt.Print(err)
	}
	ua := useragent.Parse(trk.Action.UserAgent)
	if err = e.Add(trk, ua); err != nil {
		fmt.Println(err)
	}
	fmt.Println("site id", trk.SiteID)
}

func decodeData(s string) (data Tracking, err error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return

}

func main() {
	err := e.open()
	if err != nil {
		log.Fatal(err)
		return
	}
	http.HandleFunc("/track", trackHandler)
	http.ListenAndServe(":9876", nil)
}
