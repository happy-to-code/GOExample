package main

import (
	"encoding/json"
	"testing"
)

var pensionList []Pension

func init() {
	json.Unmarshal([]byte(s), &pensionList)
}

func Benchmark_countContinueMonths1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count := countContinueMonths1(pensionList)
		if count != 120 {
			b.Fail()
		}
	}
}

func Benchmark_countContinueMonths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count := countContinueMonths(pensionList)
		if count != 120 {
			b.Fail()
		}
	}
}

func Benchmark_countContinueMonths2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count := countContinueMonths2(pensionList)
		if count != 120 {
			b.Fail()
		}
	}
}
