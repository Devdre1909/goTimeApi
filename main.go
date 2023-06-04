package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type TimeResponse struct {
	CurrentTime time.Time `json:"current_time"`
}

type TzDatabase struct {
	Identifier string
}

func main() {
	fmt.Println("Assignment 1")
	fmt.Println("Time API: Master the concept of microservices API development using the Hexagonal Architecture in Go")

	router := mux.NewRouter()

	router.HandleFunc("/api/time", getTime)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func getTime(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	tz := queryParams.Get("tz")

	if len(strings.TrimSpace(tz)) < 1 {

		loc, err := time.LoadLocation(tz)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			response := TimeResponse{
				CurrentTime: time.Now().In(loc),
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	} else {

		response := make(map[string]string, 0)
		timeZone := strings.Split(tz, ",")
		for _, tzDb := range timeZone {
			loc, err := time.LoadLocation(tzDb)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone  %s in list of input", tzDb)))
				return
			}
			now := time.Now().In(loc)
			response[tzDb] = now.String()

		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}

}
