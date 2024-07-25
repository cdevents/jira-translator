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
		Log().Error("Error occurred while Unmarshal issue_created JSON event into IssueCreatedEvent struct", err)
		return "", err
	}
	Log().Infof("jira %s event translating to dev.cdevents.ticket.created", issueCreatedEvent.IssueEventTypeName)
	cdEvent, err := issueCreatedEvent.TicketCreatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func (pEvent *JiraEvent) HandleIssueUpdatedEvent(eventType string) (string, error) {
	var issueUpdatedEvent IssueUpdatedEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueUpdatedEvent)
	if err != nil {
		Log().Errorf("Error occurred while Unmarshal %s JSON event into IssueUpdatedEvent struct, %v", eventType, err)
		return "", err
	}
	Log().Infof("jira %s event translating to dev.cdevents.ticket.updated", issueUpdatedEvent.IssueEventTypeName)
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
		Log().Error("Error occurred while Unmarshal issue_commented JSON event into IssueCommentedEvent struct ", err)
		return "", err
	}
	Log().Infof("jira %s event translating to dev.cdevents.ticket.updated", issueCommentedEvent.IssueEventTypeName)
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
		Log().Error("Error occurred while Unmarshal issue_assigned JSON event into IssueAssignedEvent struct", err)
		return "", err
	}
	Log().Infof("jira %s event translating to dev.cdevents.ticket.updated", issueAssignedEvent.IssueEventTypeName)
	cdEvent, err := issueAssignedEvent.TicketUpdatedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func (pEvent *JiraEvent) HandleIssueGenericEvent() (string, error) {
	var issueGenericEvent IssueGenericEvent
	err := json.Unmarshal([]byte(pEvent.Event), &issueGenericEvent)
	if err != nil {
		Log().Error("Error occurred while Unmarshal issue_generic JSON event into IssueGenericEvent struct", err)
		return "", err
	}
	Log().Infof("jira %s event translating to dev.cdevents.ticket.closed", issueGenericEvent.IssueEventTypeName)
	cdEvent, err := issueGenericEvent.TicketClosedCDEvent()
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}
