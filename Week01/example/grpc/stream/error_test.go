package main

import "testing"

func TestERr1(t *testing.T) {
	t.Error("err5")
}

func TestERr2(t *testing.T) {
	t.Error("err2")
}
func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {

		t.Errorf("expecting bar , got %s", result)
	}
}
