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

// Define structs for nested JSON objects
type AvatarUrls struct {
	Four8X48  string `json:"48x48"`
	Two4X24   string `json:"24x24"`
	One6X16   string `json:"16x16"`
	Three2X32 string `json:"32x32"`
}
type User struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type Issuetype struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}
type Project struct {
	Self           string     `json:"self"`
	ID             string     `json:"id"`
	Key            string     `json:"key"`
	Name           string     `json:"name"`
	ProjectTypeKey string     `json:"projectTypeKey"`
	AvatarUrls     AvatarUrls `json:"avatarUrls"`
}
type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}
type Priority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}
type Assignee struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}
type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}
type Timetracking struct {
}
type Creator struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type Reporter struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type Aggregateprogress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}
type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}
type Author struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type UpdateAuthor struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}
type Comments struct {
	Self         string       `json:"self"`
	ID           string       `json:"id"`
	Author       Author       `json:"author"`
	Body         string       `json:"body"`
	UpdateAuthor UpdateAuthor `json:"updateAuthor"`
	Created      string       `json:"created"`
	Updated      string       `json:"updated"`
}
type Comment struct {
	Comments   []Comments `json:"comments"`
	MaxResults int        `json:"maxResults"`
	Total      int        `json:"total"`
	StartAt    int        `json:"startAt"`
}
type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}
type Worklog struct {
	StartAt    int   `json:"startAt"`
	MaxResults int   `json:"maxResults"`
	Total      int   `json:"total"`
	Worklogs   []any `json:"worklogs"`
}
type Fields struct {
	Issuetype                     Issuetype         `json:"issuetype"`
	Timespent                     any               `json:"timespent"`
	Project                       Project           `json:"project"`
	FixVersions                   []any             `json:"fixVersions"`
	Customfield10110              string            `json:"customfield_10110"`
	Customfield10111              any               `json:"customfield_10111"`
	Aggregatetimespent            any               `json:"aggregatetimespent"`
	Resolution                    any               `json:"resolution"`
	Customfield10107              any               `json:"customfield_10107"`
	Customfield10108              any               `json:"customfield_10108"`
	Customfield10109              any               `json:"customfield_10109"`
	Resolutiondate                any               `json:"resolutiondate"`
	Workratio                     int               `json:"workratio"`
	LastViewed                    any               `json:"lastViewed"`
	Watches                       Watches           `json:"watches"`
	Created                       string            `json:"created"`
	Priority                      Priority          `json:"priority"`
	Customfield10100              any               `json:"customfield_10100"`
	Customfield10101              any               `json:"customfield_10101"`
	Customfield10102              any               `json:"customfield_10102"`
	Labels                        []any             `json:"labels"`
	Customfield10103              any               `json:"customfield_10103"`
	Timeestimate                  any               `json:"timeestimate"`
	Aggregatetimeoriginalestimate any               `json:"aggregatetimeoriginalestimate"`
	Versions                      []any             `json:"versions"`
	Issuelinks                    []any             `json:"issuelinks"`
	Assignee                      Assignee          `json:"assignee"`
	Updated                       string            `json:"updated"`
	Status                        Status            `json:"status"`
	Components                    []any             `json:"components"`
	Timeoriginalestimate          any               `json:"timeoriginalestimate"`
	Description                   string            `json:"description"`
	Timetracking                  Timetracking      `json:"timetracking"`
	Archiveddate                  any               `json:"archiveddate"`
	Attachment                    []any             `json:"attachment"`
	Aggregatetimeestimate         any               `json:"aggregatetimeestimate"`
	Summary                       string            `json:"summary"`
	Creator                       Creator           `json:"creator"`
	Subtasks                      []any             `json:"subtasks"`
	Reporter                      Reporter          `json:"reporter"`
	Customfield10000              string            `json:"customfield_10000"`
	Aggregateprogress             Aggregateprogress `json:"aggregateprogress"`
	Environment                   any               `json:"environment"`
	Duedate                       any               `json:"duedate"`
	Progress                      Progress          `json:"progress"`
	Comment                       Comment           `json:"comment"`
	Votes                         Votes             `json:"votes"`
	Worklog                       Worklog           `json:"worklog"`
	Archivedby                    any               `json:"archivedby"`
}
type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}
type Items struct {
	Field      string `json:"field"`
	Fieldtype  string `json:"fieldtype"`
	From       any    `json:"from"`
	FromString string `json:"fromString"`
	To         any    `json:"to"`
	ToString   string `json:"toString"`
}
type Changelog struct {
	ID    string  `json:"id"`
	Items []Items `json:"items"`
}

// Jira event types
type IssueCreatedEvent struct {
	Timestamp          int64  `json:"timestamp"`
	WebhookEvent       string `json:"webhookEvent"`
	IssueEventTypeName string `json:"issue_event_type_name"`
	User               User   `json:"user"`
	Issue              Issue  `json:"issue"`
}
type IssueUpdatedEvent struct {
	Timestamp          int64     `json:"timestamp"`
	WebhookEvent       string    `json:"webhookEvent"`
	IssueEventTypeName string    `json:"issue_event_type_name"`
	User               User      `json:"user"`
	Issue              Issue     `json:"issue"`
	Changelog          Changelog `json:"changelog"`
}
type IssueCommentedEvent struct {
	Timestamp          int64    `json:"timestamp"`
	WebhookEvent       string   `json:"webhookEvent"`
	IssueEventTypeName string   `json:"issue_event_type_name"`
	User               User     `json:"user"`
	Issue              Issue    `json:"issue"`
	Comment            Comments `json:"comment"`
}
type IssueAssignedEvent struct {
	Timestamp          int64     `json:"timestamp"`
	WebhookEvent       string    `json:"webhookEvent"`
	IssueEventTypeName string    `json:"issue_event_type_name"`
	User               User      `json:"user"`
	Issue              Issue     `json:"issue"`
	Changelog          Changelog `json:"changelog"`
	Comment            Comments  `json:"comment"`
}
