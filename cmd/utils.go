package cmd

import (
	"fmt"
	"os"

	"github.com/samueltorres/td/pkg/boards"
)

// LoadBoard tries to load the board and storage
func LoadBoard() (boards.BoardConfig, boards.BoardStorage) {
	cf := boards.NewBoardConfig()
	bfs, err := boards.NewBoardFileStorage(cf)
	if err != nil {
		CheckError(fmt.Errorf("could not load td configuration"))
	}

	return cf, bfs
}

// CheckError logs the error to the console and exits the program
func CheckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}
