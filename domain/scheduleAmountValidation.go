package domain

import (
    "time"
    //"log"
    "github.com/dowdeswells/testapi/brokenrules"
)

//Constants for invariant RuleIDs
const (
    RuleEndTimesInOrder = "EndTimeInOrder"
)

func validateEndTimesInOrder(startDate time.Time, amounts []ScheduledAmount) (br brokenrules.IBrokenRules) {
    br = brokenrules.New()
    if (amounts == nil) {
        //log.Print("none")
        return
    }
    last := startDate
    for _, a := range amounts {
        if (a.EndDate.Before(last)) {
            return br.AddRule(RuleEndTimesInOrder)
        }
        last = a.EndDate
    }
    return;
}