# Migration Guide: v0.14 to v0.15

v0.15.0 introduces a number of API changes, which should be simple to migrate.
Just follow this guide and if you still encounter problems, ask for help on Discord or feel free to create an issue.

<!-- toc -->

-   [Changed Struct Fields (#503) (#520)](#changed-struct-fields)

<!-- tocstop -->

## Changed Struct Fields

 - The `State` field at **NotificationSubject** changed from **StateType** to **NotifySubjectState**, it also contains `"open"`, `"closed"` and add `"merged"`.
 - In **Issue**, **CreateIssueOption** and **EditIssueOption** structs, `Assignee` got removed. Use `Assignees`.
 - `Type` field at **CreateHookOption** now use **HookType** instead of pure string.

Pulls:
-   [#503 Drop deprecations](https://gitea.com/gitea/go-sdk/pulls/503)
-   [#520 Introduce NotifySubjectState](https://gitea.com/gitea/go-sdk/pulls/520)
