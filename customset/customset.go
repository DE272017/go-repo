package customset

import "sync"

var exists = struct{}{}
var lock = sync.RWMutex{}

type Set struct {
	M map[interface{}]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.M = make(map[interface{}]struct{})
	return s
}

func (s *Set) Add(value interface{}) bool {
	lock.Lock()
	defer lock.Unlock()
	_, ok := s.M[value]
	if ok {
		return false
	} else {
		s.M[value] = exists
		return true
	}
}

func (s *Set) Remove(value interface{}) {
	delete(s.M, value)
}

func (s *Set) contains(value interface{}) bool {
	_, c := s.M[value]
	return c
}
