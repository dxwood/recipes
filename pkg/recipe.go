package recipe

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

	}
	},
}
