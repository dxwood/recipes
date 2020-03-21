package commands

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/chriscoffee/recipes/recipes"
)

func defaultPrompt() error {
	var selection string

	for selection != "Quit the application" {
		prompt := &survey.Select{
			Message: "Select an action:",
			Options: []string{
				"Create a new recipe",
				"Quit the application",
			},
		}
		err := survey.AskOne(prompt, &selection, nil)
		if err != nil {
			break
		}
	}
	switch selection {
	case "Create a new recipe":
		recipes.CreateRecipe()
	}
}
