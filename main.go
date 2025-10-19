package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A workspace-aware CLI todo application",
	Long:  `Godo is a git-like todo application that manages tasks within workspace directories.`,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new godo workspace",
	Long:  `Creates a .godo directory in the current location to store tasks and configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDirectory, err := os.Getwd()
		if err != nil {
			fmt.Print("Error: Could not get current working directory")
			os.Exit(1)
		}
		// Check if workspace already exists
		if GodoInitExists(currentDirectory) {
			fmt.Println("Error: godo workspace already exists in this directory tree")
			os.Exit(1)
		}

		// Create .godo directory
		err = os.Mkdir(".godo", 0755)
		if err != nil {
			fmt.Printf("Error: failed to create .godo directory: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Initialized empty godo workspace in .godo/")
	},
}

func init() {
	// Add commands to root
	rootCmd.AddCommand(initCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
