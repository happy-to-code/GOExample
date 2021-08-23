package badgerDB

import "time"

type Store interface {
	Write(KV) error
	WriteAt(KV, time.Duration) error
	WriteBatch(...KV) error

	Confirm(...KV) error
	Get([]byte) (*KV, error)
	GetValueByKey([]byte) ([]byte, error)
	GetBatch([][]byte) ([]KV, error)
	GetPrefix(pre []byte) ([]KV, error)

	Delete([]byte) error
	DeleteBatch([][]byte) error

	Close() error
}
