# Migration Guide: v0.14 to v0.15

v0.15.0 introduces a number of api changes, through which it should not be difficult to migrate.
Just follow this guid and if you still encounter problems, ask for help on discord or feel free to create an issue.

<!-- toc -->

-   [Changed Struct Fields (#520)](#changed-struct-fields)

<!-- tocstop -->

## Changed Struct Fields

 - the State field at **NotificationSubject** changed from **StateType** to **NotifySubjectState**, it also contains "open", "closed" and add "merged"

Pulls:
-   [#520 Introduce NotifySubjectState](https://gitea.com/gitea/go-sdk/pulls/520)

