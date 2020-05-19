// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"time"
)

// BranchProtection represents a branch protection for a repository
type BranchProtection struct {
	BranchName                  string   `json:"branch_name"`
	EnablePush                  bool     `json:"enable_push"`
	EnablePushWhitelist         bool     `json:"enable_push_whitelist"`
	PushWhitelistUsernames      []string `json:"push_whitelist_usernames"`
	PushWhitelistTeams          []string `json:"push_whitelist_teams"`
	PushWhitelistDeployKeys     bool     `json:"push_whitelist_deploy_keys"`
	EnableMergeWhitelist        bool     `json:"enable_merge_whitelist"`
	MergeWhitelistUsernames     []string `json:"merge_whitelist_usernames"`
	MergeWhitelistTeams         []string `json:"merge_whitelist_teams"`
	EnableStatusCheck           bool     `json:"enable_status_check"`
	StatusCheckContexts         []string `json:"status_check_contexts"`
	RequiredApprovals           int64    `json:"required_approvals"`
	EnableApprovalsWhitelist    bool     `json:"enable_approvals_whitelist"`
	ApprovalsWhitelistUsernames []string `json:"approvals_whitelist_username"`
	ApprovalsWhitelistTeams     []string `json:"approvals_whitelist_teams"`
	BlockOnRejectedReviews      bool     `json:"block_on_rejected_reviews"`
	BlockOnOutdatedBranch       bool     `json:"block_on_outdated_branch"`
	DismissStaleApprovals       bool     `json:"dismiss_stale_approvals"`
	RequireSignedCommits        bool     `json:"require_signed_commits"`
	ProtectedFilePatterns       string   `json:"protected_file_patterns"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
	// swagger:strfmt date-time
	Updated time.Time `json:"updated_at"`
}

// CreateBranchProtectionOption options for creating a branch protection
type CreateBranchProtectionOption struct {
	BranchName                  string   `json:"branch_name"`
	EnablePush                  bool     `json:"enable_push"`
	EnablePushWhitelist         bool     `json:"enable_push_whitelist"`
	PushWhitelistUsernames      []string `json:"push_whitelist_usernames"`
	PushWhitelistTeams          []string `json:"push_whitelist_teams"`
	PushWhitelistDeployKeys     bool     `json:"push_whitelist_deploy_keys"`
	EnableMergeWhitelist        bool     `json:"enable_merge_whitelist"`
	MergeWhitelistUsernames     []string `json:"merge_whitelist_usernames"`
	MergeWhitelistTeams         []string `json:"merge_whitelist_teams"`
	EnableStatusCheck           bool     `json:"enable_status_check"`
	StatusCheckContexts         []string `json:"status_check_contexts"`
	RequiredApprovals           int64    `json:"required_approvals"`
	EnableApprovalsWhitelist    bool     `json:"enable_approvals_whitelist"`
	ApprovalsWhitelistUsernames []string `json:"approvals_whitelist_username"`
	ApprovalsWhitelistTeams     []string `json:"approvals_whitelist_teams"`
	BlockOnRejectedReviews      bool     `json:"block_on_rejected_reviews"`
	BlockOnOutdatedBranch       bool     `json:"block_on_outdated_branch"`
	DismissStaleApprovals       bool     `json:"dismiss_stale_approvals"`
	RequireSignedCommits        bool     `json:"require_signed_commits"`
	ProtectedFilePatterns       string   `json:"protected_file_patterns"`
}

// EditBranchProtectionOption options for editing a branch protection
type EditBranchProtectionOption struct {
	EnablePush                  *bool    `json:"enable_push"`
	EnablePushWhitelist         *bool    `json:"enable_push_whitelist"`
	PushWhitelistUsernames      []string `json:"push_whitelist_usernames"`
	PushWhitelistTeams          []string `json:"push_whitelist_teams"`
	PushWhitelistDeployKeys     *bool    `json:"push_whitelist_deploy_keys"`
	EnableMergeWhitelist        *bool    `json:"enable_merge_whitelist"`
	MergeWhitelistUsernames     []string `json:"merge_whitelist_usernames"`
	MergeWhitelistTeams         []string `json:"merge_whitelist_teams"`
	EnableStatusCheck           *bool    `json:"enable_status_check"`
	StatusCheckContexts         []string `json:"status_check_contexts"`
	RequiredApprovals           *int64   `json:"required_approvals"`
	EnableApprovalsWhitelist    *bool    `json:"enable_approvals_whitelist"`
	ApprovalsWhitelistUsernames []string `json:"approvals_whitelist_username"`
	ApprovalsWhitelistTeams     []string `json:"approvals_whitelist_teams"`
	BlockOnRejectedReviews      *bool    `json:"block_on_rejected_reviews"`
	BlockOnOutdatedBranch       *bool    `json:"block_on_outdated_branch"`
	DismissStaleApprovals       *bool    `json:"dismiss_stale_approvals"`
	RequireSignedCommits        *bool    `json:"require_signed_commits"`
	ProtectedFilePatterns       *string  `json:"protected_file_patterns"`
}

// ListBranchProtections list branch protections for a repo
func ListBranchProtections() {
	// swagger:operation GET /repos/{owner}/{repo}/branch_protections repository repoListBranchProtection
	// ---
	// summary: List branch protections for a repository
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/BranchProtectionList"
}

// GetBranchProtection gets a branch protection
func GetBranchProtection() {
	// swagger:operation GET /repos/{owner}/{repo}/branch_protections/{name} repository repoGetBranchProtection
	// ---
	// summary: Get a specific branch protection for the repository
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: name
	//   in: path
	//   description: name of protected branch
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/BranchProtection"
	//   "404":
	//     "$ref": "#/responses/notFound"
}

// CreateBranchProtection creates a branch protection for a repo
func CreateBranchProtection(opt CreateBranchProtectionOption) {
	// swagger:operation POST /repos/{owner}/{repo}/branch_protections repository repoCreateBranchProtection
	// ---
	// summary: Create a branch protections for a repository
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/CreateBranchProtectionOption"
	// responses:
	//   "201":
	//     "$ref": "#/responses/BranchProtection"
	//   "403":
	//     "$ref": "#/responses/forbidden"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "422":
	//     "$ref": "#/responses/validationError"
}

// EditBranchProtection edits a branch protection for a repo
func EditBranchProtection(opt EditBranchProtectionOption) {
	// swagger:operation PATCH /repos/{owner}/{repo}/branch_protections/{name} repository repoEditBranchProtection
	// ---
	// summary: Edit a branch protections for a repository. Only fields that are set will be changed
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: name
	//   in: path
	//   description: name of protected branch
	//   type: string
	//   required: true
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/EditBranchProtectionOption"
	// responses:
	//   "200":
	//     "$ref": "#/responses/BranchProtection"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "422":
	//     "$ref": "#/responses/validationError"
}

// DeleteBranchProtection deletes a branch protection for a repo
func DeleteBranchProtection() {
	// swagger:operation DELETE /repos/{owner}/{repo}/branch_protections/{name} repository repoDeleteBranchProtection
	// ---
	// summary: Delete a specific branch protection for the repository
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: name
	//   in: path
	//   description: name of protected branch
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//     "$ref": "#/responses/empty"
	//   "404":
	//     "$ref": "#/responses/notFound"
}
