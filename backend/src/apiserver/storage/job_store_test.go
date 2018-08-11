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
	"testing"
	"time"

	"database/sql"

	"github.com/googleprivate/ml/backend/src/apiserver/model"
	"github.com/googleprivate/ml/backend/src/common/util"
	swfapi "github.com/googleprivate/ml/backend/src/crd/pkg/apis/scheduledworkflow/v1alpha1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

func initializeDbAndStore() (*sql.DB, *JobStore) {
	db := NewFakeDbOrFatal()
	jobStore := NewJobStore(db, util.NewFakeTimeForEpoch())
	job1 := &model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Enabled:     true,
		Conditions:  "ready",
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}
	jobStore.CreateJob(job1)
	job2 := &model.Job{
		UUID:        "2",
		DisplayName: "pp 2",
		Name:        "pp2",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Trigger: model.Trigger{
			CronSchedule: model.CronSchedule{
				CronScheduleStartTimeInSec: util.Int64Pointer(1),
				CronScheduleEndTimeInSec:   util.Int64Pointer(2),
				Cron:                       util.StringPointer("1 * *"),
			},
		},
		Enabled:        true,
		CreatedAtInSec: 2,
		UpdatedAtInSec: 2,
	}
	jobStore.CreateJob(job2)

	return db, jobStore
}

func TestListJobs_Pagination(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobsExpected := []model.Job{
		{
			UUID:        "1",
			DisplayName: "pp 1",
			Name:        "pp1",
			Namespace:   "n1",
			PipelineId:  "1",
			Conditions:  "ready",
			Enabled:     true,
			Trigger: model.Trigger{
				PeriodicSchedule: model.PeriodicSchedule{
					PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
					PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
					IntervalSecond:                 util.Int64Pointer(3),
				},
			},
			CreatedAtInSec: 1,
			UpdatedAtInSec: 1,
		}}
	jobs, nextPageToken, err := jobStore.ListJobs("" /*pageToken*/, 1 /*pageSize*/, "Name" /*sortByFieldName*/, false /*isDesc*/)
	assert.Nil(t, err)
	assert.NotEmpty(t, nextPageToken)
	assert.Equal(t, jobsExpected, jobs)
	jobsExpected2 := []model.Job{
		{
			UUID:        "2",
			DisplayName: "pp 2",
			Name:        "pp2",
			Namespace:   "n1",
			PipelineId:  "1",
			Enabled:     true,
			Trigger: model.Trigger{
				CronSchedule: model.CronSchedule{
					CronScheduleStartTimeInSec: util.Int64Pointer(1),
					CronScheduleEndTimeInSec:   util.Int64Pointer(2),
					Cron:                       util.StringPointer("1 * *"),
				},
			},
			CreatedAtInSec: 2,
			UpdatedAtInSec: 2,
			Conditions:     "ready",
		}}
	jobs, newToken, err := jobStore.ListJobs(nextPageToken, 2 /*pageSize*/, "Name" /*sortByFieldName*/, false /*isDesc*/)
	assert.Nil(t, err)
	assert.Equal(t, "", newToken)
	assert.Equal(t, jobsExpected2, jobs)
}

func TestListJobs_Pagination_Descent(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobsExpected := []model.Job{
		{
			UUID:        "2",
			DisplayName: "pp 2",
			Name:        "pp2",
			Namespace:   "n1",
			PipelineId:  "1",
			Enabled:     true,
			Conditions:  "ready",
			Trigger: model.Trigger{
				CronSchedule: model.CronSchedule{
					CronScheduleStartTimeInSec: util.Int64Pointer(1),
					CronScheduleEndTimeInSec:   util.Int64Pointer(2),
					Cron:                       util.StringPointer("1 * *"),
				},
			},
			CreatedAtInSec: 2,
			UpdatedAtInSec: 2,
		}}
	jobs, nextPageToken, err := jobStore.ListJobs("" /*pageToken*/, 1 /*pageSize*/, "Name" /*sortByFieldName*/, true /*isDesc*/)
	assert.Nil(t, err)
	assert.NotEmpty(t, nextPageToken)
	assert.Equal(t, jobsExpected, jobs)
	jobsExpected2 := []model.Job{
		{
			UUID:        "1",
			DisplayName: "pp 1",
			Name:        "pp1",
			Namespace:   "n1",
			PipelineId:  "1",
			Enabled:     true,
			Conditions:  "ready",
			Trigger: model.Trigger{
				PeriodicSchedule: model.PeriodicSchedule{
					PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
					PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
					IntervalSecond:                 util.Int64Pointer(3),
				},
			},
			CreatedAtInSec: 1,
			UpdatedAtInSec: 1,
		}}
	jobs, newToken, err := jobStore.ListJobs(nextPageToken, 2 /*pageSize*/, "Name" /*sortByFieldName*/, true /*isDesc*/)
	assert.Nil(t, err)
	assert.Equal(t, "", newToken)
	assert.Equal(t, jobsExpected2, jobs)
}

func TestListJobs_Pagination_LessThanPageSize(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobsExpected := []model.Job{
		{
			UUID:        "1",
			DisplayName: "pp 1",
			Name:        "pp1",
			Namespace:   "n1",
			PipelineId:  "1",
			Enabled:     true,
			Conditions:  "ready",
			Trigger: model.Trigger{
				PeriodicSchedule: model.PeriodicSchedule{
					PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
					PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
					IntervalSecond:                 util.Int64Pointer(3),
				},
			},
			CreatedAtInSec: 1,
			UpdatedAtInSec: 1,
		},
		{
			UUID:        "2",
			DisplayName: "pp 2",
			Name:        "pp2",
			Namespace:   "n1",
			PipelineId:  "1",
			Enabled:     true,
			Conditions:  "ready",
			Trigger: model.Trigger{
				CronSchedule: model.CronSchedule{
					CronScheduleStartTimeInSec: util.Int64Pointer(1),
					CronScheduleEndTimeInSec:   util.Int64Pointer(2),
					Cron:                       util.StringPointer("1 * *"),
				},
			},
			CreatedAtInSec: 2,
			UpdatedAtInSec: 2,
		}}
	jobs, nextPageToken, err := jobStore.ListJobs("" /*pageToken*/, 2 /*pageSize*/, model.GetJobTablePrimaryKeyColumn() /*sortByFieldName*/, false /*isDesc*/)
	assert.Nil(t, err)
	assert.Equal(t, "", nextPageToken)
	assert.Equal(t, jobsExpected, jobs)
}

func TestListJobsError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	db.Close()
	_, _, err := jobStore.ListJobs("" /*pageToken*/, 2 /*pageSize*/, "Name" /*sortByFieldName*/, false /*isDesc*/)

	assert.Equal(t, codes.Internal, err.(*util.UserError).ExternalStatusCode(),
		"Expected to list job to return error")
}

func TestGetJob(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobExpected := model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		Enabled:        true,
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}

	job, err := jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job, "Got unexpected job")
}

func TestGetJob_NotFoundError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	_, err := jobStore.GetJob("notexist")
	assert.Equal(t, codes.NotFound, err.(*util.UserError).ExternalStatusCode(),
		"Expected get job to return not found error")
}

func TestGetJob_InternalError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	db.Close()
	_, err := jobStore.GetJob("1")
	assert.Equal(t, codes.Internal, err.(*util.UserError).ExternalStatusCode(),
		"Expected get job to return internal error")
}

func TestDeleteJob(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	err := jobStore.DeleteJob("1")
	assert.Nil(t, err)
	_, err = jobStore.GetJob("1")
	assert.Equal(t, codes.NotFound, err.(*util.UserError).ExternalStatusCode())
}

func TestDeleteJob_InternalError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	db.Close()

	err := jobStore.DeleteJob("1")
	assert.Equal(t, codes.Internal, err.(*util.UserError).ExternalStatusCode(),
		"Expected delete job to return internal error")
}

func TestCreateJob(t *testing.T) {
	db := NewFakeDbOrFatal()
	defer db.Close()
	jobStore := NewJobStore(db, util.NewFakeTimeForEpoch())
	job := &model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Enabled:     true,
	}

	job, err := jobStore.CreateJob(job)
	assert.Nil(t, err)
	jobExpected := &model.Job{
		UUID:           "1",
		DisplayName:    "pp 1",
		Name:           "pp1",
		Namespace:      "n1",
		PipelineId:     "1",
		Enabled:        true,
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}
	assert.Equal(t, jobExpected, job, "Got unexpected jobs")
}

func TestCreateJobError(t *testing.T) {
	db := NewFakeDbOrFatal()
	defer db.Close()
	jobStore := NewJobStore(db, util.NewFakeTimeForEpoch())
	db.Close()
	job := &model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Enabled:     true,
	}

	job, err := jobStore.CreateJob(job)
	assert.Equal(t, codes.Internal, err.(*util.UserError).ExternalStatusCode(),
		"Expected create job to return error")
}

func TestEnableJob(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	err := jobStore.EnableJob("1", false)
	assert.Nil(t, err)

	jobExpected := model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Enabled:     false,
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		CreatedAtInSec: 1,
		UpdatedAtInSec: 3,
	}

	job, err := jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job, "Got unexpected job")
}

func TestEnableJob_SkipUpdate(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	err := jobStore.EnableJob("1", true)
	assert.Nil(t, err)

	jobExpected := model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Enabled:     true,
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}

	job, err := jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job, "Got unexpected job")
}

func TestEnableJob_DatabaseError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	db.Close()

	// Enabling the job.
	err := jobStore.EnableJob("1", true)
	println(err.Error())
	assert.Contains(t, err.Error(), "Error when enabling job 1 to true: sql: database is closed")
}

func TestUpdateJob_Success(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobExpected := model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Enabled:     true,
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}

	job, err := jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job)

	swf := util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "MY_NAME",
			Namespace: "MY_NAMESPACE",
			UID:       "1",
		},
		Spec: swfapi.ScheduledWorkflowSpec{
			Enabled:        false,
			MaxConcurrency: util.Int64Pointer(200),
			Workflow: &swfapi.WorkflowResource{
				Parameters: []swfapi.Parameter{
					{Name: "PARAM1", Value: "NEW_VALUE1"},
				},
			},
			Trigger: swfapi.Trigger{
				CronSchedule: &swfapi.CronSchedule{
					StartTime: util.MetaV1TimePointer(metav1.NewTime(time.Unix(10, 0).UTC())),
					EndTime:   util.MetaV1TimePointer(metav1.NewTime(time.Unix(20, 0).UTC())),
					Cron:      "MY_CRON",
				},
				PeriodicSchedule: &swfapi.PeriodicSchedule{
					StartTime:      util.MetaV1TimePointer(metav1.NewTime(time.Unix(30, 0).UTC())),
					EndTime:        util.MetaV1TimePointer(metav1.NewTime(time.Unix(40, 0).UTC())),
					IntervalSecond: 50,
				},
			},
		},
		Status: swfapi.ScheduledWorkflowStatus{
			Conditions: []swfapi.ScheduledWorkflowCondition{{
				Type:               swfapi.ScheduledWorkflowEnabled,
				Status:             core.ConditionTrue,
				LastProbeTime:      metav1.NewTime(time.Unix(10, 0).UTC()),
				LastTransitionTime: metav1.NewTime(time.Unix(20, 0).UTC()),
				Reason:             string(swfapi.ScheduledWorkflowEnabled),
				Message:            "The schedule is enabled.",
			},
			},
		},
	})

	err = jobStore.UpdateJob(swf)
	assert.Nil(t, err)

	jobExpected = model.Job{
		UUID:           "1",
		DisplayName:    "pp 1",
		Name:           "MY_NAME",
		Namespace:      "MY_NAMESPACE",
		PipelineId:     "1",
		Enabled:        false,
		Conditions:     "Enabled:",
		CreatedAtInSec: 1,
		UpdatedAtInSec: 3,
		MaxConcurrency: 200,
		Parameters:     "[{\"name\":\"PARAM1\",\"value\":\"NEW_VALUE1\"}]",
		Trigger: model.Trigger{
			CronSchedule: model.CronSchedule{
				CronScheduleStartTimeInSec: util.Int64Pointer(10),
				CronScheduleEndTimeInSec:   util.Int64Pointer(20),
				Cron:                       util.StringPointer("MY_CRON"),
			},
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(30),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(40),
				IntervalSecond:                 util.Int64Pointer(50),
			},
		},
	}

	job, err = jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job)
}

func TestUpdateJob_MostlyEmptySpec(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	jobExpected := model.Job{
		UUID:        "1",
		DisplayName: "pp 1",
		Name:        "pp1",
		Namespace:   "n1",
		PipelineId:  "1",
		Conditions:  "ready",
		Enabled:     true,
		Trigger: model.Trigger{
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: util.Int64Pointer(1),
				PeriodicScheduleEndTimeInSec:   util.Int64Pointer(2),
				IntervalSecond:                 util.Int64Pointer(3),
			},
		},
		CreatedAtInSec: 1,
		UpdatedAtInSec: 1,
	}

	job, err := jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job)

	swf := util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "MY_NAME",
			Namespace: "MY_NAMESPACE",
			UID:       "1",
		},
	})

	err = jobStore.UpdateJob(swf)
	assert.Nil(t, err)

	jobExpected = model.Job{
		UUID:           "1",
		DisplayName:    "pp 1",
		Name:           "MY_NAME",
		Namespace:      "MY_NAMESPACE",
		PipelineId:     "1",
		Enabled:        false,
		Conditions:     "NO_STATUS:",
		CreatedAtInSec: 1,
		UpdatedAtInSec: 3,
		Parameters:     "[]",
		Trigger: model.Trigger{
			CronSchedule: model.CronSchedule{
				CronScheduleStartTimeInSec: nil,
				CronScheduleEndTimeInSec:   nil,
				Cron:                       util.StringPointer(""),
			},
			PeriodicSchedule: model.PeriodicSchedule{
				PeriodicScheduleStartTimeInSec: nil,
				PeriodicScheduleEndTimeInSec:   nil,
				IntervalSecond:                 util.Int64Pointer(0),
			},
		},
	}

	job, err = jobStore.GetJob("1")
	assert.Nil(t, err)
	assert.Equal(t, jobExpected, *job)
}

func TestUpdateJob_MissingField(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	// Name
	swf := util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "MY_NAMESPACE",
			UID:       "1",
		},
	})

	err := jobStore.UpdateJob(swf)
	assert.NotNil(t, err)
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "The resource must have a name")
	assert.Equal(t, err.(*util.UserError).ExternalStatusCode(), codes.InvalidArgument)

	// Namespace
	swf = util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name: "MY_NAME",
			UID:  "1",
		},
	})

	err = jobStore.UpdateJob(swf)
	assert.NotNil(t, err)
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "The resource must have a namespace")
	assert.Equal(t, err.(*util.UserError).ExternalStatusCode(), codes.InvalidArgument)

	// UID
	swf = util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "MY_NAME",
			Namespace: "MY_NAMESPACE",
		},
	})

	err = jobStore.UpdateJob(swf)
	assert.NotNil(t, err)
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "The resource must have a UID")
	assert.Equal(t, err.(*util.UserError).ExternalStatusCode(), codes.InvalidArgument)
}

func TestUpdateJob_RecordNotFound(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	defer db.Close()

	swf := util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "MY_NAME",
			Namespace: "MY_NAMESPACE",
			UID:       "UNKNOWN_UID",
		},
	})

	err := jobStore.UpdateJob(swf)
	assert.NotNil(t, err)
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "There is no job")
	assert.Equal(t, err.(*util.UserError).ExternalStatusCode(), codes.InvalidArgument)
}

func TestUpdateJob_InternalError(t *testing.T) {
	db, jobStore := initializeDbAndStore()
	db.Close()
	swf := util.NewScheduledWorkflow(&swfapi.ScheduledWorkflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "MY_NAME",
			Namespace: "MY_NAMESPACE",
			UID:       "UNKNOWN_UID",
		},
	})

	err := jobStore.UpdateJob(swf)
	assert.NotNil(t, err)
	assert.Contains(t, err.(*util.UserError).ExternalMessage(), "Internal Server Error")
	assert.Contains(t, err.(*util.UserError).Error(), "database is closed")
	assert.Equal(t, err.(*util.UserError).ExternalStatusCode(), codes.Internal)
}
