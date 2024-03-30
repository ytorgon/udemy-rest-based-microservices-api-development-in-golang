package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const contentType = "Content-Type"
const applicationJson = "application/json"

type Map = map[string]string

type Time struct {
	currentTime string `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(Map, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")
	fmt.Println(timezones)

	if len(timezones) <= 1 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			sendInvalid(w, tz)
		} else {
			response["current_time"] = time.Now().In(loc).String()
			jsonEncode(w, response)
		}
	} else {
		for _, tz := range timezones {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				sendInvalid(w, tz)
				return
			}
			response[tz] = time.Now().In(loc).String()
		}
		jsonEncode(w, response)
	}
}

func sendInvalid(w http.ResponseWriter, tz string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
}

func jsonEncode(w http.ResponseWriter, response Map) {
	w.Header().Add(contentType, applicationJson)
	json.NewEncoder(w).Encode(response)
}
