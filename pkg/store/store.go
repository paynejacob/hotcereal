package store

type Store interface {
	Get(key string) ([]byte, error)
	List(prefix string, process func([]byte) error) error

	Save(string, []byte) error
	BulkSave(map[string][]byte) error

	Delete(key ...string) error

	Close() error
}
