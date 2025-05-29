package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

/*
*	PromptProjectName to ask the name of the project.
 */
func PromptProjectName() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Project Name",
		Default: "archi-project",
		Validate: func(input string) error {
			if len(input) < 1 {
				return fmt.Errorf("le nom du projet ne peut pas être vide")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}
