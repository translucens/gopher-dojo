package main

import (
	"os"
	"testing"
)

const TESTFILE string = "5lines.txt"

func TestHeadSingleLine(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetHead(fd, 1)

	if len(lines) != 1 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
	}
}

func TestHead3Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetHead(fd, 3)

	if len(lines) != 3 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[1] != "Next line is empty" {
		t.Error("Does not match 2nd line content")
		t.Log("Content: " + lines[1])
	}
	if lines[2] != "" {
		t.Error("Does not match 3rd line content")
		t.Log("Content: " + lines[2])
	}
}

func TestHead5Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetHead(fd, 5)

	if len(lines) != 5 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[4] != "5" {
		t.Error("Does not match 5th line content")
		t.Log("Content: " + lines[4])
	}
}

func TestHead6Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetHead(fd, 6)

	if len(lines) != 5 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[4] != "5" {
		t.Error("Does not match 5th line content")
		t.Log("Content: " + lines[4])
	}
}

func TestTailSingleLine(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetTail(fd, 1)

	if len(lines) != 1 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "5" {
		t.Error("Does not match last line content")
		t.Log("Content: " + lines[0])
	}
}

func TestTail3Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetTail(fd, 3)

	if len(lines) != 3 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[1] != "四" {
		t.Error("Does not match 2nd line content")
		t.Log("Content: " + lines[1])
	}
	if lines[2] != "5" {
		t.Error("Does not match 3rd line content")
		t.Log("Content: " + lines[2])
	}
}

func TestTail5Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetTail(fd, 5)

	if len(lines) != 5 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[4] != "5" {
		t.Error("Does not match 5th line content")
		t.Log("Content: " + lines[4])
	}
}

func TestTail6Line(t *testing.T) {
	fd, err := os.Open("./testmaterial/" + TESTFILE)

	if err != nil {
		t.Fatal("Cannot open file " + TESTFILE)
	}

	defer fd.Close()
	var lines []string
	lines, err = GetTail(fd, 6)

	if len(lines) != 5 {
		t.Error("Does not match line number")
		t.Logf("lines: %d", len(lines))
	}

	if lines[0] != "1" {
		t.Error("Does not match 1st line content")
		t.Log("Content: " + lines[0])
	}
	if lines[4] != "5" {
		t.Error("Does not match 5th line content")
		t.Log("Content: " + lines[4])
	}
}
