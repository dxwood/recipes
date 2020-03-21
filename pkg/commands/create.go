package commands

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [item]",
	Short: "Create a new event, organizer, program, speaker, sponsor, or talk",
	Long:  `Use this to create events, sponsors, talks, etc`,
}

func init() {
	RootCmd.AddCommand(createCmd)
}
