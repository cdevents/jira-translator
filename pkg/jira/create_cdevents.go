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
	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cdevents04 "github.com/cdevents/sdk-go/pkg/api/v04"
)

func (issueCreatedEvent *IssueCreatedEvent) TicketCreatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.created for ", issueCreatedEvent.IssueEventTypeName)
	cdEvent, err := cdevents04.NewTicketCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketCreatedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(issueCreatedEvent.Issue.Fields.Status.IconURL)
	cdEvent.SetSubjectId(issueCreatedEvent.Issue.ID)
	cdEvent.SetSubjectUri(issueCreatedEvent.Issue.Self)
	cdEvent.SetSubjectSummary(issueCreatedEvent.Issue.Fields.Summary)
	cdEvent.SetSubjectCreator(issueCreatedEvent.Issue.Fields.Creator.Name)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketCreatedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}
func (issueUpdatedEvent *IssueUpdatedEvent) TicketUpdatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.updated for ", issueUpdatedEvent.IssueEventTypeName)
	cdEvent, err := cdevents04.NewTicketUpdatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketUpdatedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(issueUpdatedEvent.Issue.Fields.Status.IconURL)
	cdEvent.SetSubjectId(issueUpdatedEvent.Issue.ID)
	cdEvent.SetSubjectUri(issueUpdatedEvent.Issue.Self)
	cdEvent.SetSubjectUpdatedBy(issueUpdatedEvent.Issue.Fields.Creator.Name)
	cdEvent.SetCustomData("application/json", issueUpdatedEvent.Changelog)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketUpdatedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}

func (issueCommentedEvent *IssueCommentedEvent) TicketUpdatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.updated for ", issueCommentedEvent.IssueEventTypeName)
	cdEvent, err := cdevents04.NewTicketUpdatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketUpdatedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(issueCommentedEvent.Issue.Fields.Status.IconURL)
	cdEvent.SetSubjectId(issueCommentedEvent.Issue.ID)
	cdEvent.SetSubjectUri(issueCommentedEvent.Issue.Self)
	cdEvent.SetSubjectUpdatedBy(issueCommentedEvent.Issue.Fields.Creator.Name)
	cdEvent.SetCustomData("application/json", issueCommentedEvent.Comment)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketUpdatedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}

func (issueAssignedEvent *IssueAssignedEvent) TicketUpdatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.updated for ", issueAssignedEvent.IssueEventTypeName)
	cdEvent, err := cdevents04.NewTicketUpdatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketUpdatedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(issueAssignedEvent.Issue.Fields.Status.IconURL)
	cdEvent.SetSubjectId(issueAssignedEvent.Issue.ID)
	cdEvent.SetSubjectUri(issueAssignedEvent.Issue.Self)
	cdEvent.SetSubjectUpdatedBy(issueAssignedEvent.Issue.Fields.Creator.Name)
	cdEvent.SetCustomData("application/json", issueAssignedEvent.Changelog)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketUpdatedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}
