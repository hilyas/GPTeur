/*
Copyright © 2023 Ilyas Hamdi <ilyas.hamdi@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var apiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpteur",
	Short: "A CLI to interact with the ChatGPT API",
	Long: `GPTeur is a command line tool to interact 
	with OpenAI's ChatGPT API. It allows you to
	input a prompt and get a response from the
	ChatGPT model.`,
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
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&apiKey, "apikey", "", "API key for ChatGPT authentication")
}


