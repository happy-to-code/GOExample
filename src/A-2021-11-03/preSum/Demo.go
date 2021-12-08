package main

func main() {
	var ints = []int{1, 2, 5, -1, 10, 3}
	preSum(ints)

}

func preSum(ints []int) {
	newIntList := make([]int, len(ints)+1)
	for i := 1; i < len(newIntList); i++ {
		newIntList[i] = newIntList[i-1] + ints[i-1]
	}
	for _, v := range newIntList {
		println(v)
	}
}
