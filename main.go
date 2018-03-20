package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Sprint struct {
	Id   int8   `json: "id"`
	Name string `json: "name"`
}

type Issue struct {
	Key     string `json: "key"`
	Summary string `json: "summary"`
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]Issue)

	m["TAA-1"] = Issue{Key: "TAA-1", Summary: "Summary 1"}
	m["TAA-2"] = Issue{Key: "TAA-2", Summary: "Summary 2"}
	m["TAA-3"] = Issue{Key: "TAA-3", Summary: "Summary 3"}

	issueKeys, ok := r.URL.Query()["key"]
	if !ok || len(issueKeys) < 1 {
		log.Println("Url Param 'key' is missing ")
		return
	}

	key := issueKeys[0]
	log.Println("key is " + string(key))

	js, err := json.Marshal(m[key])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(js)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	issue1 := Issue{Key: "TAA-1", Summary: "Issue 1"}
	issue2 := Issue{Key: "TAA-2", Summary: "Issue 2"}

	var iss []Issue
	iss = append(iss, issue1)
	iss = append(iss, issue2)

	js, err := json.Marshal(iss)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(js)
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/issue", issueHandler)
	http.ListenAndServe(":8080", nil)

}
