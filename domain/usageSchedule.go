package domain

import (
    "time"
    //"log"
    "github.com/dowdeswells/testapi/brokenrules"
)

//NewUsageSchedule is a factory
func NewUsageSchedule(startDate time.Time, scale DurationUnits) (u IUsageSchedule) {

    u = &UsageSchedule {
        StartDate: startDate,
        Scale: scale,
    }

    return
}

//AddUsageAmount add another usage to the end on the usage schedule
func (u *UsageSchedule) AddUsageAmount(endDate time.Time, amount int64) brokenrules.IBrokenRules {
    s := ScheduledAmount{
        EndDate:endDate,
        UsageAmount: amount,
    }
    appended := appendUsageAmount(u.ScheduledAmounts, s)
    br := validateEndTimesInOrder(u.StartDate, appended)

    if (!br.HasBrokenRules()) {
        u.ScheduledAmounts = appended
    }

    return br
}

//GetScheduledAmounts is to return data
func (u *UsageSchedule) GetScheduledAmounts() []ScheduledAmount {
    return u.ScheduledAmounts
}

func appendUsageAmount(o []ScheduledAmount, s ScheduledAmount) (appended []ScheduledAmount) {

    if (o == nil) {
        appended = make([]ScheduledAmount, 1)
        appended[0] = s
    } else {
        origLen := len(o)
        appended = make([]ScheduledAmount, origLen + 1)
        copy(appended, o)
        appended[origLen] = s
    }

    return appended
}