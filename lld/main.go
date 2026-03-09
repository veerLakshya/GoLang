package main

import (
	"fmt"
	ttt "github.com/veerlakshya/my-go-all/lld/tic-tac-toe"
)

func main() {
	system := ttt.GetInstance()

	alice, _ := ttt.NewPlayer("Alice", ttt.SymbolX)
	bob, _ := ttt.NewPlayer("Bob", ttt.SymbolO)

	fmt.Println("=========Game=========")
	system.CreateGame(alice, bob)

	system.MakeMove(alice, 0, 0)
	system.MakeMove(bob, 1, 0)
	system.MakeMove(alice, 0, 1)
	system.MakeMove(bob, 1, 1)
	system.MakeMove(alice, 0, 2)

	status, _ := system.GameStatus()
	fmt.Printf("Game 1 Result: %s\n", status)

}
