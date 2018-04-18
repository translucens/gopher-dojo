package main

import "testing"

func TestUranai1(t *testing.T) {

	uranai1 := uranai(1)
	if uranai1 != "凶" {
		t.Error("Unexpected value of 1: " + uranai1)
	}
	t.Log("1: " + uranai1)
}
func TestUranai2(t *testing.T) {
	uranai2 := uranai(2)
	if uranai2 != "吉" {
		t.Error("Unexpected value of 2: " + uranai2)
	}
	t.Log("2: " + uranai2)
}
func TestUranai4(t *testing.T) {
	uranai4 := uranai(4)
	if uranai4 != "中吉" {
		t.Error("Unexpected value of 4: " + uranai4)
	}
	t.Log("4: " + uranai4)
}
func TestUranai6(t *testing.T) {
	uranai6 := uranai(6)
	if uranai6 != "大吉" {
		t.Error("Unexpected value of 6: " + uranai6)
	}
	t.Log("6: " + uranai6)
}

func TestUranai0(t *testing.T) {

	uranai0 := uranai(0)
	if uranai0 != "？？？" {
		t.Error("Unexpected value of 0: " + uranai0)
	}
	t.Log("0: " + uranai0)
}
