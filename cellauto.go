package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
)

func main () {
	// Initialize ('seed') the random number generator with the current
	// second. Computed from the unix "epoch".
	rand.Seed(time.Now().Unix())

	clearScreen()

	// Here we define the constants we want to use as parameters.
	const CellsNumber = 40
	const TurnsNumber = 15

	// Initialize an array of CellsNumber elements of type int.
	var state [CellsNumber]int

	// Assign a random integer in the interval [0,2).
	for c := 0; c < CellsNumber; c++ {
		state[c] = rand.Intn(2)
	}

	// At each turn.
	fmt.Println("turn", 0, state)
	time.Sleep(1*time.Second)
	for t := 1; t <= TurnsNumber; t++ {
		// Create a new array of cells to contain the state for this
		// turn. This is necessary because we need to refer to the
		// previous state to build it.
		var nextState [CellsNumber]int

		for c := 0; c < CellsNumber; c++ {
			previous := state[(c-1+CellsNumber)%CellsNumber]
			next := state[(c+1)%CellsNumber]

			if previous != next {
				nextState[c] = 1
			} else {
				nextState[c] = 0
			}
		}

		// Set the state of the automaton to the computed new state.
		state = nextState
		clearScreen()
		fmt.Println("turn", t, state)
		time.Sleep(1*time.Second)
	}

	clearScreen()
}

func clearScreen() {
	cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}
