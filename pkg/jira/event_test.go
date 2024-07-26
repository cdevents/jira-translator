package jira

import (
	"net/http"
	"os"
	"testing"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdevents04 "github.com/cdevents/sdk-go/pkg/api/v04"
	assert "github.com/go-playground/assert/v2"
)

func TestHandleTranslateIssueCreatedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_created.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_created into dev.cdevents.ticket.created to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.created")
}

func TestHandleTranslateIssueUpdatedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_updated.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_updated into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueCommentedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_commented.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_commented into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueAssignedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_assigned.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_assigned into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueWorkLoggedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_work_logged.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_work_logged into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueWorklogUpdatedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_worklog_updated.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_worklog_updated into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueWorklogDeletedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_worklog_deleted.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_worklog_deleted into dev.cdevents.ticket.updated to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10023")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.updated")
}

func TestHandleTranslateIssueGenericClosedEvent(t *testing.T) {
	newCDEvent, err := translateJiraEvent("testdata/issue_generic.json")
	if err != nil {
		t.Errorf("Expected translate Jira issue_generic into dev.cdevents.ticket.closed to be successful")
		return
	}
	assert.Equal(t, newCDEvent.GetSource(), "http://localhost:9090/")
	assert.Equal(t, newCDEvent.GetSubjectId(), "10008")
	assert.Equal(t, newCDEvent.GetSubject().GetSubjectType().String(), "ticket")
	assert.Equal(t, newCDEvent.GetType().UnversionedString(), "dev.cdevents.ticket.closed")
}

func translateJiraEvent(testFile string) (cdevents.CDEventV04, error) {
	event, err := os.ReadFile(testFile)
	if err != nil {
		Log().Errorf("Failed to read %s file: %v", testFile, err)
		return nil, err
	}
	headers := http.Header{}
	headers.Set("X-Origin-Url", "http://jira.est.tech")

	cdEvent, err := HandleTranslateJiraEvent(string(event), headers)
	if err != nil {
		Log().Errorf("Failed to translate Jira Event into CDEvent: %v", err)
		return nil, err
	}
	newCDEvent, err := cdevents04.NewFromJsonString(cdEvent)
	if err != nil {
		Log().Errorf("Failed to create CDEvent as Json string: %v", err)
		return nil, err
	}
	return newCDEvent, nil
}
