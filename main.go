package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/guiphh/jira-helper/pkg/jclient"
)

func issueHandler(w http.ResponseWriter, r *http.Request) {
	issueKeys, ok := r.URL.Query()["key"]
	if !ok || len(issueKeys) < 1 {
		log.Println("Url Param 'key' is missing ")
		return
	}

	key := issueKeys[0]
	log.Println("key is " + string(key))

	issue := jclient.GetIssue(key)
	js, err := json.Marshal(issue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(js)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	jclient.Connect("https://server.net", "username", "password")
	board := "1"
	sprint, err := jclient.GetCurrentSprint(board)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sprint)
	}
	issues := jclient.GetIssuesSprint(sprint.ID)

	js, err := json.Marshal(issues)
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
