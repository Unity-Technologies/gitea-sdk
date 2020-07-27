# Migration Guide: v0.12 to v0.13

v0.13.0 introduces a number of breaking changes, throu it should not be hard to
migrate.
Just follow this guid and if issues still ocure ask for help on discord or
feel free to create an issue.

<!-- toc -->

-   [EditMilestoneOption use StateType (#350)](#EditMilestoneOption-use-StateType)
-   [RepoSearch Options Struct was rewritten (#346)](#RepoSearch-Options-Struct-was-rewritten)
-   [Remove structs witch only have one arguemt (#387)](#Remove-structs-witch-only-have-one-arguemt)

<!-- tocstop -->

## EditMilestoneOption use StateType

Instead of a raw string StateType is now used for State too.
just replace old strings with new enum.


Pulls:

-   [#350 EditMilestoneOption also use StateType](https://gitea.com/gitea/go-sdk/pulls/350)


## RepoSearch Options Struct was rewritten

Since the API itself is ugly and there was no nameconvention whats o ever.
You easely can pass the wrong options and dont get the result you want.

Now it is rewritten and translated for the API.
The easyest way to migrate is to look at who this function is used and rewritten that code block.

If there is a special edgecase you have you can pass a `RawQuery` to the API endpoint.

Pulls:

-   [#346 Refactor RepoSearch to be easy usable](https://gitea.com/gitea/go-sdk/pulls/346)


## Remove structs witch only have one arguemt

Some Functions have Option structs witch will not be extended in the future
and the struct only contain one variable.
So there is no need to warp them up in a extra struct, so they got be removed.
Just pass the variable you set before to the option directly.

Pulls:

-   [#387 Replace IssueLabelsOption with []int64](https://gitea.com/gitea/go-sdk/pulls/387)
