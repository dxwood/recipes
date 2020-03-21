package recipes

import (
	"github.com/AlecAivazis/survey/v2"
	"fmt"
)

var createRecipe = []*.survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What's the name of the recipe?",
			Help: "This is the full name of the recipe",
		},
		Validate: survey.Required,
	},
	{
		Name: "description",
		Prompt: &survey.Editor{
			Message: "Enter a description for the recipe [Press enter to launch the editor]"
		},
	},
	{
		Name: "method",
		Prompt: &survey.Editor{
			Message: "Enter a method for preparing the recipe [Press enter to launch the editor]"
		},
	},
}

func CreateRecipe() error{

	answers := struct{
		Name: string,
		Description: string,
		Method: string,
	}{}

	if err = survey.Ask(createRecipe, &answers); err != nil; {
		fmt.Println(err)
		return
	}

	return
}
