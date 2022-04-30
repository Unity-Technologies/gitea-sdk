// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMilestones(t *testing.T) {
	log.Println("== TestMilestones ==")
	c := newTestClient()

	repo, _ := createTestRepo(t, "TestMilestones", c)
	now := time.Now()
	future := time.Unix(1896134400, 0) // 2030-02-01
	closed := "closed"
	sClosed := StateClosed

	// CreateMilestone 4x
	m1, _, err := c.CreateMilestone(repo.Owner.UserName, repo.Name, CreateMilestoneOption{Title: "v1.0", Description: "First Version", Deadline: &now})
	assert.NoError(t, err)
	_, _, err = c.CreateMilestone(repo.Owner.UserName, repo.Name, CreateMilestoneOption{Title: "v2.0", Description: "Second Version", Deadline: &future})
	assert.NoError(t, err)
	_, _, err = c.CreateMilestone(repo.Owner.UserName, repo.Name, CreateMilestoneOption{Title: "v3.0", Description: "Third Version", Deadline: nil})
	assert.NoError(t, err)
	m4, _, err := c.CreateMilestone(repo.Owner.UserName, repo.Name, CreateMilestoneOption{Title: "temp", Description: "part time milestone"})
	assert.NoError(t, err)

	// EditMilestone
	m1, _, err = c.EditMilestone(repo.Owner.UserName, repo.Name, m1.ID, EditMilestoneOption{Description: &closed, State: &sClosed})
	assert.NoError(t, err)

	// DeleteMilestone
	_, err = c.DeleteMilestone(repo.Owner.UserName, repo.Name, m4.ID)
	assert.NoError(t, err)

	// ListRepoMilestones
	ml, _, err := c.ListRepoMilestones(repo.Owner.UserName, repo.Name, ListMilestoneOption{})
	assert.NoError(t, err)
	assert.Len(t, ml, 2)
	ml, _, err = c.ListRepoMilestones(repo.Owner.UserName, repo.Name, ListMilestoneOption{State: StateClosed})
	assert.NoError(t, err)
	assert.Len(t, ml, 1)
	ml, _, err = c.ListRepoMilestones(repo.Owner.UserName, repo.Name, ListMilestoneOption{State: StateAll})
	assert.NoError(t, err)
	assert.Len(t, ml, 3)
	ml, _, err = c.ListRepoMilestones(repo.Owner.UserName, repo.Name, ListMilestoneOption{State: StateAll, Name: "V3.0"})
	assert.NoError(t, err)
	assert.Len(t, ml, 1)
	assert.EqualValues(t, "v3.0", ml[0].Title)

	// test fallback resolveMilestoneByName
	m, _, err := c.resolveMilestoneByName(repo.Owner.UserName, repo.Name, "V3.0")
	assert.NoError(t, err)
	assert.EqualValues(t, ml[0].ID, m.ID)
	_, _, err = c.resolveMilestoneByName(repo.Owner.UserName, repo.Name, "NoEvidenceOfExist")
	assert.Error(t, err)
	assert.EqualValues(t, "milestone 'NoEvidenceOfExist' do not exist", err.Error())

	// GetMilestone
	_, _, err = c.GetMilestone(repo.Owner.UserName, repo.Name, m4.ID)
	assert.Error(t, err)
	m, _, err = c.GetMilestone(repo.Owner.UserName, repo.Name, m1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, m1, m)
	m2, _, err := c.GetMilestoneByName(repo.Owner.UserName, repo.Name, m.Title)
	assert.NoError(t, err)
	assert.EqualValues(t, m, m2)
}
