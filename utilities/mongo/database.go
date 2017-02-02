package mongo

import "gopkg.in/mgo.v2"

type Database struct {
	BaseSession *mgo.Session
	name        string
	Session     *mgo.Database
}

func (this *Database) Connect() {
	this.BaseSession = service.Session()
	session := *this.BaseSession.DB(this.name)
	this.Session = &session
}

func newDBSession(name string) *Database {
	var database = Database{
		name: name,
	}

	database.Connect()

	return &database
}
