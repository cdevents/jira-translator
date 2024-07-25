package jira

import (
	"net/http"
	"os"
	"testing"
)

func TestHandleTranslateIssueCreatedEvent(t *testing.T) {
	event, err := os.ReadFile("testdata/issue-created.json")
	if err != nil {
		t.Fatalf("Failed to read issue-created.json file: %v", err)
	}
	headers := http.Header{}
	headers.Set("X-Origin-Url", "http://jira.est.tech")

	cdEvent, err := HandleTranslateJiraEvent(string(event), headers)
	if err != nil {
		t.Errorf("Expected ticket.created CDEvent to be successful.")
		return
	}
	Log().Info("Handle issue-created gerrit event into dev.cdevents.ticket.created is successful ", cdEvent)
}
