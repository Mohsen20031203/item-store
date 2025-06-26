package db

type Client interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Update(map[string][]byte) error
	List() (map[string][]byte, error)
	Close() error
}
