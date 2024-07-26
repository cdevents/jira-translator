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
	Log().Info("Creating dev.cdevents.ticket.created for Jira ", issueCreatedEvent.IssueEventTypeName)
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
	Log().Info("Creating dev.cdevents.ticket.updated for Jira ", issueUpdatedEvent.IssueEventTypeName)
	return createTicketUpdatedCDEvent(
		issueUpdatedEvent.Issue.Fields.Status.IconURL,
		issueUpdatedEvent.Issue.ID,
		issueUpdatedEvent.Issue.Self,
		issueUpdatedEvent.Issue.Fields.Creator.Name,
		issueUpdatedEvent.Changelog)
}

func (issueCommentedEvent *IssueCommentedEvent) TicketUpdatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.updated for Jira ", issueCommentedEvent.IssueEventTypeName)
	return createTicketUpdatedCDEvent(
		issueCommentedEvent.Issue.Fields.Status.IconURL,
		issueCommentedEvent.Issue.ID,
		issueCommentedEvent.Issue.Self,
		issueCommentedEvent.Issue.Fields.Creator.Name,
		issueCommentedEvent.Comment)
}

func (issueAssignedEvent *IssueAssignedEvent) TicketUpdatedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.updated for Jira ", issueAssignedEvent.IssueEventTypeName)
	return createTicketUpdatedCDEvent(
		issueAssignedEvent.Issue.Fields.Status.IconURL,
		issueAssignedEvent.Issue.ID,
		issueAssignedEvent.Issue.Self,
		issueAssignedEvent.Issue.Fields.Creator.Name,
		issueAssignedEvent.Changelog)
}

func createTicketUpdatedCDEvent(source, id, uri, updatedBy string, customData interface{}) (string, error) {
	cdEvent, err := cdevents04.NewTicketUpdatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketUpdatedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(source)
	cdEvent.SetSubjectId(id)
	cdEvent.SetSubjectUri(uri)
	cdEvent.SetSubjectUpdatedBy(updatedBy)
	cdEvent.SetCustomData("application/json", customData)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketUpdatedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}

func (issueGenericEvent *IssueGenericEvent) TicketClosedCDEvent() (string, error) {
	Log().Info("Creating dev.cdevents.ticket.closed for Jira ", issueGenericEvent.IssueEventTypeName)
	cdEvent, err := cdevents04.NewTicketClosedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent TicketClosedEvent, ", err)
		return "", err
	}

	cdEvent.SetSource(issueGenericEvent.Issue.Fields.Status.IconURL)
	cdEvent.SetSubjectId(issueGenericEvent.Issue.ID)
	cdEvent.SetSubjectUri(issueGenericEvent.Issue.Self)
	cdEvent.SetSubjectResolution(issueGenericEvent.Issue.Fields.Resolution.Name)
	cdEvent.SetSubjectSummary(issueGenericEvent.Issue.Fields.Summary)
	cdEvent.SetSubjectCreator(issueGenericEvent.Issue.Fields.Creator.Name)
	cdEvent.SetCustomData("application/json", issueGenericEvent.Changelog)

	cdEventStr, err := cdevents.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating TicketClosedEvent CDEvent as Json string, ", err)
		return "", err
	}

	return cdEventStr, nil
}
