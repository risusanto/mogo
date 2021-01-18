package mogo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client
var MongoDB *mongo.Database

// Model interface contain method must be implemented by each model
type Model interface {
	Create(model Model, opts ...*options.InsertOneOptions) error
	Delete(model Model) error
	Find(model Model, filter interface{}, results interface{}, opts ...*options.FindOptions)
	FindByID(id string, model Model) error
	GetID() interface{}
	PrepareID(id interface{}) (interface{}, error)
	SetID(id interface{})
	Update(model Model, opts ...*options.UpdateOptions) error
}

// CollectionNameGetter interface contain method to return
// collection name of model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

// CollectionGetter interface contain method to return
// collection of model.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

type BaseModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

type DefaultFilter struct {
	Skip    int    `json:"skip"`
	Take    int    `json:"take"`
	Filters bson.M `json:"filters"`
}

type DBConfig struct {
	MongoURI string
	DBName   string
}

func NewConnection(config DBConfig) (*mongo.Client, *mongo.Database, error) {
	err := SetDefaultConfig(nil, config.DBName, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	_, client, db, _ := DefaultConfigs()
	MongoDB = db
	MongoDBClient = client

	return client, db, err
}

func (m *BaseModel) Create(model Model, opts ...*options.InsertOneOptions) error {
	err := Coll(model).Create(model, opts...)
	if err != nil {
		return err
	}
	return err
}

func (m *BaseModel) Find(model Model, filter interface{}, results interface{}, opts ...*options.FindOptions) {
	err := Coll(model).SimpleFind(results, filter, opts...)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *BaseModel) FindByID(id string, model Model) error {
	return Coll(model).FindByID(id, model)
}

func (m *BaseModel) Delete(model Model) error {
	return Coll(model).Delete(model)
}

func (m *BaseModel) Update(model Model, opts ...*options.UpdateOptions) error {
	return Coll(model).Update(model, opts...)
}
