package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr []string
	for i := 0; i < 10000000; i++ {
		rand := rand.Intn(1000000)
		arr = append(arr, fmt.Sprintf("%d", rand))
	}

	t1 := time.Now()
	removeDuplication(arr)
	t2 := time.Now()
	fmt.Println("1--->", t2.Sub(t1))

	removeRepeat(arr)
	t3 := time.Now()
	fmt.Println("2--->", t3.Sub(t2))

}
func removeDuplication(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func removeRepeat(slc []string) (result []string) {
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
