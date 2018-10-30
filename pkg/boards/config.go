package boards

import (
	"github.com/spf13/viper"
)

// BoardConfig exposes the configuration of a board
type BoardConfig struct {
	BoardsFile   string
	CurrentBoard string
}

// NewBoardConfig creates a new configuration for todo boards
func NewBoardConfig() BoardConfig {
	bf := viper.GetString("boards_file")
	cb := viper.GetString("current_board")

	return BoardConfig{
		BoardsFile:   bf,
		CurrentBoard: cb,
	}
}

// SaveConfig stores in the configuration file the current configuration
func (bc *BoardConfig) SaveConfig() error {
	viper.Set("boards_file", bc.BoardsFile)
	viper.Set("current_board", bc.CurrentBoard)

	return viper.WriteConfig()
}
