package datatype

import "gopkg.in/mgo.v2/bson"

// MongoDB interface for DB
type MongoDB interface {
	Count(collection string, selector interface{}) (int, error)
	FindOne(collection string, selector interface{}, fields interface{},
		model interface{}) error
	FindAll(collection string, selector interface{}, fields interface{},
		sort string, model interface{}) error
	Create(collection string, document interface{}) error
	CreateMany(collection string, documents []interface{}) error
	Upsert(collection string, selector interface{},
		document interface{}) (interface{}, error)
	UpdateOne(collection string, selector interface{}, update interface{}) error
	UpdateAll(collection string, selector interface{}, update interface{}) error
	RemoveOne(collection string, selector interface{}) error
	RemoveAll(collection string, selector interface{}) error
	EnsureIndex(collection string, indexDef interface{}) error
	DropIndex(collection string, key []string) error
	DropCollection(collection string) error
	Aggregate(collection string, pipeline interface{}, result *[]bson.M) error
}
