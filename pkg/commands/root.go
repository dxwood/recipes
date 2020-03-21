package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "recipes",
	Short: "Add recipes to list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultPrompt()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(myBanner))

	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "enable debug mode")

}
