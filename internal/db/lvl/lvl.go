package lvl

import (
	"errors"
	"item-store/internal/db"

	"github.com/syndtr/goleveldb/leveldb"
)

type LvlDB struct {
	DB *leveldb.DB
}

func NewDB(path string) (db.Client, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LvlDB{
		DB: db,
	}, err
}

func (l *LvlDB) Get(key string) ([]byte, error) {
	value, err := l.DB.Get([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, errors.New("key not found")
		}
		return nil, err
	}
	return []byte(value), nil
}

func (l *LvlDB) Put(key string, value []byte) error {
	err := l.DB.Put([]byte(key), value, nil)
	if err != nil {
		return err
	}
	return nil
}

func (l *LvlDB) Delete(key string) error {
	err := l.DB.Delete([]byte(key), nil)
	if err != nil {
		return err
	}
	return nil
}

func (l *LvlDB) Update(m map[string][]byte) error {
	for k, v := range m {
		err := l.DB.Put([]byte(k), v, nil)
		if err != nil {
			return err
		}
	}
	return nil

}

func (l *LvlDB) Close() error {
	err := l.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (l *LvlDB) GetAll() (map[string][]byte, error) {
	iter := l.DB.NewIterator(nil, nil)
	defer iter.Release()

	items := make(map[string][]byte)
	for iter.First(); iter.Valid(); iter.Next() {
		key := string(iter.Key())
		value := iter.Value()
		items[key] = value
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return items, nil
}
