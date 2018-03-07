package main

import (
	"fmt"

	"github.com/guiphh/jira-helper/pkg/excel"
	"github.com/guiphh/jira-helper/pkg/jclient"
)

func main() {
	jclient.Connect("https://server.net", "username", "password")
	board := "1"
	sprintName, err := jclient.GetCurrentSprint(board)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sprintName)
	}
	issues := jclient.GetIssuesSprint()

	excel.WriteXlsx("MyXLSXFile.xlsx", issues)
}
