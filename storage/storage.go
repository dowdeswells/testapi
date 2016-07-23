package storage

type IState interface {
    GetID() string
    SetID(id string)
}

type IStorage interface {
    Get(id string, o interface{}) (error)
    CreateKey() (string)
    Insert(r interface{}) (error)
    Update(id string, r interface{}) (error)
}

