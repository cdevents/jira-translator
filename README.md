# Jira Translator CDEvents
A translator plugin for translating [Jira issue and worklog events](https://developer.atlassian.com/platform/forge/events-reference/jira/#jira-events) into [Continuous Operations Ticket events](https://github.com/cdevents/spec/blob/v0.4.1/continuous-operations.md#ticket).
This plugin is served using Hashicorp's [go-plugin](https://github.com/hashicorp/go-plugin/). 

The binary of this plugin is published with a release URL and is used by external applications like [cdevents/webhook-adapter](https://github.com/cdevents/webhook-adapter)

The published plugin's binary can be downloaded and loaded by creating a new plugin client using HashiCorp's go-plugin, which manages the lifecycle of this plugin and establishes the RPC connection.

### How to build locally
Run the `make` command from the project root directory, which creates a plugin's binary with the name `jira-translator-cdevents`
````make
make all
````

### Jira-CDEvents type mapping for translation
Below are the Jira events that currently have mappings with CDEvents and are supported by this translator.

| CDEvent Type  | Jira Event Type  |
| :------------ |:-------------------|
|  dev.cdevents.ticket.created| issue_created |
|  dev.cdevents.ticket.updated   | issue_updated    |
|  dev.cdevents.ticket.updated   | issue_commented    |
|  dev.cdevents.ticket.updated   | issue_assigned    |
|  dev.cdevents.ticket.updated   | issue_work_logged    |
|  dev.cdevents.ticket.updated   | issue_worklog_updated    |
|  dev.cdevents.ticket.updated   | issue_worklog_deleted    |
|  dev.cdevents.ticket.closed    | issue_generic     |

Note: 
CDEvents do not have corresponding mapping with Jira Project, User, Configuration and Software related events, these events can be mapped using [Custom CDEvents](https://github.com/cdevents/spec/tree/v0.4.1/custom)