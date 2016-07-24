package domain

import (
    "time"
)

type AddUsageScheduleCmd struct {
    StartDate           time.Time         `json:"startDate"`
    Scale               DurationUnits     `json:"scale"`
}

type AddScheduledAmountCmd struct {
    EndDate             time.Time         `json:"endDate"`
    UsageAmount         int64             `json:"usageAmount"`
}

type RemoveScheduledAmountCmd struct {
    EndDate             time.Time         `json:"endDate"`
}

type IUsageScheduleCommand interface {
    Execute(u UsageSchedule) (UsageSchedule, error)
}

func(c *AddUsageScheduleCmd) Execute(u UsageSchedule) (r UsageSchedule, err error) {
    r = u
    return r, nil
}

func(c *AddScheduledAmountCmd) Execute(u UsageSchedule) (r UsageSchedule, err error) {

    err = u.AddScheduledAmount(c.EndDate, c.UsageAmount)
    r = u
    return r, err
}

func(c *RemoveScheduledAmountCmd) Execute(u UsageSchedule) (r UsageSchedule, err error) {

    r = u
    return r, nil
}
