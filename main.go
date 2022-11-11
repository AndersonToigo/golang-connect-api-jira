package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AndersonToigo/golang-connect-api-jira/helpers"
	"github.com/AndersonToigo/golang-connect-api-jira/integration"
)

const (
	email     = "ander.toigo@gmail.com"
	token     = "6nANkF36aNCSU3KsxIUd408D"
	accountID = "636d8c6cc383ad842146ed20"
)

func main() {

	authorization := helpers.BasicAuth(email, token)
	equalData := helpers.NowDate()

	var jiraResponse integration.ResponseJiraIssue

	jql := fmt.Sprintf(`{"jql": "worklogDate>='%v' and worklogDate<='%v' and (worklogAuthor in ('%v'))", "fields":["worklog"] }`, equalData, equalData, accountID)

	payload := strings.NewReader(jql)

	result, err := integration.RequestHttpJiraReport(authorization, payload)

	if err != nil {
		panic(err.Error())
	}

	er := json.Unmarshal(result, &jiraResponse)

	if er != nil {
		panic(er.Error())
	}

	var count = 0.0
	for _, r := range jiraResponse.Issues {
		for _, a := range r.Fields.Worklog.Worklogs {
			count += float64(a.TimeSpentSeconds)
		}
	}

	fmt.Printf("Suas horas trabalhadas: %.2f", helpers.ConvertHour(count))
}
