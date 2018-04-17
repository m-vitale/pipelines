// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"errors"
	"ml/src/message"
	"ml/src/util"
	"net/http"
	"testing"

	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultScheduledAtInSec = 10
	defaultCreatedAtInSec   = 20
)

func initializeJobDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("sqlite3", db)
	return gormDB, mock
}

func createWorkflow(name string) *v1alpha1.Workflow {
	return &v1alpha1.Workflow{
		ObjectMeta: v1.ObjectMeta{Name: name},
		Status:     v1alpha1.WorkflowStatus{Phase: "Pending"}}
}

func TestCreateJob(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()

	wf1 := createWorkflow("wf1")

	jobExpected := message.Job{
		CreatedAtInSec:   defaultCreatedAtInSec,
		Name:             wf1.Name,
		ScheduledAtInSec: defaultScheduledAtInSec,
		Status:           message.JobExecutionPending,
		UpdatedAtInSec:   1,
		PipelineID:       1,
	}
	wfExpected := createWorkflow(wf1.Name)
	jobDetailExpect := message.JobDetail{
		Workflow: wfExpected,
		Job:      &jobExpected}
	jobDetail, err := store.JobStore().CreateJob(1, wf1, defaultScheduledAtInSec, defaultCreatedAtInSec)

	assert.Nil(t, err)
	assert.Equal(t, jobDetailExpect, *jobDetail, "Unexpected Job parsed.")

	job, err := getJobMetadata(store.DB(), 1, wf1.Name)
	assert.Equal(t, jobExpected, *job)
}

func TestCreateJob_CreateWorkflowFailed(t *testing.T) {
	fakeClients := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer fakeClients.Close()
	store := &JobStore{db: fakeClients.DB(), wfClient: &FakeBadWorkflowClient{}, time: fakeClients.time}

	wf1 := createWorkflow("wf1")
	jobExpected := message.Job{
		CreatedAtInSec:   defaultCreatedAtInSec,
		UpdatedAtInSec:   defaultCreatedAtInSec,
		Name:             wf1.Name,
		Status:           message.JobCreationPending,
		ScheduledAtInSec: defaultScheduledAtInSec,
		PipelineID:       1,
	}
	wfExpected := createWorkflow(wf1.Name)
	jobDetailExpect := message.JobDetail{
		Workflow: wfExpected,
		Job:      &jobExpected}

	jobDetail, err := store.CreateJob(1, wf1, defaultScheduledAtInSec, defaultCreatedAtInSec)
	assert.Nil(t, err)
	assert.Equal(t, jobDetailExpect, *jobDetail, "Unexpected Job parsed.")

	job, err := getJobMetadata(fakeClients.DB(), 1, wf1.Name)
	assert.Equal(t, jobExpected, *job)
}

func TestCreateJob_CreateMetadataError(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	store.DB().Close()

	_, err := store.JobStore().CreateJob(1, &v1alpha1.Workflow{},
		defaultScheduledAtInSec, defaultCreatedAtInSec)
	assert.Equal(t, http.StatusInternalServerError, err.(*util.UserError).ExternalStatusCode(),
		"Expected to throw an internal error")
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "Internal Server Error")
	assert.Contains(t, err.(*util.UserError).Error(), "Failed to store job metadata")
}

func TestCreateJob_UpdateMetadataFailed(t *testing.T) {
	db, mock := initializeJobDB()
	mock.ExpectExec("INSERT INTO \"jobs\"").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE jobs").WillReturnError(errors.New("something"))

	store := NewJobStore(db, NewWorkflowClientFake(), util.NewFakeTimeForEpoch())
	wf1 := createWorkflow("wf1")
	jobExpected := message.Job{
		CreatedAtInSec:   defaultCreatedAtInSec,
		UpdatedAtInSec:   defaultCreatedAtInSec,
		Name:             wf1.Name,
		Status:           message.JobExecutionPending,
		ScheduledAtInSec: defaultScheduledAtInSec,
		PipelineID:       1,
	}
	wfExpected := createWorkflow(wf1.Name)
	jobDetailExpect := message.JobDetail{
		Workflow: wfExpected,
		Job:      &jobExpected}

	jobDetail, err := store.CreateJob(1, wf1, defaultScheduledAtInSec, defaultCreatedAtInSec)

	assert.Nil(t, err)
	assert.Equal(t, jobDetailExpect, *jobDetail)
}

func TestListJobs(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	jobDetail, err := store.JobStore().CreateJob(1, createWorkflow("wf1"),
		defaultScheduledAtInSec, defaultCreatedAtInSec)
	store.JobStore().CreateJob(2, createWorkflow("wf2"),
		defaultScheduledAtInSec, defaultCreatedAtInSec)

	jobsExpected := []message.Job{
		{
			CreatedAtInSec:   defaultCreatedAtInSec,
			UpdatedAtInSec:   1,
			Name:             jobDetail.Job.Name,
			Status:           message.JobExecutionPending,
			ScheduledAtInSec: defaultScheduledAtInSec,
			PipelineID:       1,
		}}
	jobs, err := store.JobStore().ListJobs(1)
	assert.Nil(t, err)
	assert.Equal(t, jobsExpected, jobs, "Unexpected Job listed.")

	jobs, err = store.JobStore().ListJobs(3)
	assert.Empty(t, jobs)
}

func TestListJobsError(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	store.DB().Close()
	_, err := store.JobStore().ListJobs(1)
	assert.Equal(t, http.StatusInternalServerError, err.(*util.UserError).ExternalStatusCode(),
		"Expected to throw an internal error")
}

func TestGetJob(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	wf1 := createWorkflow("wf1")
	createdJobDetail, err := store.JobStore().CreateJob(1, wf1,
		defaultScheduledAtInSec, defaultCreatedAtInSec)
	assert.Nil(t, err)
	jobDetailExpect := message.JobDetail{
		Workflow: wf1,
		Job: &message.Job{
			CreatedAtInSec:   defaultCreatedAtInSec,
			UpdatedAtInSec:   1,
			Status:           message.JobExecutionPending,
			Name:             createdJobDetail.Job.Name,
			ScheduledAtInSec: defaultScheduledAtInSec,
			PipelineID:       1,
		}}

	jobDetail, err := store.JobStore().GetJob(1, wf1.Name)
	assert.Nil(t, err)
	assert.Equal(t, jobDetailExpect, *jobDetail)
}

func TestGetJob_NotFoundError(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()

	_, err := store.JobStore().GetJob(1, "wf1")
	assert.Equal(t, http.StatusNotFound, err.(*util.UserError).ExternalStatusCode(),
		"Expected not to find the job")
}

func TestGetJob_InternalError(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	wf1 := createWorkflow("wf1")
	store.JobStore().CreateJob(1, wf1,
		defaultScheduledAtInSec, defaultCreatedAtInSec)
	store.DB().Close()

	_, err := store.JobStore().GetJob(1, wf1.Name)
	assert.Equal(t, http.StatusInternalServerError, err.(*util.UserError).ExternalStatusCode(),
		"Expected get job to return internal error")
}

func TestGetJob_GetWorkflowError(t *testing.T) {
	store := NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	defer store.Close()
	wf1 := createWorkflow("wf1")
	store.JobStore().CreateJob(1, wf1,
		defaultScheduledAtInSec, defaultCreatedAtInSec)

	jobStore := NewJobStore(store.DB(), &FakeBadWorkflowClient{}, util.NewFakeTimeForEpoch())
	_, err := jobStore.GetJob(1, wf1.Name)
	assert.Equal(t, http.StatusInternalServerError, err.(*util.UserError).ExternalStatusCode(),
		"Expected to throw an internal error")
}
