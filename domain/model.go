package domain

import (
    "time"
    "github.com/dowdeswells/testapi/brokenrules"
)
//DurationUnits is the type of the duration of schedule in which the amount is applicable
type DurationUnits int

//DurationUnits constants
const (
    Second DurationUnits = 1 + iota
    Minute
    Hour
    Day
    Week
    Month
    Year
)

//ScheduledAmount is the date range in which the Usage is applicable
type ScheduledAmount struct {
    EndDate             time.Time         `json:"endDate"`
    UsageAmount         int64             `json:"usageAmount"`
}

//UsageSchedule is a group of consecutive usage
type UsageSchedule struct {
    ID   string                 `bson:"_id"`
    //id   bson.ObjectId          `bson:"_id"`
    StartDate   time.Time       `json:"startDate"`
    Scale       DurationUnits   `json:"scale"`
    ScheduledAmounts []ScheduledAmount `json:"scheduledAmounts"`
}

//IUsageSchedule is the contract for Usage Schedule behaviour
type IUsageSchedule interface {
    GetScheduledAmounts() []ScheduledAmount
    AddScheduledAmount(endDate time.Time, amount int64) brokenrules.IBrokenRules
}