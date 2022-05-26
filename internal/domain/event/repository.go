package event

import (
	"github.com/upper/db/v4"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	DbUpdate(id int64, event *Event) (*Event, error)
	DbCreated(event *Event) (*Event, error)
	DbDelete(id int64) (bool, error)
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
		log.Fatal("allparty.Find", err)
	}
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	var event Event
	allevents := r.ses.Collection("events")


		s := allevents.Find(id)
		err := s.One(&event)
		if err != nil {
			log.Println("Find One", err)
			return nil, nil
		}
		return &event, nil
		

}

func (r *repository) DbUpdate(id int64, updateEvent *Event) (*Event, error) {
	var event Event
	allevents := r.ses.Collection("events")
	s := allevents.Find(id)
	err := s.One(&event)
	if err != nil {
		log.Println("Find One", err)
		return nil, nil
	}
	if updateEvent.Name != "" {
		event.Name = updateEvent.Name
	}
	if updateEvent.Location != "" {
		event.Location = updateEvent.Location
	}
	if err := s.Update(event); err != nil {
		log.Printf("Update: %v\n", err)
	}
	return &event, nil

}

func (r *repository) DbCreated(event *Event) (*Event, error) {
	allevents := r.ses.Collection("events")
	err := allevents.InsertReturning(event)
	if err != nil {
		log.Println("InsertReturning: ", err)
	}
	return event, nil
}

func (r *repository) DbDelete(id int64) (bool, error) {
	var event Event
	allevents := r.ses.Collection("events")
	s := allevents.Find(id)
	err := s.One(&event)
	if err != nil {
		log.Println("Find One", err)
		return false, nil
	}

	if err := s.Delete(); err != nil {
		log.Printf("Update: %v\n", err)
	}
	return true, nil
}
