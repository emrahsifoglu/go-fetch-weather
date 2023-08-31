package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/emrahsifoglu/go-fetch-weather/pkg/weather"
)

type Alerts struct {
    Alerts []string `json:"alerts"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    area := query.Get("area")
    client := weather.NewClient("https://api.weather.gov", time.Second * 10)

    fc, err := client.GetAlerts(area)
    if err != nil {
        log.Panicf("Error occurred in getting feature collection Err: %s", err)
    }

    var alerts []string

    for _, feature := range fc.Features {
        alerts = append(alerts, feature.Properties["headline"].(string))
    }

    jsondat := &Alerts{Alerts: alerts}

    encjson, err := json.Marshal(jsondat)
    if err != nil {
        log.Fatalf("Error happened in JSON marshal. Err: %s", err)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write(encjson)
}

func main() {
    http.HandleFunc("/alerts", handleRequest)
    http.ListenAndServe(":9090", nil)
}
