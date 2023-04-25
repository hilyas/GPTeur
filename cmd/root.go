package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// TODO: add debug flag to root command

var apiKey string

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

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


