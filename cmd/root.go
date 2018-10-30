package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var tdFolder string

var rootCmd = &cobra.Command{
	Use:   "td",
	Short: "todo list",
	Long:  `a cli todolist`,
}

// Execute executes the command line
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tdPath := path.Join(home, ".td")
	if _, err := os.Stat(tdPath); os.IsNotExist(err) {
		os.MkdirAll(tdPath, os.ModePerm)
	}

	viper.AddConfigPath(tdPath)
	viper.Set("tdPath", tdPath)
	viper.SetConfigName(".config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.WriteConfigAs(path.Join(tdPath, ".config.json"))
	}

	viper.SetDefault("boards_file", path.Join(tdPath, "boards.json"))
	viper.SetDefault("current_board", "default")
}
