package main

import "testing"

func TestBasicDraw(t *testing.T) {
	g := new(Game)
	expected := `-------------
|   |   |   |
-------------
|   |   |   |
-------------
|   |   |   |
-------------`

	if g.String() != expected {
		t.Errorf("Board output was incorrect: \n%s, want: \n%s.", g.String(), expected)
	}
}

func TestDraw(t *testing.T) {
	g := Game{[9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}}
	expected := `-------------
| 1 | 2 | 3 |
-------------
| 4 | 5 | 6 |
-------------
| 7 | 8 | 9 |
-------------`

	if g.String() != expected {
		t.Errorf("Board output was incorrect: \n%s, want: \n%s.", g.String(), expected)
	}
}

func TestMoveValidation(t *testing.T) {
	g := new(Game)

	if g.Move("0") {
		t.Errorf("Validation should have blocked move: 0")
	}
	if g.Move("10") {
		t.Errorf("Validation should have blocked move: 10")
	}
	if g.Move("abc") {
		t.Errorf("Validation should have blocked move: abc")
	}
	if g.Move("") {
		t.Errorf("Validation should have blocked move: abc")
	}
}

func TestMoveOnOnlyEmptySpace(t *testing.T) {
	g := Game{[9]string{"o"}}

	if g.Move("1") || g.Board[0] != "o" {
		t.Errorf("Able to make move when computer occupied that space")
	}
}

func TestMoveWorks(t *testing.T) {
	g := new(Game)

	if !g.Move("5") || g.Board[4] != "x" {
		t.Errorf("Unable to make move when space is free")
	}
}

func TestComputerMove(t *testing.T) {
	g := Game{[9]string{"", "x", "o", "o", "o", "x", "x", "o", "x"}}

	g.ComputerMove()

	if g.Board[0] != "o" {
		t.Errorf("Unable to make computer move when space is free")
	}
}

var resultsTests = []struct {
	expected1 bool
	expected2 string
	board     [9]string
}{
	{false, "", [9]string{"o"}},
	{true, "o", [9]string{0: "o", 1: "o", 2: "o"}},
	{true, "o", [9]string{3: "o", 4: "o", 5: "o"}},
	{true, "o", [9]string{6: "o", 7: "o", 8: "o"}},
	{true, "o", [9]string{0: "o", 3: "o", 6: "o"}},
	{true, "o", [9]string{1: "o", 4: "o", 7: "o"}},
	{true, "o", [9]string{2: "o", 5: "o", 8: "o"}},
	{true, "o", [9]string{0: "o", 4: "o", 8: "o"}},
	{true, "o", [9]string{6: "o", 4: "o", 2: "o"}},
	{true, "", [9]string{"x", "x", "o", "o", "o", "x", "x", "o", "x"}},
}

func TestCheckResults(t *testing.T) {
	for tid, tt := range resultsTests {
		g := Game{tt.board}
		fin, winner := g.CheckResults()
		if fin != tt.expected1 {
			t.Errorf("failed test %d expected %t, got %t", tid, fin, tt.expected1)
		}
		if winner != tt.expected2 {
			t.Errorf("failed test %d expected %s, got %s", tid, winner, tt.expected2)
		}
	}
}
