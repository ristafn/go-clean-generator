package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Create a new project",
	Long:  "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		folders := []string{"bin", "domain", "repositories", "services", "delivery", "cmd", "configs", "utils"}
		files := []string{"main.go", "config.json", "README.md", ".gitignore"}

		// Project name
		name := cmd.Flag("name").Value.String()
		if name == "" {
			name = "go-clean-project"
		}

		// Project path
		path := cmd.Flag("path").Value.String()
		if path == "" {
			path = "."
		}

		// Create project folder
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", path, name), 0755); err != nil {
			panic(err)
		}

		for _, folder := range folders {
			err := os.MkdirAll(name+"/"+folder, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}

		for _, file := range files {
			_, err := os.Create(fmt.Sprintf("%s/%s", name, file))
			if err != nil {
				fmt.Println(err)
			}
		}

		// Write to .gitignore
		contentGitignore := "bin/*"
		if err := os.WriteFile(fmt.Sprintf("%s/.gitignore", name), []byte(contentGitignore), 0755); err != nil {
			fmt.Println(err)
		}

		// Write to main.go
		contentMain := `package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// TODO: Add your routes here

	// TODO: Add your repositories here

	// TODO: Add your services here

	// TODO: Add your delivery here

	app.Listen(":3000")
}
		`
		if err := os.WriteFile(fmt.Sprintf("%s/main.go", name), []byte(contentMain), 0755); err != nil {
			fmt.Println(err)
			panic(err)
		}

		// Write to config.json
		contentConfig := `{
	"app": {
		"status": "development",
		"port": "3000"
	},
	"database": {
		"host": "localhost",
		"port": "5432",
		"user": "postgres",
		"password": "postgres",
		"dbname": "postgres"
	},
	"minio": {
		"endpoint": "localhost:9000",
		"accessKeyID": "minio",
		"secretKeyID": "minio123",
		"useSSL": false
	},
	"redis": {
		"addr": "localhost:6379",
		"password": "",
		"db": 0
	}
}
		`
		if err := os.WriteFile(fmt.Sprintf("%s/config.json", name), []byte(contentConfig), 0755); err != nil {
			fmt.Println(err)
		}

		// Write to README.md
		if err := os.WriteFile(fmt.Sprintf("%s/README.md", name), []byte("# "+name), 0755); err != nil {
			fmt.Println(err)
		}

		// Change directory to the project
		if err := os.Chdir(name); err != nil {
			fmt.Println(err)
			panic(err)
		}

		// Running go mod init
		if err := exec.Command("go", "mod", "init", name).Run(); err != nil {
			fmt.Println(err)
			panic(err)
		}

		// Running go mod tidy
		if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println("Project created successfully")
	},
}

func Init() {
	// rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for the command")
	rootCmd.Flags().StringP("path", "P", "", "Path to the project")
	rootCmd.Flags().StringP("name", "N", "", "Name of the project")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
