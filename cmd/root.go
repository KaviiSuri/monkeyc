/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KaviiSuri/monkeyc/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const words = "possible since through follow one long can way know want part out she lead could call so general here we should find how late end where look how both year present end"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monkeyc",
	Short: "A monkey type clone, write in your terminal! Monkey C, Monky Type!",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// f, _ := tea.LogToFile("debug.log", "debug")
		// fmt.Fprintln(f, words)
		ui := ui.New(words)
		program := tea.NewProgram(ui)

		if _, err := program.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.monkeyc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
