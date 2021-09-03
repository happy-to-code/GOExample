package main

import (
	"testing"
)

func Test_convertToObj(t *testing.T) {
	s := "6665535,864310308f36426783d84afb1bfbae98,0,registration"
	obj := convertToObj(s)
	index := obj.Index
	t.Logf("%+v", obj)
	t.Log(index == 6665535)
}
