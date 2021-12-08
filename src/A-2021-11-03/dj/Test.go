package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 	1       0001
	// 	2       0010
	// 	3       0011
	// 	4       0100
	// 	5       0101
	// 	6       0110
	// 	7       0111
	// 	8       1000
	// 	9       1001
	// 	10      1010

	// 	假设是第n瓶酒有毒   (活着：0    死了：1)
	// 	第一个人    1        （8、9、10）
	// 	第二个人    2        （4、5、6、7）
	// 	第三个人    3        （2、3、6、7、10）
	// 	第四个人    4        （1、3、5、7、9）

	for i := 1; i < 11; i++ {
		fmt.Println(convertToBin(i))
	}

}

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	sLen := len(s)
	if sLen < 4 {
		for i := 0; i < 4-sLen; i++ {
			s = "0" + s
		}
	}
	return s
}
