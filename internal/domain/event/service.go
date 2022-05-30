package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	FindMap(la1, lo1, la2, lo2 float32) ([]Event, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) FindMap(la1, lo1, la2, lo2 float32) ([]Event, error) {
	allEvents, err := (*s.repo).FindAll()
	var countForMap int
	countForMap = 0
	for i := int64(0); i < EventsCount; i++ {
		if la1 < allEvents[i].Lat && allEvents[i].Lat < la2 &&
			lo1 < allEvents[i].Long && allEvents[i].Long < lo2 {
			countForMap++
		}
	}
	mapEvents := make([]Event, countForMap)
	countForMap = 0
	for i := int64(0); i < EventsCount; i++ {
		if la1 < allEvents[i].Lat && allEvents[i].Lat < la2 &&
			lo1 < allEvents[i].Long && allEvents[i].Long < lo2 {
			mapEvents[countForMap] = allEvents[i]
			countForMap++
		}
	}
	return mapEvents, err
}
