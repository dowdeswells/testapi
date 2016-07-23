package domain

import (
	. "testing"
	"time"
)

func TestAddUsageShouldAppend(t *T) {

	start := time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)
	u := NewUsageSchedule(start, Day)
	d := time.Hour * 24 * 7
	end := start.Add(d)
	u.AddUsageAmount(end, 2000)

	amounts := u.GetScheduledAmounts()
	if len(amounts) != 1 {
		t.Fatalf("Expected %d Got: %d", 1, len(amounts))
	}
	if amounts[0].EndDate != end {
		t.Fatalf("Expected %s Got: %s", end, amounts[0].EndDate)
	}

	end2 := end.Add(d)
	err := u.AddUsageAmount(end2, 4000)

	if err != nil {
		t.Fatalf("Did not expect error %s", err)
	}

	amounts = u.GetScheduledAmounts()

	if len(amounts) != 2 {
		t.Fatalf("Expected %d Got: %d", 2, len(amounts))
	}
	if amounts[0].EndDate != end {
		t.Fatalf("First Amount Expected %s Got: %s", end, amounts[0].EndDate)
	}
	if amounts[1].EndDate != end2 {
		t.Fatalf("Second Amount Expected %s Got: %s", end2, amounts[1].EndDate)
	}
}

func TestAddScheduleAmountShouldValidateEndDate(t *T) {
	start := time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)
	u := NewUsageSchedule(start, Day)
	d := time.Hour * 24 * 7
	end := start.Add(-d)
	br := u.AddUsageAmount(end, 2000)

    hasRules := br.HasBrokenRules()
	if !hasRules {
		t.Fatalf("Expected a validation error")
	}
    n := len(br.GetBrokenRules())
    if (n != 1){
        t.Fatalf("Expected 1 rule, got %d", n)
    }
    n = len(u.GetScheduledAmounts())
    if (n != 0) {
         t.Fatalf("Should not create amounts when invalid: Expected 0 Amounts, got %d", n)
    }
}

func makeUsageSchedule(d time.Duration) IUsageSchedule {
	start := time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)
	u := NewUsageSchedule(start, Day)
	end := start.Add(d)
	u.AddUsageAmount(end, 2000)
	return u
}
