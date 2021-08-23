package badgerDB

import (
	"fmt"
	"os"
	"testing"
)

var db Store

func TestMain(m *testing.M) {
	path := "E:\\20.06.16Project\\GOExample\\src\\A-2021-08-11\\DBTest\\BadgerDB\\badger.db"
	os.RemoveAll(path)

	db = NewDB(path)
	m.Run()
}

func TestGetWrite(t *testing.T) {

	for i := 0; i < 1000000; i++ {
		kv := KV{
			Key:   []byte(fmt.Sprintf("key_%d", i)),
			Value: []byte(fmt.Sprintf("value_%d", i)),
		}
		db.Write(kv)
	}

}

// --- PASS: TestGetWrite (21.70s)
