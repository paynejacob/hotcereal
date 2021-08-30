package store

import "io"

type Store interface {
	Get(key Key) ([]byte, error)
	List(prefix TypeKey, process func([]byte) error) error
	ReadLazy(key FieldKey, w io.Writer) error

	WriteLazy(key FieldKey, r io.Reader) error
	Save(key Key, data []byte) error
	BulkSave(map[Key][]byte) error

	Delete(key ...Key) error

	Close() error
}
