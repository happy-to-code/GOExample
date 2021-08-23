package badgerDB

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var db Store

func TestMain(m *testing.M) {
	path := "E:\\20.06.16Project\\GOExample\\src\\20210716Demo\\badgerDB\\badger.db"
	os.RemoveAll(path)

	db = NewDB(path)
	m.Run()
}

func TestGetWrite(t *testing.T) {
	kv := KV{
		Key:   []byte("test1"),
		Value: []byte("value1"),
	}

	err := db.Write(kv)
	require.NoError(t, err)
	res, err := db.Get(kv.Key)
	fmt.Printf("-->%s\n", res)
	fmt.Printf("-->%s\n", res.Value)
	fmt.Println("-------------------")
	v, e := db.GetValueByKey(kv.Key)
	if e != nil {
		panic(e)
	}
	fmt.Println("=========>", string(v))
}
