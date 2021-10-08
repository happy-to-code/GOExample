package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 发行
	txId := "txid123sdfs"

	go func() {
		time.Sleep(time.Second * 5)
		Signal(txId, nil)
	}()
	// Wait是同步操作 保证业务数据上链成功
	if err := Wait(string(txId), "hex", false); err != nil {
		fmt.Println("-->err:", err)
		return
	}
	fmt.Println("------------------last--------------")
}

const (
	V1 = "base64"
	V2 = "hex"
)

// Wait 同步操作等待.
func Wait(txid string, version string, isWait bool) error {
	timeout := 10

	if isWait {
		return nil
	}
	ch, err := Put(txid)
	if err != nil {
		return err
	}
	select {
	case err = <-ch:
		return err
	case <-time.After(time.Second * time.Duration(timeout)):
		Delete(txid)
		if version == V1 {
			txid = base64.StdEncoding.EncodeToString([]byte(txid))
		} else if version == V2 {
			txid = hex.EncodeToString([]byte(txid))
		}
		return fmt.Errorf("上链超时:%s", txid)
	}
}

var ErrDuplicateKey = errors.New("duplicate key")

var pool *sync.Map

func init() {
	pool = &sync.Map{}
}

// Put 根据Key值创建并返回同步信号通道.
func Put(key string) (<-chan error, error) {
	if _, has := pool.Load(key); has {
		return nil, ErrDuplicateKey
	}
	ch := make(chan error)
	pool.Store(key, ch)
	return ch, nil
}

// Signal 根据指定Key值发送同步信号.
func Signal(key string, value error) {
	v, has := pool.Load(key)
	if has {
		ch := v.(chan error)
		ch <- value
		close(ch)
	}
	pool.Delete(key)
}

func Delete(key string) {
	pool.Delete(key)
}

// SignalAPI 正确交易
func SignalAPI(key string) {
	v, has := pool.Load(key)
	if has {
		ch := v.(chan error)
		ch <- nil
		close(ch)
	}
	pool.Delete(key)
}
