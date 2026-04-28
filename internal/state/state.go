package state

import "time"

type Row struct {
	Number int
	Date   time.Time
	Data   []int
	Sum    int
	Even   int
	Odd    int
}

type State struct {
	Data        []Row
	Status      status
	Error       error
	subscribers []func(s *State)
}

func (s *State) SetError(err error) {
	s.Status = Error
	s.Error = err
	s.Data = nil
	s.notify()
}

func (s *State) SetSuccess(data []Row) {
	if data != nil {
		s.Data = data
	}
	s.Error = nil
	s.Status = Ready
	s.notify()
}

func (s *State) SetLoading() {
	s.Status = Loading
	s.notify()
}

func (s *State) Subscribe(f ...func(*State)) {
	s.subscribers = append(s.subscribers, f...)
}

func (s *State) notify() {
	for _, f := range s.subscribers {
		f(s)
	}
}
