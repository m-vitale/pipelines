package model

import "fmt"

type PipelineV2 struct {
	UUID           string `gorm:"column:UUID; not null; primary_key"`
	Name           string `gorm:"column:Name; not null"`
	Namespace      string `gorm:"column:Namespace; not null;"`
	Description    string `gorm:"column:Description"`
	PackageId      uint32 `gorm:"column:PackageId; not null"`
	Enabled        bool   `gorm:"column:Enabled; not null"`
	Condition      string `gorm:"column:Condition; not null"`
	MaxConcurrency int64  `gorm:"column:MaxConcurrency"`
	Trigger        `gorm:"column:Trigger;"`
	Parameters     string `gorm:"column:Parameters; size:10000"`   /* Json format argo.v1alpha1.parameter. Set max size to 10,000 */
	CreatedAtInSec int64  `gorm:"column:CreatedAtInSec; not null"` /* The time this record is stored in DB*/
	UpdatedAtInSec int64  `gorm:"column:UpdatedAtInSec; not null"`
}

// Trigger specifies when to create a new workflow.
type Trigger struct {
	// Create workflows according to a cron schedule.
	CronSchedule
	// Create workflows periodically.
	PeriodicSchedule
}

type CronSchedule struct {
	// Time at which scheduling starts.
	// If no start time is specified, the StartTime is the creation time of the schedule.
	CronScheduleStartTimeInSec *int64 `gorm:"column:CronScheduleStartTimeInSec;"`

	// Time at which scheduling ends.
	// If no end time is specified, the EndTime is the end of time.
	CronScheduleEndTimeInSec *int64 `gorm:"column:CronScheduleEndTimeInSec;"`

	// Cron string describing when a workflow should be created within the
	// time interval defined by StartTime and EndTime.
	Cron *string `gorm:"column:Schedule;"`
}

type PeriodicSchedule struct {
	// Time at which scheduling starts.
	// If no start time is specified, the StartTime is the creation time of the schedule.
	PeriodicScheduleStartTimeInSec *int64 `gorm:"column:PeriodicScheduleStartTimeInSec;"`

	// Time at which scheduling ends.
	// If no end time is specified, the EndTime is the end of time.
	PeriodicScheduleEndTimeInSec *int64 `gorm:"column:PeriodicScheduleEndTimeInSec;"`

	// Interval describing when a workflow should be created within the
	// time interval defined by StartTime and EndTime.
	IntervalSecond *int64 `gorm:"column:IntervalSecond;"`
}

type PipelineDetailV2 struct {
	PipelineV2
	/* ScheduledWorkflow CRD. Set size to 65535 so it will be stored as longtext. https://dev.mysql.com/doc/refman/8.0/en/column-count-limit.html */
	ScheduledWorkflow string `gorm:"column:ScheduledWorkflow; size:65535"`
}

func (p PipelineV2) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.UUID)
}

func GetPipelineV2TablePrimaryKeyColumn() string {
	return "UUID"
}
