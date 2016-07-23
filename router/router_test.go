package router

import (
	. "testing"
	"encoding/json"
	"strings"
	//"fmt"
	"time"
	"net/http"
	"net/http/httptest"
	"github.com/dowdeswells/testapi/domain"
	"github.com/dowdeswells/testapi/repository"
)

func TestReadBody(t *T) {

	v := buildUsageSchedule()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf(" %s", err.Error())
	}
	reader := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", "/one", reader)

	v1 := new (domain.UsageSchedule)
	readBody(r, v1)

	if !v1.StartDate.Equal(v.StartDate) {
		t.Errorf("Expected: %s got: %s", v.StartDate, v1.StartDate)
	}

}

func TestPostHandler(t *T) {

	v := buildUsageSchedule()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf(" %s", err.Error())
	}

	reader := strings.NewReader(string(b))
	r, _ := http.NewRequest("GET", "/one", reader)

	w := httptest.NewRecorder()

	s := makeMockStorage()
	postHandler(w, r, s)

	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	// We can also get the full body out of the httptest.Recorder, and check
	// that its contents are what we expect
	//fmt.Sprintf("Here's your number: 1\n")
	body := w.Body.String()
	if body != "" {
		t.Fatalf("wrong body returned: %s", body)
	}

}

func buildUsageSchedule() domain.UsageSchedule {
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

type mockStorage struct {

}

func(m *mockStorage) GetByID(id string) (domain.UsageSchedule, error) {
	return buildUsageSchedule(), nil
}

func(m *mockStorage) Save(u domain.UsageSchedule) (id string, err error) {
    err = nil
	id = "1"
    return
}

func makeMockStorage() repository.IRepository {
	return new (mockStorage)
}