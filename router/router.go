package router

import (
    "log"
    //"errors"
    "encoding/json"
    //"fmt"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
    //"github.com/gorilla/context"
    //"github.com/dowdeswells/testapi/domain"
    "github.com/dowdeswells/testapi/repository"
)

type requestResponseFunc func(http.ResponseWriter, *http.Request)
type requestResponseWithErrorFunc func(http.ResponseWriter, *http.Request) error

type commandHandler func (*http.Request, repository.IRepository) (interface{}, error)

type commandRoute struct {
    Route       string
    Method      string
    Handler     commandHandler
}

var routes = []commandRoute {
    commandRoute {
        Route: "/api/usageschedule",
        Method: "POST",
        Handler: addUsageScheduleHandler,
    },
    commandRoute {
        Route: "/api/usageschedule/{id}/addusage",
        Method: "PUT",
        Handler: addUsageAmountHandler,
    },
}

// BuildRouter builds the scheduled usage router
func NewRouter() http.Handler {

    r := mux.NewRouter().StrictSlash(true)
    for _, cr := range routes {
        r.HandleFunc(cr.Route, addMiddleware(cr.Handler)).Methods(cr.Method)
    }
    r1 := handlers.CORS(
        handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
        )(r)

    return r1
}

func addMiddleware(h commandHandler) http.HandlerFunc {
    return errorHandler(injectStorage(h));
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

func injectStorage(f commandHandler) func(http.ResponseWriter, *http.Request) (error) {
    return func(w http.ResponseWriter, r *http.Request) (err error) {
        rep := repository.NewRepository()
        content, err := f(r, rep)
        if (err != nil) {
            return
        }
        if (content != nil) {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            if err = json.NewEncoder(w).Encode(content); err != nil {
                return
            }
        }
        return
    }
}




