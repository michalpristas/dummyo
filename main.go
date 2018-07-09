package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/eventlog/", event)
	http.HandleFunc("/eventlog/1", event1)

	http.HandleFunc("/cachemiss", cachemiss)

	http.HandleFunc("/moduxe/@v/v1.0.mod", mod)
	http.HandleFunc("/moduxe/@v/v1.0.zip", zip)

	http.ListenAndServe(":8081", nil)
}

func cachemiss(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cm")
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	mod := Module{}
	d.Decode(&mod)

	fmt.Println(mod)
}

func event(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event")
	var ee []Event
	ee = append(ee, Event{ID: "1", Time: time.Now(), Module: "moduxe", Version: "v1.0"})

	j, _ := json.Marshal(ee)
	w.Write(j)
}

func event1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event1")

	var ee []Event

	j, _ := json.Marshal(ee)
	w.Write(j)

}

func mod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mod")
	m := []byte("mod")
	w.Write(m)
}

func zip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("zip")
	z := []byte("zip")
	w.Write(z)
}

type Module struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Event is entry of event log specifying demand for a module.
type Event struct {
	// ID is identifier, also used as a pointer reference target.
	ID string `json:"_id" bson:"_id,omitempty"`
	// Time is cache-miss created/handled time.
	Time time.Time `json:"time_created" bson:"time_created"`
	// Module is module name.
	Module string `json:"module" bson:"module"`
	// Version is version of a module e.g. "1.10", "1.10-deprecated"
	Version string `json:"version" bson:"version"`
}
