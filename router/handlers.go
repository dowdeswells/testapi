package router

import (
    "errors"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/dowdeswells/testapi/domain"
    "github.com/dowdeswells/testapi/repository"
)

//IDContent alows encoding Id as json content
type IDContent struct {
    ID  string `json:"id"`
}

func addUsageScheduleHandler(r *http.Request, rep repository.IRepository) (content interface{}, err error)  {

    cmd := new(domain.AddUsageScheduleCmd)
    err = readBody(r, cmd)
    if err != nil {
        return
    }
    u := domain.UsageSchedule{}
    id, err := executeAndSave(u, cmd, rep)

    if (err != nil) {
        return
    }
    content = convertID(id)
    return
}

func addUsageAmountHandler(r *http.Request, rep repository.IRepository) (content interface{}, err error) {

    cmd := new(domain.AddScheduledAmountCmd)
    err = readBody(r, cmd)
    if err != nil {
        return
    }

    content, err = updateCommandHandler(r,rep,cmd)

    if (err != nil) {
        return
    }
    return
}

func updateCommandHandler(r *http.Request, rep repository.IRepository, cmd domain.IUsageScheduleCommand) (content interface{}, err error) {

    id, err := getRouteID(r)
    if err != nil {
        return
    }

    u, err := rep.GetByID(id)
    if (err != nil) {
        return
    }

    id, err = executeAndSave(u, cmd, rep)
    if (err != nil) {
        return
    }
    content = convertID(id)
    return
}

func executeAndSave(u domain.UsageSchedule, cmd domain.IUsageScheduleCommand, rep repository.IRepository) (id string, err error){

    u2, err := cmd.Execute(u)

    if (err == nil) {
        id, err = rep.Save(u2)
    }
    return
}


func readBody(r *http.Request, v interface{}) (err error) {

    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(v)
    return err
}

func getRouteID(r *http.Request) (id string, err error) {
    vars := mux.Vars(r)
    id = vars["id"]
    if id == "" {
        err = NoIdInRequestError()
    }
    return id, err
}

func NoIdInRequestError() (err error) {
    err = errors.New("Incorrect route parameters - no id present")
    return
}

func convertID(id string) IDContent {
    return IDContent{
        ID:id,
    }
}