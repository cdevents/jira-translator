/*
Copyright (C) 2024 Nordix Foundation.
For a full list of individual contributors, please see the commit history.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
SPDX-License-Identifier: Apache-2.0
*/

package jira

import (
	"encoding/json"
)

func (pEvent *JiraEvent) HandleIssueCreatedEvent() (string, error) {
	var issueCreatedEvent IssueCreatedEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueCreatedEvent)
	if err != nil {
		Log().Error("Error occurred while Unmarshal JSON event into JiraIssueEvent struct", err)
		return "", err
	}
	Log().Info("jira issue_created event received : ", issueCreatedEvent.Issue.Fields.Issuetype.Name, issueCreatedEvent.User.Name, issueCreatedEvent.IssueEventTypeName)
	cdEvent, err := issueCreatedEvent.TicketCreatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func (pEvent *JiraEvent) HandleIssueUpdatedEvent() (string, error) {
	var issueUpdatedEvent IssueUpdatedEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueUpdatedEvent)
	if err != nil {
		Log().Error("Error occurred while Unmarshal JSON event into JiraIssueEvent struct", err)
		return "", err
	}
	Log().Info("jira issue_updated event received : ", issueUpdatedEvent.Issue.Fields.Issuetype.Name, issueUpdatedEvent.User.Name, issueUpdatedEvent.IssueEventTypeName)
	cdEvent, err := issueUpdatedEvent.TicketUpdatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func (pEvent *JiraEvent) HandleIssueCommentedEvent() (string, error) {
	var issueCommentedEvent IssueCommentedEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueCommentedEvent)
	if err != nil {
		Log().Error("Error occurred while Unmarshal JSON event into JiraIssueEvent struct", err)
		return "", err
	}
	Log().Info("jira issue_commented event received : ", issueCommentedEvent.Issue.Fields.Issuetype.Name, issueCommentedEvent.User.Name, issueCommentedEvent.IssueEventTypeName)
	cdEvent, err := issueCommentedEvent.TicketUpdatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func (pEvent *JiraEvent) HandleIssueAssignedEvent() (string, error) {
	var issueAssignedEvent IssueAssignedEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueAssignedEvent)
	if err != nil {
		Log().Error("Error occurred while Unmarshal JSON event into JiraIssueEvent struct", err)
		return "", err
	}
	Log().Info("jira issue_assigned event received : ", issueAssignedEvent.Issue.Fields.Issuetype.Name, issueAssignedEvent.User.Name, issueAssignedEvent.IssueEventTypeName)
	cdEvent, err := issueAssignedEvent.TicketUpdatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}
