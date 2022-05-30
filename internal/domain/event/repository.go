package event

import (
	"fmt"
	"math/rand"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

const EventsCount int64 = 20

type repository struct {
	baseEvents []Event
}

func NewRepository() Repository {
	var r repository
	r.baseEvents = make([]Event, EventsCount)
	for i := int64(0); i < EventsCount; i++ {
		r.baseEvents[i] = Event{
			Id:    i + 1,
			Title: fmt.Sprintf("Event #%d", i+1),
			Lat:   rand.Float32() * 90,
			Long:  rand.Float32()*360 - 180,
		}
	}
	return &r
}

func (r *repository) FindAll() ([]Event, error) {
	return r.baseEvents, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	if id <= EventsCount && id > 0 {
		return &r.baseEvents[id-1], nil
	} else {
		return nil, nil
	}
}
