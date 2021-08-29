package memory

import (
	"strings"
	"sync"
)

type Store struct {
	mu sync.RWMutex

	data map[string][]byte
}

func (s *Store) Get(key string) ([]byte, error) {
	var rval []byte

	s.mu.RLock()
	rval = s.data[key]
	s.mu.RUnlock()

	return rval, nil
}

func (s *Store) List(prefix string, process func([]byte) error) error {
	var err error

	s.mu.RLock()
	for k, v := range s.data {
		if strings.HasPrefix(k, prefix) {
			err = process(v)
			if err != nil {
				break
			}
		}
	}
	s.mu.RUnlock()

	return err
}

func (s *Store) Save(key string, bytes []byte) error {
	s.mu.Lock()
	s.data[key] = bytes
	s.mu.RUnlock()

	return nil
}

func (s *Store) BulkSave(m map[string][]byte) error {
	s.mu.Lock()
	for k, v := range m {
		s.data[k] = v
	}
	s.mu.Unlock()

	return nil
}

func (s *Store) Delete(key ...string) error {
	s.mu.Lock()
	for _, k := range key {
		delete(s.data, k)
	}
	s.mu.Unlock()

	return nil
}

func (s *Store) Close() error {
	return nil
}
