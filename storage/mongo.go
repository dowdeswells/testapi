package storage

import (
    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "os"
)


type storeImp struct {
    session *mgo.Session
    db string
    col string
}

type mongoAction func(c *mgo.Collection)

var mongoURI = "mongodb://mongodb"

func init() {
    if m := os.Getenv("MONGO_URI"); m != "" {
		mongoURI = m
	}
}

// CreateMongo builds the underlying storage mechanism
func CreateMongo(db string, col string) (s IStorage, err error) {
    mongo, err := startMongo()
    if (err == nil) {
        mongo.db = db
        mongo.col = col
    }
    return mongo, err
}


func (s *storeImp) Get(id string, o interface{}) (err error) {
    s.execMongoAction(func (c *mgo.Collection) {
        err = c.Find(bson.M{"_id": id}).One(o)
    })
    return
}

func(s *storeImp) CreateKey() (id string) {
    id = bson.NewObjectId().Hex()
    return
}

func(s *storeImp) Insert(u interface{}) (err error) {
    s.execMongoAction(func (c *mgo.Collection) {
        err = c.Insert(u)
    })
    return
}

func(s *storeImp) Update(id string, u interface{}) (err error) {
    s.execMongoAction(func (c *mgo.Collection) {
        err = c.UpdateId(id, u)
    })
    return
}


func startMongo() (s *storeImp, err error) {
    ses, err := mgo.Dial(mongoURI)

    if (err == nil) {
        ses.SetMode(mgo.Monotonic, true)
        s := &storeImp{
            session: ses,
        }
        return s, err
    }
    return nil, err
}


func(s *storeImp) execMongoAction(f mongoAction) {
    session := s.session.Copy()
    defer session.Close()
    c := session.DB(s.db).C(s.col)
    f(c)
}
