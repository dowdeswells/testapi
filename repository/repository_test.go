package repository

import (
	. "testing"
	"time"
    "github.com/dowdeswells/testapi/domain"
)

func TestInsertNew(t *T) {
	r := NewRepository()

	u := createUsageSchedule()
	id, err := r.Save(u)

	if (err != nil) {
		t.Errorf("Save returned: %s", err)
	}
	t.Logf("Save returned ID: %s", id)

	u2, err := r.GetByID(id)
	if (err != nil) {
		t.Errorf("GetByID returned: %s", err)
	}

	if (u2.StartDate != u.StartDate) {
		t.Errorf("Expected: %s Got %s", u.StartDate, u2.StartDate)
	}
}
func TestInsertNewThenUpdate(t *T) {
	r := NewRepository()

	u := createUsageSchedule()
	id, err := r.Save(u)

	if (err != nil) {
		t.Errorf("Save returned: %s", err)
	}
	t.Logf("Save returned ID: %s", id)

	u2, err := r.GetByID(id)
	if (err != nil) {
		t.Errorf("GetByID returned: %s", err)
	}

	if (u2.StartDate != u.StartDate) {
		t.Errorf("Expected: %s Got %s", u.StartDate, u2.StartDate)
	}

	nextDate := u.StartDate.AddDate(3,6,2)
	br := u2.AddScheduledAmount(nextDate, 8000)
	if (br.HasBrokenRules()) {
		t.Errorf("Should be valid")
	}

	id, err = r.Save(u2)
	if (err != nil) {
		t.Errorf("2nd Save returned: %s", err)
	}
	t.Logf("Save returned ID: %s", id)
}

func createUsageSchedule() domain.UsageSchedule {

    start := time.Date(2016, 1,1,0,0,0,0,time.Local)
    v := domain.UsageSchedule {
        StartDate: start,
        Scale: domain.Day,
        ScheduledAmounts: []domain.ScheduledAmount {
            domain.ScheduledAmount {
                EndDate: start.AddDate(1,0,0),
                UsageAmount: 2000,
            },
            domain.ScheduledAmount {
                EndDate: start.AddDate(2,0,0),
                UsageAmount: 5000,
            },
        },
    }

    return v
}