package router

import (
	"log"
	//"errors"
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/dowdeswells/testapi/repository"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type requestResponseFunc func(http.ResponseWriter, *http.Request)
type requestResponseWithErrorFunc func(http.ResponseWriter, *http.Request) error

type commandHandler func(*http.Request, repository.IRepository) (interface{}, error)

type commandRoute struct {
	Route   string
	Method  string
	Handler commandHandler
}

var routes = []commandRoute{
	commandRoute{
		Route:   "/api/usageschedule",
		Method:  "POST",
		Handler: addUsageScheduleHandler,
	},
	commandRoute{
		Route:   "/api/usageschedule/{id}/addusage",
		Method:  "PUT",
		Handler: addUsageAmountHandler,
	},
}

// NewRouter builds the scheduled usage router
func NewRouter(repositoryFactory func() repository.IRepository) http.Handler {

	r := mux.NewRouter().StrictSlash(true)
	for _, cr := range routes {
		r.HandleFunc(cr.Route, addMiddleware(repositoryFactory, cr.Handler)).Methods(cr.Method)
	}
	r1 := handlers.CORS(
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
	)(r)

	return r1
}

func addMiddleware(repositoryFactory func() repository.IRepository, h commandHandler) http.HandlerFunc {
	return errorHandler(injectStorage(repositoryFactory, h))
}

func errorHandler(f requestResponseWithErrorFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("handling %q: %v", r.RequestURI, err)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}
}

func injectStorage(repositoryFactory func() repository.IRepository, f commandHandler) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		rep := repositoryFactory()
		content, err := f(r, rep)
		if err != nil {
			return
		}
		if content != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err = json.NewEncoder(w).Encode(content); err != nil {
				return
			}
		}
		return
	}
}
