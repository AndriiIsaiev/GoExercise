package event

import (
	"fmt"
	"math/rand"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, EventsCount)
	for i := int64(0); i < EventsCount; i++ {
		events[i] = Event{
			Id:    i + 1,
			Title: fmt.Sprintf("Event #%d", i+1),
			Lat:   rand.Float32() * 90,
			Long:  rand.Float32()*360 - 180,
		}
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	if id <= EventsCount {
		return &Event{
			Id:    id,
			Title: fmt.Sprintf("Event #%d", id),
			Lat:   rand.Float32()*180 - 90,
			Long:  rand.Float32()*360 - 180,
		}, nil
	} else {
		return nil, nil
	}
}
