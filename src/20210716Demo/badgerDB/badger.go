package badgerDB

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"time"
)

type badgerDB struct {
	db *badger.DB
}

// NewDB example path: ./db
func NewDB(path string) Store {
	badgerdb := new(badgerDB)
	var err error

	opt := badger.DefaultOptions(path)
	// 截断日志文件删除错误数据，以便数据库能实现异常退出的再次访问
	opt.Truncate = true
	badgerdb.db, err = badger.Open(opt)
	if err != nil {
		panic(err)
	}
	return badgerdb
}

// 单个kv 永久写入
func (db *badgerDB) Write(kv KV) error {
	txn := db.db.NewTransaction(true)
	txn.SetEntry(badger.NewEntry(kv.Key, kv.Value))
	return txn.Commit()
}

// WriteAt 单个kv 写入带有效期，ts为有效期时长
func (db *badgerDB) WriteAt(kv KV, ts time.Duration) error {
	txn := db.db.NewTransaction(true)
	txn.SetEntry(badger.NewEntry(kv.Key, kv.Value).WithTTL(ts))
	return txn.Commit()
}

func (db *badgerDB) Confirm(kvs ...KV) error {
	return db.db.Update(func(txn *badger.Txn) error {
		for _, kv := range kvs {
			err := txn.SetEntry(badger.NewEntry(kv.Key, kv.Value))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// WriteBatch kv 批量写入
func (db *badgerDB) WriteBatch(kvs ...KV) error {
	if len(kvs) < 1 {
		return fmt.Errorf("%s", "empty kvs!")
	}
	txn := db.db.NewTransaction(true)
	defer txn.Discard()

	for _, kv := range kvs {
		txn.SetEntry(badger.NewEntry(kv.Key, kv.Value))
	}

	return txn.Commit()
}

func (db *badgerDB) Get(key []byte) (*KV, error) {
	var value []byte
	if err := db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		if value, err = item.ValueCopy(value); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return &KV{Key: key, Value: value}, nil
}

// GetValueByKey 通过key 获取value
func (db *badgerDB) GetValueByKey(k []byte) ([]byte, error) {
	var value []byte
	if err := db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(k)
		if err != nil {
			return err
		}
		if value, err = item.ValueCopy(value); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return value, nil
}

// GetBatch 批量get
func (db *badgerDB) GetBatch(keys [][]byte) ([]KV, error) {
	kvs := make([]KV, len(keys))
	if err := db.db.View(func(txn *badger.Txn) error {
		for k, key := range keys {
			item, err := txn.Get(key)
			if err != nil {
				return err
			}
			if value, err := item.ValueCopy(nil); err != nil {
				return err
			} else {
				kvs[k] = KV{Key: key, Value: value}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return kvs, nil
}
func (db *badgerDB) GetPrefix(pre []byte) ([]KV, error) {
	opt := badger.DefaultIteratorOptions
	opt.PrefetchValues = false
	kvs := []KV{}
	txn := db.db.NewTransaction(false)
	it := txn.NewIterator(opt)
	defer it.Close()
	for it.Seek(pre); it.ValidForPrefix(pre); it.Next() {
		item := it.Item()
		key := item.Key()
		val, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}
		newkey := make([]byte, len(key))
		copy(newkey, key)
		kvs = append(kvs, KV{Key: newkey, Value: val})
	}

	return kvs, nil
}

func (db *badgerDB) Delete(key []byte) error {
	txn := db.db.NewTransaction(true)
	if err := txn.Delete(key); err != nil {
		return err
	}

	return txn.Commit()
}

func (db *badgerDB) DeleteBatch(keys [][]byte) error {
	txn := db.db.NewTransaction(true)
	for _, key := range keys {
		if err := txn.Delete(key); err != nil {
			return err
		}
	}

	return txn.Commit()
}

func (db *badgerDB) Close() error {
	return db.db.Close()
}
