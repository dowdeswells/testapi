package router

import (
    "log"
    "errors"
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
    //"github.com/gorilla/context"
    "github.com/dowdeswells/testapi/domain"
    "github.com/dowdeswells/testapi/repository"
)

type adapter func(http.Handler) http.Handler
type requestResponseFunc func(http.ResponseWriter, *http.Request)
type requestResponseWithErrorFunc func(http.ResponseWriter, *http.Request) error

// func adapt(h http.Handler, adapters ...adapter) http.Handler {
//   for _, adapter := range adapters {
//     h = adapter(h)
//   }
//   return h
// }

// BuildRouter builds the scheduled usage router
func BuildRouter() http.Handler {

    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/api/usageschedule/{id}", errorHandler(injectStorage(getHandler))).Methods("GET")
    r.HandleFunc("/api/usageschedule", errorHandler(injectStorage(postHandler))).Methods("POST")
    r1 := handlers.CORS(
        handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
        )(r)

    return r1
}

func addMiddleware() {

}

func errorHandler(f requestResponseWithErrorFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := f(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            log.Printf("handling %q: %v", r.RequestURI, err)
        }
    }
}

func injectStorage(f func(http.ResponseWriter, *http.Request, repository.IRepository) error) func(http.ResponseWriter, *http.Request) error {
    return func(w http.ResponseWriter, r *http.Request) error {
        rep := repository.NewRepository()
        err := f(w, r, rep)
        return err
    }
}

func getHandler(w http.ResponseWriter, r *http.Request, rep repository.IRepository) (err error) {
    id, err := getRouteID(r)
    if err != nil {
        return
    }

    log.Println("Get usageschedule start")
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    v, err := rep.GetByID(id)
    if (err != nil) {
        return
    }
    if err = json.NewEncoder(w).Encode(v); err != nil {
        return
    }
    w.WriteHeader(http.StatusOK)
    log.Println("Get usageschedule end")
    return
}

func postHandler(w http.ResponseWriter, r *http.Request, rep repository.IRepository) (err error) {

    usageschedule := new(domain.UsageSchedule)

    err = readBody(r, usageschedule)
    if (err != nil) {
        return
    }
    fmt.Println(usageschedule.StartDate);
    w.WriteHeader(http.StatusOK)
    return
}

func addUsageAmountHandler(w http.ResponseWriter, r *http.Request, rep repository.IRepository) (err error) {

    id, err := getRouteID(r)
    if err != nil {
        return
    }

    cmd := new(domain.AddScheduledAmountCmd)
    err = readBody(r, cmd)
    if err != nil {
        return
    }

    u, err := rep.GetByID(id)
    if (err != nil) {
        return
    }

    br := u.AddScheduledAmount(cmd.EndDate, cmd.UsageAmount)
    if (br.HasBrokenRules()){
        err = br
        return
    }

    w.WriteHeader(http.StatusOK)
    return
}



func readBody(r *http.Request, v interface{}) (err error) {

    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(v)
    // if error != nil {
    //     log.Println(error.Error())
    //     http.Error(res, error.Error(), http.StatusInternalServerError)
    //     return
    // }
    return err
}

func getRouteID(r *http.Request) (id string, err error) {
    vars := mux.Vars(r)
    id = vars["id"]
    if id == "" {
        err = errors.New("Incorrect route parameters")
    }
    return id, err
}
