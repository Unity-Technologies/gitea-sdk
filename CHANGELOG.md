# Changelog

## [v0.15.1](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.15.1) - 2022-01-04

* FEATURES
  * Add ignoreVersion & manuall version set option (#560) (#562)
* BUGFIXES
  * Fix version string for next release (#559)


## [v0.15.0](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.15.0) - 2021-08-13

* BREAKING
  * Introduce NotifySubjectState (#520)
  * Drop deprecations (#503)
* FEATURES
  * Add Repo Team Management Functions (#537)
  * Add CreateRepoFromTemplate (#536)
  * Add GetReviewers & GetAssignees (#534)
  * Add GetTag, GetAnnotatedTag & CreateTag (#533)
  * Add GetUserSettings & UpdateUserSettings (#531)
  * Add ListPullRequestCommits (#530)
  * Add GetUserByID (#513)
  * Add GetRepoByID (#511)
* ENHANCEMENTS
  * Update List Options (#527)
  * Update Structs (#524)
  * ListFunctions: option to disable pagination (#509)


## [v0.14.1](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.14.1) - 2021-06-30

* BUGFIXES
  * Fix setDefaults (#508) (#510)


## [v0.14.0](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.14.0) - 2021-03-21

* BREAKING
  * Update Structs (#486)
  * Added repo ListContents and changed GetContents doc to talk about a single file (#485)
  * Remove & Rename  TrackedTimes list functions (#467)
  * UrlEscape Function Arguments used in UrlPath (#273)
* FEATURES
  * Add Create/Delete ReviewRequests (#493)
  * Add Un-/DismissPullReview funcs (#489)
  * Add Repo Un-Star Functions (#483)
  * introduce Client.GetArchiveReader (#476)
  * Add DeleteRepoTag function (#461)
  * Add GetReleaseByTag (#427)
* BUGFIXES
  * Handle Contents Edge-Case (#492)
  * Fix GetCombinedStatus() (#470)
  * Use Predefind Versions & Compare Function (#442)
  * Return resp on NotFound too (#428)
* ENHANCEMENTS
  * Add workaround to get head branch sha of pulls with deleted head branch (#498)
  * GetFile: Use "ref" in-query if posible (#491)
  * Add DeleteTag & Correct DeleteReleaseByTag (#488)
  * Add html_url field to Release struct (#477)
  * Add Ref to Issue structs (#466)
  * Update Issue Struct (#458)
  * Use sync.Once for loading ServerVersion (#456)
  * Add Gitea2Gitea Migration Support (#454)
  * Add Helper for Optional Values (#448)
  * Update CreateRepoOption struct (#445)
  * Update CommitMeta Struct (#434)
  * Update Struct NotificationSubject (#424)
  * Add Debug Mode (#422)
* DOCS
  * Make Client thread-safe & add docs (#495)
  * Improve PullReview docs (#469)


## [v0.13.3](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.13.3) - 2021-03-22

* BUGFIXES
  * Fix GetCombinedStatus() (#470) (#472)
* ENHANCEMENTS
  * Add html_url field to Release struct (#477) (#478)


## [v0.13.2](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.13.2) - 2020-12-07

* BUGFIXES
  * Use Predefind Versions & Compare Function (#442) (#446)
* ENHANCEMENTS
  * Add Gitea2Gitea Migration Support (#454) (#455)
  * Update CreateRepoOption struct (#445) (#447)
  * Update CommitMeta Struct (#434) (#437)


## [v0.13.1](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.13.1) - 2020-09-29

* FEATURES
  * Add GetReleaseByTag (#427) (#430)
* BUGFIXES
  * Return http Response on NotFound too (#428) (#429)
* ENHANCEMENTS
  * Update Struct NotificationSubject (#424) (#425)
  * Add Debug Mode (#422) (#423)


## [v0.13.0](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.13.0) - 2020-09-15

* BREAKING
  * Check Gitea Version Requirement (#419)
  * All Function return http responce (#416)
  * Remove opts from ListPullReviewComments (#411)
  * Use enum AccessMode for OrgTeam and Collaborator functions (#408)
  * CreateOrgOption rename UserName to Name (#386)
  * EditMilestoneOption also use StateType (#350)
  * Refactor RepoSearch to be easy usable (#346)
* FEATURES
  * Milestone Functions accept name to identify (#418)
  * Make http requests with context (#417)
  * Add GetGlobalAttachmentSettings (#414)
  * Add GetArchive (#413)
  * Add GetRepoLanguages + TESTs (#412)
  * Add CreateBranch (#407)
  * Add Admin CronTask functions (#406)
  * Add GetGlobalAPISettings Function (#404)
  * Add Get Diff and Patch endpoints for pull requests (#398)
  * Add Validate func for Create/Edit Options (#370)
  * Add Function to get GetGlobalSettings and GetSettingAllowedReactions (#359)
* ENHANCEMENTS
  * TrackedTime API >= 1.11.x needed (#415)
  * Update Milestone struct (#410)
  * Add Fallback for GetPullRequestDiff/Patch (#399)
  * DeleteToken Accept Names too (#394)
  * Update ListMilestoneOption struct (#393)
  * Migration Api Changed (#392)
  * Refactor Visibletype Orgs (#382)
  * Extend Notification Functions (#381)
  * Update GetGlobalSettings Functions (#376)
  * Allow Creating Closed Milestones (#373)
  * CreateLabel correct Color if needed for old versions (#365)
  * Issue/Pull add IsLocked Property (#357)
  * Update EditPullRequestOption Add Base (#353)
  * File Create/Update/Delete detect DefaultBranch if Branch not set for old Versions (#352)
  * Improve Error Handling (#351)

## [v0.12.2](https://gitea.com/gitea/go-sdk/releases/tag/gitea/v0.12.2) - 2020-09-05

* ENHANCEMENTS
  * Extend Notification Functions (#381) (#385)

## [v0.12.1](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1268) - 2020-07-09

* ENHANCEMENTS
  * Improve Error Handling (#351) (#377)
  * Allow Creating Closed Milestones (#373) (#375)
  * File Create/Update/Delete detect DefaultBranch if Branch not set for old Versions (#352) (#372)
  * CreateLabel correct Color if needed for old versions (#365) (#371)
  * Update EditPullRequestOption Add Base (#353) (#363)

## [v0.12.0](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1223) - 2020-05-21

* BREAKING
  * Support 2FA for basic auth & refactor Token functions (#335)
  * PullMerge: use enum for MergeStyle (#328)
  * Refactor List/SetRepoTopics (#276)
  * Remove ListUserIssues() ... (#262)
  * Extend SearchUsers (#248)
  * Fix & Refactor UserApp Functions (#247)
  * Add ListMilestoneOption to ListRepoMilestones (#244)
  * Add ListIssueCommentOptions for optional param (#243)
  * Refactor RepoWatch (#241)
  * Add Pagination Options for List Requests (#205)
* FEATURES
  * Add BranchProtection functions (#341)
  * Add PullReview functions (#338)
  * Add Issue Subscription Check & Fix DeleteIssueSubscription (#318)
  * Add Branch Deletion (#317)
  * Add Get/Update for oauth2 apps (#311)
  * Add Create/Get/Delete for oauth2 apps (#305)
  * Add DeleteFile() (#302)
  * Add Get/Update/Create File (#281)
  * Add List/Check/SetPublic/Delete OrgMembership functions (#275)
  * Add ListRepoCommits (#266)
  * Add TransferRepo (#264)
  * Add SearchRepo API Call (#254)
  * Add ListOptions struct (#249)
  * Add Notification functions (#226)
  * Add GetIssueComment (#216)
* BUGFIXES
  * Add missing JSON header to AddCollaborator() (#306)
  * On Internal Server Error, show request witch caused this (#296)
  * Fix MergePullRequest & extend Tests (#278)
  * Fix AddEmail (#260)
* ENHANCEMENTS
  * Check if gitea is able to squash-merge via API (#336)
  * ListIssues: add milestones filter (#327)
  * Update CreateRepoOption struct (#300)
  * Add IssueType as filter for ListIssues (#286)
  * Extend ListDeployKeys (#268)
  * Use RepositoryMeta struct on Issues (#267)
  * Use StateType (#265)
  * Extend Issue Struct (#258)
  * IssueSubscribtion: Check http Status responce (#242)

## [v0.11.3](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1259) - 2020-04-27
* BUGFIXES
  * Fix MergePullRequest (#278) (#316)
  * Add missing JSON header to AddCollaborator() (#307)

## [v0.11.2](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1256) - 2020-03-31
* ENHANCEMENTS
  * On Internal Server Error, show request witch caused this (#297)

## [v0.11.1](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1235) - 2020-03-29
* BUGFIXES
  * Fix SetRepoTopics (#276) (#274)
  * Fix AddEmail (#260) (#261)
  * Fix UserApp Functions (#247) (#256)
* ENHANCEMENTS
  * Add IssueType as filter for ListIssues (#288)
  * Correct version (#259)

## [v0.11.0](https://gitea.com/gitea/go-sdk/pulls?q=&type=all&state=closed&milestone=1222) - 2020-01-27
* FEATURES
  * Add VersionCheck (#215)
  * Add Issue Un-/Subscription function (#214)
  * Add Reaction struct and functions (#213)
  * Add GetBlob (#212)
* BUGFIXES
  * Fix ListIssue Functions (#225)
  * Fix ListRepoPullRequests (#219)
* ENHANCEMENTS
  * Add some pull list options (#217)
  * Extend StopWatch struct & functions (#211)
* TESTING
  * Add Test Framework (#227)
* BUILD
  * Use golangci-lint and revive for linting (#220)
