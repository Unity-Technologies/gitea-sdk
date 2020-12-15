# Migration Guide: v0.13 to v0.14

v0.14.0 introduces a number of breaking changes, throu it should not be hard to migrate.  
Just follow this guid and if issues still ocure ask for help on discord or  
feel free to create an issue.

<!-- toc -->

-   [Remove Functions for deprecated endpoints (#467)](#Remove-Functions-for-deprecated-endpoints)

<!-- tocstop -->

## Remove Functions for deprecated endpoints

 - for **GetUserTrackedTimes** use **GetRepoTrackedTimes** with user set in options

Pulls:
-   [#467 Remove GetUserTrackedTimes](https://gitea.com/gitea/go-sdk/pulls/467)
