package storage

import (
	"fmt"
	"sync"
	"time"
)

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

type Data struct {
	Duration int64 `json:"duration"`
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
func (s *Task) MapDelete() { //delete by Task struct
	delete(MapByID, s.ID)
}

func (t *Data) AddTask() (id int) {
	MapMutex.Lock() //lock for ID
	task := Task{
		ID:          len(MapByID) + 1,
		Date:        time.Now().Unix(),
		AllDuration: t.Duration,
		Remaining:   t.Duration,
		IsComplete:  false,
	}
	task.MapCreate()
	MapMutex.Unlock()

	go func() { // remaining -= 2;update.
		for task.Remaining > 0 {
			time.Sleep(time.Second * 2)

			if _, exists := MapByID[task.ID]; exists {
				task.Remaining -= 2
				task.MapUpdate()
			} else {
				break
			}
		}
		if task.Remaining <= 0 {
			fmt.Printf("Task %d successfully end\n", task.ID)
			task.IsComplete = true
			task.MapUpdate()
		} else {
			fmt.Printf("Task %d was deleted\n", task.ID)
		}
	}()
	return task.ID
}

func GetTasks() (tasklist []TaskOutput) {
	var tasks []TaskOutput
	for _, v := range MapByID {
		percent := (1 - (float32(v.Remaining) / float32(v.AllDuration))) * 100
		task := TaskOutput{
			ID:         v.ID,
			Date:       time.Unix(v.Date, 0).Format("2006-01-02 15:04"),
			Remaining:  v.Remaining,
			Percent:    int(percent),
			IsComplete: v.IsComplete,
		}
		tasks = append(tasks, task)
	}
	return tasks
}
