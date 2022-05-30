package controllers

import (
	"fmt"
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

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) FindMap() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		la1, err := parsePos(w, r, "la1")
		if err != nil {
			return
		}

		lo1, err := parsePos(w, r, "lo1")
		if err != nil {
			return
		}

		la2, err := parsePos(w, r, "la2")
		if err != nil {
			return
		}

		lo2, err := parsePos(w, r, "lo2")
		if err != nil {
			return
		}
		//		 strconv.ParseFloat(chi.URLParam(r, "la1"), 32)
		//if err != nil {
		//	fmt.Printf("EventController.FindMap(): %s", err)
		//	err = internalServerError(w, err)
		//	if err != nil {
		//		fmt.Printf("EventController.FindMap(): %s", err)
		//	}
		//	return
		//}

		//lo1, err := strconv.ParseFloat(chi.URLParam(r, "lo1"), 32)
		//if err != nil {
		//	fmt.Printf("EventController.FindMap(): %s", err)
		//	err = internalServerError(w, err)
		//	if err != nil {
		//		fmt.Printf("EventController.FindMap(): %s", err)
		//	}
		//	return
		//}
		//
		//la2, err := strconv.ParseFloat(chi.URLParam(r, "la2"), 32)
		//if err != nil {
		//	fmt.Printf("EventController.FindMap(): %s", err)
		//	err = internalServerError(w, err)
		//	if err != nil {
		//		fmt.Printf("EventController.FindMap(): %s", err)
		//	}
		//	return
		//}
		//
		//lo2, err := strconv.ParseFloat(chi.URLParam(r, "lo2"), 32)
		//if err != nil {
		//	fmt.Printf("EventController.FindMap(): %s", err)
		//	err = internalServerError(w, err)
		//	if err != nil {
		//		fmt.Printf("EventController.FindMap(): %s", err)
		//	}
		//	return
		//}

		event, err := (*c.service).FindMap(float32(la1), float32(lo1), float32(la2), float32(lo2))
		if err != nil {
			fmt.Printf("EventController.FindMap(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindMap(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindMap(): %s", err)
		}
	}
}

func parsePos(w http.ResponseWriter, r *http.Request, sPos string) (float64, error) {
	pos, err := strconv.ParseFloat(chi.URLParam(r, sPos), 32)
	if err != nil {
		fmt.Printf("EventController.FindMap(): %s", err)
		err = internalServerError(w, err)
		if err != nil {
			fmt.Printf("EventController.FindMap(): %s", err)
		}
		return 0, err
	}
	return pos, err
}
