package repository

import (
    "github.com/dowdeswells/testapi/storage"
    "github.com/dowdeswells/testapi/domain"
    "os"
)

//IRepository is Get and Save of UsageSchedule
type IRepository interface {
    GetByID (id string) (u domain.UsageSchedule, err error)
    Save (u domain.UsageSchedule) (id string, err error)
}

type repository struct {
    store storage.IStorage
}

var mDB = "store"
var mCOL = "usageSchedule"

func init() {
    if m := os.Getenv("MONGO_DB"); m != "" {
		mDB = m
	}
    if m := os.Getenv("MONGO_COL"); m != "" {
		mCOL = m
	}
}
//NewRepository is the factory for IRepository
func NewRepository() IRepository {
    s, _ := storage.CreateMongo(mDB, mCOL)
    r := &repository {
        store : s,
    }
    return r
}

func (r *repository) GetByID(id string) (u domain.UsageSchedule, err error) {
    err = r.store.Get(id, &u)
    return
}

func (r *repository) Save(u domain.UsageSchedule) (id string, err error) {
    if (u.ID == "") {
        id = r.store.CreateKey()
        u.ID = id
        err = r.store.Insert(u)
    } else {
        id = u.ID
        err = r.store.Update(id, u)
    }
    return
}



