package jclient

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

// Issue is a wrapper with all the required fields
type Issue struct {
	Key      string
	Summary  string
	EpicKey  string
	EpicName string
}

// Sprint wrapper
type Sprint struct {
	ID   int
	Name string
}

// JIRACLIENT allows to maintain connection for rest calls
var JIRACLIENT *jira.Client

func Connect(server, username, password string) {
	jiraClient, err := jira.NewClient(nil, server)
	if err != nil {
		panic(err)
	}
	jiraClient.Authentication.SetBasicAuth(username, password)

	JIRACLIENT = jiraClient
}

// GetCurrentSprint returns a pointer to the current sprint
// This is the current sprint of the board passed in parameter
func GetCurrentSprint(BoardID string) (*Sprint, error) {
	sprints, _, err := JIRACLIENT.Board.GetAllSprints(BoardID)

	if err != nil {
		fmt.Println(err)
	}

	currentSprint := &Sprint{}
	for _, sprint := range sprints {
		if sprint.State == "active" {
			currentSprint.ID = sprint.ID
			currentSprint.Name = sprint.Name
		}
	}

	if currentSprint != nil {
		return currentSprint, nil
	}

	return nil, fmt.Errorf("No Active sprint found")

}

// GetIssuesSprint returns a list of customized issues
func GetIssuesSprint(sprintID int) []Issue {
	issues, _, err := JIRACLIENT.Sprint.GetIssuesForSprint(sprintID)
	if err != nil {
		panic(err)
	}

	var myIssues []Issue

	for _, issue := range issues {
		//customFields, _, err := JIRACLIENT.Issue.GetCustomFields(issue.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

		var iss Issue

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

func GetIssue(IssueID string) Issue {
	issue, _, err := JIRACLIENT.Issue.Get(IssueID, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	var iss Issue

	iss.Key = issue.Key
	iss.Summary = issue.Fields.Summary
	if issue.Fields.Epic != nil {
		iss.EpicKey = issue.Fields.Epic.Key
		iss.EpicName = issue.Fields.Epic.Name
	}

	return iss
}
