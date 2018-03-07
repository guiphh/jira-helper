package jclient

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

type MyIssue struct {
	Key      string
	Summary  string
	EpicKey  string
	EpicName string
}

var JIRACLIENT *jira.Client

func Connect(server, username, password string) {
	jiraClient, err := jira.NewClient(nil, server)
	if err != nil {
		panic(err)
	}
	jiraClient.Authentication.SetBasicAuth(username, password)

	JIRACLIENT = jiraClient
}

func GetCurrentSprint(BoardId string) (string, error) {
	sprints, _, err := JIRACLIENT.Board.GetAllSprints(BoardId)

	if err != nil {
		fmt.Println(err)
	}

	for _, sprint := range sprints {
		if sprint.State == "active" {
			return sprint.Name, nil
		}
	}

	return "", fmt.Errorf("No Active sprint found")

}

// GetIssuesSprint returns a list of customized issues
func GetIssuesSprint() []MyIssue {
	issues, _, err := JIRACLIENT.Sprint.GetIssuesForSprint(1)
	if err != nil {
		panic(err)
	}

	var myIssues []MyIssue

	for _, issue := range issues {
		//customFields, _, err := JIRACLIENT.Issue.GetCustomFields(issue.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

		var iss MyIssue

		iss.Key = issue.Key
		iss.Summary = issue.Fields.Summary
		if issue.Fields.Epic != nil {
			iss.EpicKey = issue.Fields.Epic.Key
			iss.EpicName = issue.Fields.Epic.Name
		}
		myIssues = append(myIssues, iss)
	}

	return myIssues
}
