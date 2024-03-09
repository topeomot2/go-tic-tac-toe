package helper

import (
	"fmt"
	"testing"
)

func setupTest(t *testing.T) ([3][3]string, func(t *testing.T)) {
	fmt.Println("set up")
	store := [3][3]string{
		{"X", "", ""},
		{"", "O", ""},
		{"", "", "X"},
	}

	return store, func(t *testing.T) {
		fmt.Println("clean up")
	}
}

func TestDrawGameBoard(t *testing.T) {
	store, tearDownTest := setupTest(t)
	defer tearDownTest(t)

	expectedBoard := fmt.Sprintf("%v  | %v  |    %v    \n--------------------------\n%v  |   %v    |  %v \n--------------------------\n  %v    | %v  |  %v\n",
		"(0,2)", "(1,2)", store[2][2], "(0,1)", store[1][1], "(2,1)", store[0][0], "(1,0)", "(2,0)")
	board := DrawGameBoard(store)

	if board != expectedBoard {
		t.Fatal(`Board not correctly drawn`)
	}
}

func TestValidateInput(t *testing.T) {
	store, tearDownTest := setupTest(t)
	defer tearDownTest(t)

	x, y, err := ValidateInput("0", store)
	if err.Error() != "Wrong input" {
		t.Error("Input of 0 not throwing error 'Wrong input'")
	}

	_, _, err = ValidateInput("1,1", store)
	if err.Error() != "Location not empty" {
		t.Error("Input of 1,1 not throwing error 'Location not empty'")
	}

	x, y, err = ValidateInput("2,1", store)
	if err != nil {
		t.Error("Input of 2,1 throwing error")
	}

	if x != 2 || y != 1 {
		t.Error("Input of 2,1 return wrong values")
	}
}

func TestFullPathExists(t *testing.T) {
	store, tearDownTest := setupTest(t)
	defer tearDownTest(t)

	if FullPathExists(store) {
		t.Error("Winning Path should not exist")
	}

	store[1][1] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = "O"
	store[1][1] = "O"
	store[2][2] = "X"
	store[1][2] = "X"
	store[0][2] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = "O"
	store[1][1] = "O"
	store[2][2] = "X"
	store[1][2] = ""
	store[0][2] = ""
	store[2][1] = "X"
	store[2][0] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = "O"
	store[1][1] = ""
	store[2][2] = ""
	store[1][0] = "O"
	store[2][0] = "O"
	store[2][1] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = "O"
	store[1][0] = ""
	store[2][0] = "X"
	store[2][1] = "X"
	store[0][1] = "O"
	store[0][2] = "O"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = "O"
	store[1][0] = "X"
	store[2][0] = ""
	store[2][1] = ""
	store[0][1] = "O"
	store[0][2] = ""
	store[1][1] = "X"
	store[1][2] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}

	store[0][0] = ""
	store[1][0] = ""
	store[2][0] = "X"
	store[2][1] = "O"
	store[0][1] = "O"
	store[0][2] = ""
	store[1][1] = "O"
	store[1][2] = "X"

	if !FullPathExists(store) {
		t.Error("Winning Path should exist")
	}
}
