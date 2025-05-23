package storage

import "sync"

var (
	MapByID  = map[int]Quotes{}
	MapMutex sync.Mutex
)

type Quotes struct {
	ID     int
	Author string
	Quote  string
}

func (s *Quotes) MapCreate() { //By ID
	if _, exists := MapByID[s.ID]; exists {
		s.ID = s.ID + 1
	}
	MapByID[s.ID] = *s
}
func (s *Quotes) MapDelete() { //By ID
	delete(MapByID, s.ID)
}
