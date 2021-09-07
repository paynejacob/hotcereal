package memory

import (
	"github.com/paynejacob/hotcereal/pkg/store"
	"io"
	"strings"
	"sync"
)

type Store struct {
	mu sync.RWMutex

	data map[string][]byte
}

func New() *Store {
	return &Store{
		data: make(map[string][]byte, 0),
	}
}

func (s *Store) Get(id store.Key) ([]byte, error) {
	var rval []byte

	s.mu.RLock()
	rval = s.data[id.String()]
	s.mu.RUnlock()

	return rval, nil
}

func (s *Store) List(prefix store.TypeKey, process func([]byte) error) error {
	var err error

	s.mu.RLock()
	for k, v := range s.data {
		if strings.HasPrefix(k, prefix.String()) {
			err = process(v)
			if err != nil {
				break
			}
		}
	}
	s.mu.RUnlock()

	return err
}

func (s *Store) ReadLazy(key store.FieldKey, w io.Writer) error {
	var rval []byte

	s.mu.RLock()
	rval = s.data[key.String()]
	s.mu.RUnlock()

	_, err := w.Write(rval)
	return err
}

func (s *Store) WriteLazy(key store.FieldKey, r io.Reader) error {
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.data[key.String()] = content
	s.mu.Unlock()

	return err
}

func (s *Store) Save(key store.Key, data []byte) error {
	s.mu.Lock()
	s.data[key.String()] = data
	s.mu.Unlock()

	return nil
}

func (s *Store) BulkSave(m map[store.Key][]byte) error {
	s.mu.Lock()
	for k, v := range m {
		s.data[k.String()] = v
	}
	s.mu.Unlock()

	return nil
}

func (s *Store) Delete(key ...store.Key) error {
	s.mu.Lock()
	for _, k := range key {
		delete(s.data, k.String())
	}
	s.mu.Unlock()

	return nil
}

func (s *Store) Close() error {
	return nil
}
