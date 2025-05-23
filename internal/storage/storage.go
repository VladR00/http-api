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
	MapByID[s.ID] = *s
}
func (s *Quotes) MapDelete() { //By ID
	delete(MapByID, s.ID)
}
