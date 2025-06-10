package storage

import "sync"

var (
	MapByID  = map[int]Task{} //Default Task map
	MapMutex sync.Mutex
)

type Task struct { //Default Task sctruct
	ID          int
	Date        int64
	AllDuration int64
	Remaining   int64
	IsComplete  bool
}

type TaskOutput struct { // Using in GET request.
	ID         int
	Date       string
	Remaining  int64
	Percent    int
	IsComplete bool
}

func (s *Task) MapCreate() { // if exist -> ID++ -> create
	if _, exists := MapByID[s.ID]; exists {
		s.ID = s.ID + 1
	}
	MapByID[s.ID] = *s
}
func (s *Task) MapUpdate() { // if exist -> update
	if _, exists := MapByID[s.ID]; exists {
		MapByID[s.ID] = *s
	}
}
func (s *Task) MapDelete() { //delete by id
	delete(MapByID, s.ID)
}
