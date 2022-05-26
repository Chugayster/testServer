package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}
		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		if event != nil {
			err = success(w, event)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
		} else {
			err = notFound(w)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
		}
	}

}

func (c *EventController) DbUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Wrong DbUpadate post", err)
		}
		var updateEvent event.Event
		err = json.Unmarshal(body, &updateEvent)
		if err != nil {
			log.Println("Json read", err)
		}
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DbUpdate()): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DbUpdate(): %s", err)
			}
			return
		}
		event, err := (*c.service).DbUpdate(id, &updateEvent)
		if err != nil {
			fmt.Printf("EventController.DbUpdate(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DbUpdate(): %s", err)
			}
			return
		}
		if event != nil {
			err = success(w, event)
			if err != nil {
				fmt.Printf("EventController.Dbupdate(): %s", err)
			}
		} else {
			err = notFound(w)
			if err != nil {
				fmt.Printf("EventController.Dbupdate(): %s", err)
			}
		}

	}
}

func (c *EventController) DbCreated() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Wrong DbCreate post", err)
		}
		var createEvent event.Event
		err = json.Unmarshal(body, &createEvent)
		if err != nil {
			log.Println("Json read", err)
		}
		event, err := (*c.service).DbCreated(&createEvent)
		if err != nil {
			log.Printf("EventController.DbCreate(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("EventController.DbCreate(): %s", err)
			}
			return
		}
		err = created(w, event)
		if err != nil {
			fmt.Printf("EventController.DbCreate(): %s", err)
		}
	}
}

func (c *EventController) DbDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DbDelete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DbDelete(): %s", err)
			}
			return
		}
		status, err := (*c.service).DbDelete(id)
		if err != nil {
			fmt.Printf("EventController.DbDelete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DbDelete(): %s", err)
			}
			return
		}
		if status {
			err = delete(w, fmt.Sprintf("id %d is deleted", id))
			if err != nil {
				fmt.Printf("EventController.DbDelete(): %s", err)
			}
		} else {
			err = notFound(w)
			if err != nil {
				fmt.Printf("EventController.Dbupdate(): %s", err)
			}

		}
	}
}
