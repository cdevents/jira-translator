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
	"fmt"
	"net/http"
)

type JiraEvent struct {
	Event string
}

func NewJiraEvent(event string) (pEvent *JiraEvent) {
	pEvent = &JiraEvent{event}
	return
}

func HandleTranslateJiraEvent(event string, header http.Header) (string, error) {
	Log().Debug("Handle translation into CDEvent from Jira event %s\n", event)
	jiraEvent := NewJiraEvent(event)
	cdEvent, err := jiraEvent.TranslateIntoCDEvent()
	if err != nil {
		Log().Error("Error translating Jira event into CDEvent, ", err)
		return "", err
	}
	Log().Info("Jira Event translated into CDEvent", cdEvent)
	return cdEvent, nil
}

func (pEvent *JiraEvent) TranslateIntoCDEvent() (string, error) {
	eventMap := make(map[string]interface{})
	cdEvent := ""
	err := json.Unmarshal([]byte(pEvent.Event), &eventMap)
	if err != nil {
		Log().Error("Error occurred while Unmarshal jiraEvent data into eventMap", err)
		return "", err
	}
	eventType := eventMap["issue_event_type_name"]
	Log().Info("handling translating to CDEvent from Jira event type name: ", eventType)

	switch eventType {
	case "issue_created":
		cdEvent, err = pEvent.HandleIssueCreatedEvent()
		if err != nil {
			return "", err
		}
	case "issue_updated":
		cdEvent, err = pEvent.HandleIssueUpdatedEvent()
		if err != nil {
			return "", err
		}
	case "issue_commented":
		cdEvent, err = pEvent.HandleIssueCommentedEvent()
		if err != nil {
			return "", err
		}
	case "issue_assigned":
		cdEvent, err = pEvent.HandleIssueAssignedEvent()
		if err != nil {
			return "", err
		}
	default:
		Log().Info("Not handling CDEvent translation for Jira event type: ", eventType)
		return "", fmt.Errorf("jira event type %q, not supported for translation", eventType)
	}
	return cdEvent, nil
}
