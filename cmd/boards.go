package cmd

import (
	"fmt"

	"github.com/samueltorres/td/pkg/boards"
	"github.com/samueltorres/td/pkg/printer"
	"github.com/spf13/cobra"
)

var boardCmd = &cobra.Command{
	Use:                   "board SUBCOMMAND",
	DisableFlagsInUseLine: true,
	Short:                 "Manages and lists todo boards",
}

var boardListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all boards",
	Run: func(cmd *cobra.Command, args []string) {
		_, bs := LoadBoard()
		boards, err := bs.GetBoards()
		CheckError(err)

		printer.PrintBoards(boards)
	},
}

var boardAddCmd = &cobra.Command{
	Use:   "add <board>",
	Short: "Adds a board",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, bs := LoadBoard()
		err := bs.AddBoard(args[0])
		CheckError(err)

		printer.PrintMessage(fmt.Sprintf("%s board was added.", args[0]))
	},
}

var boardSetCmd = &cobra.Command{
	Use:   "set [board]",
	Short: "Sets current board",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cf, bs := LoadBoard()
		boards, err := bs.GetBoards()
		CheckError(err)

		found := false
		for _, v := range boards {
			if v == args[0] {
				found = true
				break
			}
		}

		if !found {
			CheckError(fmt.Errorf("board %v does not exist", args[0]))
		}

		cf.CurrentBoard = args[0]
		err = cf.SaveConfig()
		CheckError(err)

		printer.PrintMessage(fmt.Sprintf("%s was set as the current board.", args[0]))
	},
}

var boardGetCurrentCmd = &cobra.Command{
	Use:   "get current",
	Short: "Gets current board",
	Run: func(cmd *cobra.Command, args []string) {
		cf := boards.NewBoardConfig()
		printer.PrintMessage(cf.CurrentBoard)
	},
}

func init() {
	boardCmd.AddCommand(boardListCmd)
	boardCmd.AddCommand(boardAddCmd)
	boardCmd.AddCommand(boardSetCmd)
	boardCmd.AddCommand(boardGetCurrentCmd)

	rootCmd.AddCommand(boardCmd)
}
