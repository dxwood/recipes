package commands

import (
	"github.com/chriscoffee/recipes/recipes"
	"github.com/spf13/cobra"
)

var createRecipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Example: `  recipes create recipe`,

	Run: func(cmd *cobra.Command, args []string) {
		recipes.CreateRecipe()
	},
}

func init() {
	createCmd.AddCommand(createRecipeCmd)
}
