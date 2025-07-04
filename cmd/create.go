package cmd

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/prompts"
	"archi-ts-cli/internal/templates/project"
	"archi-ts-cli/internal/utils"

	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func runCreate(cmd *cobra.Command, args []string) error {

	// Banner
	utils.DisplayBanner()
	color.New(color.FgCyan, color.Bold).Println("🚀 ArchiTS CLI - Project Scaffolding")
	fmt.Println()

	/////////////////////////////////////////////////
	// Start collecting the result of each prompts //
	/////////////////////////////////////////////////

	// Prompt 1 - Prompt to ask the user for the project name
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("error with this project name: %w", err)
	}

	// Prompt 2 - Architecture Type
	architecture, err := prompts.PromptArchitecture()
	if err != nil {
		return fmt.Errorf("error when choosing an architecture: %w", err)
	}

	// Prompt 3 - Language Type (JS or TS)
	language, err := prompts.PromptLanguage()
	if err != nil {
		return fmt.Errorf("error when choosing a language: %w", err)
	}

	// Prompt 4 - Express Library ?
	express, err := prompts.PromptExpress()
	if err != nil {
		return fmt.Errorf("error when choosing express library: %w", err)
	}

	///////////////////////////////////////////////
	// End collecting the result of each prompts //
	///////////////////////////////////////////////

	// Step 3 - Create the project directory
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("error when trying to create the folder: %w", err)
	}

	// Step 4 - Select the project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("error during the change of directory: %w", err)
	}

	color.New(color.FgYellow).Printf("📁 Project created at: %s\n", projectPath)
	color.New(color.FgBlue).Printf("Architecture: %s\n", architecture)
	color.New(color.FgGreen).Printf("Language: %s\n", language)
	expressStatus := "No"
	if express {
		expressStatus = "Yes"
	}
	color.New(color.FgHiMagenta).Printf("Express Library: %s\n", expressStatus)
	fmt.Println()

	// Store the configuration in a struct and generate the project
	config := models.ProjectConfigBuilder{
		Name:         projectName,
		Architecture: models.Architecture(architecture),
		Language:     models.Language(language),
		Express:      express,
	}

	if err := project.GenerateProject(config); err != nil {
		return fmt.Errorf("error during the generation of the prject: %w", err)
	}

	// dev dependencies and main dependencies installation
	color.New(color.FgBlue).Println("📦 Dependencies installation...")
	if err := utils.InstallDependencies(string(language), express); err != nil {
		color.New(color.FgRed).Printf("⚠️  Error during dependencies installation: %v\n", err)
	}

	color.New(color.FgBlue).Println("📦 Installation de pnpm...")
	if err := utils.InstallPnpm(); err != nil {
		color.New(color.FgRed).Printf("⚠️  Error during pnpm installation: %v\n", err)
	}

	// Success message
	fmt.Println()
	color.New(color.FgGreen, color.Bold).Println("✅ Project structure created successfully!")
	fmt.Println()
	color.New(color.FgCyan).Printf("📂 Project name: %s\n", projectName)
	color.New(color.FgCyan).Printf("Architecture: %s\n", architecture)
	color.New(color.FgCyan).Printf("Language: %s\n", language)
	color.New(color.FgCyan).Printf("ExpressJS: %s\n", expressStatus)
	fmt.Println()
	color.New(color.FgYellow).Println("To start:")
	fmt.Printf("  cd %s\n", projectName)

	if language == "typescript" {
		fmt.Printf("  npm run dev (local development with nodemon)\n")
		fmt.Printf("  npm run build (before npm start)\n")
		fmt.Printf("  npm start (after npm build)\n")
	} else {
		fmt.Printf("  npm run dev (local development with nodemon)\n")
		fmt.Printf("  npm start \n")
	}

	fmt.Println()

	return nil
}
