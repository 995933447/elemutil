package elemutil

import "sync"

type Set struct {
	m sync.Map
}

func NewSet() *Set {
	return &Set{}
}

func (s Set) Exist(val interface{}) bool {
	_, ok := s.m.Load(val)
	return ok
}

func (s *Set) Put(val interface{}) {
	s.m.Store(val, struct {}{})
}
