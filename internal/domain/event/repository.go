package event

import (
	"github.com/upper/db/v4"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	DbUpdate(event *Event) (*Event, error)
	DbCreated(event *Event) (*Event, error)
	DbDelete(id int64) error
}

const EventsCount int64 = 10

type repository struct {
	ses db.Session // Some internal data
}

func NewRepository(ses *db.Session) Repository {
	return &repository{*ses}
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event
	s := r.ses.SQL().SelectFrom("events")
	s = s.OrderBy("id")
	err := s.All(&events)
	if err != nil {
		log.Println("allparty.Find", err)
		return nil, err
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	var event Event
	allevents := r.ses.Collection("events").Find(id)
	err := allevents.One(&event)
	if err != nil {
		log.Println("Find One", err)
		return nil, err
	}
	return &event, nil

}

func (r *repository) DbUpdate(updateEvent *Event) (*Event, error) {
	var event Event
	s := r.ses.Collection("events").Find(updateEvent.Id)
	err := s.One(&event)
	if err != nil {
		log.Println("Find One", err)
		return nil, err
	}
	event.Name = updateEvent.Name
	event.Location = updateEvent.Location
	if err := s.Update(event); err != nil {
		log.Printf("Update: %v\n", err)
		return nil, err
	}
	return &event, nil

}

func (r *repository) DbCreated(event *Event) (*Event, error) {
	err := r.ses.Collection("events").InsertReturning(event)
	if err != nil {
		log.Println("InsertReturning: ", err)
		return nil, err
	}
	return event, nil
}

func (r *repository) DbDelete(id int64) error {
	var event Event
	s := r.ses.Collection("events").Find(id)
	err := s.One(&event)
	if err != nil {
		log.Println("Find One", err)
		return err
	}
	if err := s.Delete(); err != nil {
		log.Println("Delete: ", err)
		return err
	}
	return nil
}
