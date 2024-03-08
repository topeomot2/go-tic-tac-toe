package main

import (
	"bufio"
	"fmt"
	"os"

	"topeomot.com/tictactoe/helper"
)

func main() {
	fmt.Println("Hello, Welcome to Tic-Tac-Toe!")
	var gameStore [3][3]string
	scanner := bufio.NewScanner(os.Stdin)

	signs := [2]string{"X", "O"}
	turn := 1
	steps := 0

	for {
		if steps == 9 {
			fmt.Println("No Winner")
			break
		}

		helper.DrawGameBoard(gameStore)
		fmt.Println(`Player ` + signs[turn-1])
		fmt.Println("Select Empty Board Position from 0,0 to 2,2: ")
		scanner.Scan()
		input := scanner.Text()
		x, y, err := helper.ValidateInput(input, gameStore)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		gameStore[x][y] = signs[turn-1]
		if helper.FullPathExists(gameStore) {
			fmt.Println(`Player ` + signs[turn-1] + ` wins`)
			helper.DrawGameBoard(gameStore)
			break
		}

		if turn == 2 {
			turn = 1
		} else {
			turn = 2
		}
		steps++

	}
	os.Exit(0)
}
