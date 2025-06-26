package db

type Client interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Update(map[string][]byte) error
	GetAll() (map[string][]byte, error)
	Close() error
}
