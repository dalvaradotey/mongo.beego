package mongo

import "gopkg.in/mgo.v2"

type Service struct {
	BaseSession *mgo.Session
	Url         string
	open        int
	queue       chan int
}

var service Service

func (this *Service) New() error {
	var err error
	this.queue = make(chan int, maxPool)

	for i := 0; i < maxPool; i = i + 1 {
		this.queue <- 1
	}

	this.open = 0
	this.BaseSession, err = mgo.Dial(this.Url)

	return err
}

func (this *Service) Session() *mgo.Session {
	<-this.queue
	this.open++

	return this.BaseSession.Copy()
}

func (this *Service) Close(collection *Collection) {
	collection.Database.BaseSession.Close()
	this.queue <- 1
	this.open--
}
