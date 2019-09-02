package dbservice

import (
	"fmt"
	"log"

	"github.com/qjouda/dignity-platform/backend/datatype"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	//ErrRecordDoesntExist record does not exist error
	ErrRecordDoesntExist = fmt.Errorf("Record does not exist")

	//ErrMultipleResults error for getting multipe results when expecting one
	ErrMultipleResults = fmt.Errorf("Multiple results found")
)

// MongoDB defines MongoDB type
type MongoDB struct {
	Session *mgo.Session
	DBHost  string
	DBName  string
}

//NewMongoDB constructs a DB
func NewMongoDB(cfg datatype.Config) *MongoDB {
	return &MongoDB{
		DBHost: cfg.MongoDBHost,
		DBName: cfg.MongoDBName,
	}
}

//mustGetDB gets an mgo.Session copy
func (db *MongoDB) mustGetDB() *mgo.Session {
	if db.Session == nil {
		//session maintains a connection pool
		var err error
		db.Session, err = mgo.Dial(db.DBHost)
		if err != nil {
			log.Fatal(err)
		}
	}
	return db.Session.Copy()
}

//Count find count of docs matching selector
func (db *MongoDB) Count(collection string, selector interface{}) (int, error) {
	s := db.mustGetDB()
	defer s.Close()
	c := s.DB(db.DBName).C(collection)
	return c.Find(selector).Count()
}

//FindOne finds one document
func (db *MongoDB) FindOne(collection string, selector interface{}, fields interface{}, model interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	c := s.DB(db.DBName).C(collection)
	doc := c.Find(selector).Select(fields)
	count, err := doc.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrRecordDoesntExist
	}
	if count > 1 {
		return ErrMultipleResults
	}

	return doc.One(model)
}

//FindAll find all documents matching selector
func (db *MongoDB) FindAll(
	collection string,
	selector interface{},
	fields interface{},
	sort string,
	model interface{},
) error {
	s := db.mustGetDB()
	defer s.Close()
	c := s.DB(db.DBName).C(collection)
	if sort == "" {
		sort = "$natural"
	}
	return c.Find(selector).Select(fields).Sort(sort).All(model)
}

//Create persists a new document
func (db *MongoDB) Create(collection string, doc interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Insert(doc)
}

// CreateMany persists an array of documents
func (db *MongoDB) CreateMany(collection string, docs []interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Insert(docs...)
}

// Upsert updates or persists a document
func (db *MongoDB) Upsert(collection string, selector interface{}, doc interface{}) (interface{}, error) {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Upsert(selector, doc)
}

//UpdateOne updates a single document matching selector
func (db *MongoDB) UpdateOne(collection string, selector interface{}, update interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Update(selector, bson.M{"$set": update})
}

//UpdateAll updates all documents matching selector
func (db *MongoDB) UpdateAll(collection string, selector interface{}, update interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	_, err := s.DB(db.DBName).C(collection).UpdateAll(selector, bson.M{"$set": update})
	return err
}

//RemoveOne deletes a single document matching selector
func (db *MongoDB) RemoveOne(collection string, selector interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Remove(selector)
}

//RemoveAll deletes all documents matching selector
func (db *MongoDB) RemoveAll(collection string, selector interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	_, err := s.DB(db.DBName).C(collection).RemoveAll(selector)
	return err
}

// EnsureIndex creates an index in mongo
func (db *MongoDB) EnsureIndex(collection string, indexDef interface{}) error {
	s := db.mustGetDB()
	defer s.Close()
	ix := indexDef.(mgo.Index)
	return s.DB(db.DBName).C(collection).EnsureIndex(ix)
}

// DropIndex drops a mongo index
func (db *MongoDB) DropIndex(collection string, key []string) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).DropIndex(key...)
}

// DropCollection drops a mongo collection
func (db *MongoDB) DropCollection(collection string) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).DropCollection()
}

// Aggregate executes an aggregation pipeline query
func (db *MongoDB) Aggregate(collection string, pipeline interface{}, result *[]bson.M) error {
	s := db.mustGetDB()
	defer s.Close()
	return s.DB(db.DBName).C(collection).Pipe(pipeline).All(result)
}
