package helper

import "time"

type TaskManager struct {
	availableTime map[string]time.Time
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		availableTime: map[string]time.Time{},
	}
}

func (s *TaskManager) Finish(task string) {
	s.availableTime[task] = time.Time{}
}

func (s *TaskManager) CheckLater(task string) {
	s.availableTime[task] = time.Now().Add(time.Hour)
}

func (s *TaskManager) Available(task string) bool {
	t, ok := s.availableTime[task]
	if !ok {
		return true
	}
	if t.IsZero() {
		return false
	}
	return time.Now().After(t)
}
