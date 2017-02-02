package mongo

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type Collection struct {
	Database *Database
	Session  *mgo.Collection
	name     string
}

func (this *Collection) Connect() {
	session := *this.Database.Session.C(this.name)
	this.Session = &session
}

func NewCollectionSession(name string) *Collection {

	var collection = Collection{
		Database: newDBSession(beego.AppConfig.String("mongo::database")),
		name:     name,
	}

	collection.Connect()

	return &collection
}

func (this *Collection) Close() {
	service.Close(this)
}
