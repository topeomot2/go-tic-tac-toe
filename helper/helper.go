package helper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func DrawGameBoard(store [3][3]string) string {
	var displayBoard [3][3]string

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if len(store[i][j]) != 0 {
				displayBoard[i][j] = fmt.Sprint("  ", store[i][j], "  ")
			} else {
				displayBoard[i][j] = fmt.Sprint(`(`, i, `,`, j, `)`)
			}
		}
	}

	message := fmt.Sprintf("%v  | %v  |  %v  \n--------------------------\n%v  | %v  |  %v \n--------------------------\n%v  | %v  |  %v\n",
		displayBoard[0][2], displayBoard[1][2], displayBoard[2][2], displayBoard[0][1], displayBoard[1][1], displayBoard[2][1], displayBoard[0][0], displayBoard[1][0], displayBoard[2][0])

	fmt.Println(message)
	return message
}

func ValidateInput(input string, store [3][3]string) (int, int, error) {
	re := regexp.MustCompile(`^([0-2]{1},[0-2]{1})+$`)
	match := re.MatchString(input)
	if !match {
		return 0, 0, errors.New("Wrong input")
	}

	parts := strings.Split(input, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("Wrong x coordinate")
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, errors.New("Wrong y coordinate")
	}

	if len(store[x][y]) > 0 {
		return 0, 0, errors.New("Location not empty")
	}

	return x, y, nil
}

func FullPathExists(store [3][3]string) bool {
	return (len(store[0][0]) > 0 && (store[0][0] == store[1][0] && store[1][0] == store[2][0])) ||
		(len(store[0][0]) > 0 && (store[0][0] == store[0][1] && store[0][1] == store[0][2])) ||
		(len(store[0][0]) > 0 && (store[0][0] == store[1][1] && store[1][1] == store[2][2])) ||
		(len(store[2][2]) > 0 && (store[2][2] == store[1][2] && store[1][2] == store[0][2])) ||
		(len(store[2][2]) > 0 && (store[2][2] == store[2][1] && store[2][1] == store[2][0])) ||
		(len(store[1][1]) > 0 && (store[1][1] == store[1][2] && store[1][1] == store[1][0])) ||
		(len(store[1][1]) > 0 && (store[1][1] == store[0][1] && store[1][1] == store[2][1]))
}
