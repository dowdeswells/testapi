package router

import (
	"encoding/json"
	"strings"
	. "testing"
	//"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/dowdeswells/testapi/domain"
	"github.com/dowdeswells/testapi/repository"
)

var start = time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)

func TestReadBody(t *T) {

	v := buildUsageSchedule()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf(" %s", err.Error())
	}
	reader := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", "/one", reader)

	v1 := new(domain.UsageSchedule)
	readBody(r, v1)

	if !v1.StartDate.Equal(v.StartDate) {
		t.Errorf("Expected: %s got: %s", v.StartDate, v1.StartDate)
	}

}

func TestPost(t *T) {

	cmd := domain.AddUsageScheduleCmd{
		StartDate: start,
		Scale:     domain.Day,
	}

	b, _ := json.Marshal(cmd)
	reader := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", "/api/usageschedule", reader)

	w := httptest.NewRecorder()

	//s := makeMockStorage()
	router := NewRouter(makeMockStorage)

	router.ServeHTTP(w, r)

	body := w.Body.String()

	if body == "" {
		t.Fatalf("No body returned")
	}

	content := new(IDContent)
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(content)
	if content.ID == "" {
		t.Fatalf("No ID returned: %s", body)
	}
	t.Logf("ID returned: %s", content.ID)
}

func buildUsageSchedule() domain.UsageSchedule {
	start := time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)
	v := domain.UsageSchedule{
		StartDate: start,
		Scale:     domain.Day,
		ScheduledAmounts: []domain.ScheduledAmount{
			domain.ScheduledAmount{
				EndDate:     start.AddDate(1, 0, 0),
				UsageAmount: 2000,
			},
			domain.ScheduledAmount{
				EndDate:     start.AddDate(2, 0, 0),
				UsageAmount: 5000,
			},
		},
	}
	return v
}

type mockStorage struct {
}

func (m *mockStorage) GetByID(id string) (domain.UsageSchedule, error) {
	return buildUsageSchedule(), nil
}

func (m *mockStorage) Save(u domain.UsageSchedule) (id string, err error) {
	err = nil
	id = "1"
	return
}

func makeMockStorage() repository.IRepository {
	return new(mockStorage)
}
