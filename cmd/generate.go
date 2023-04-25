package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a response from the ChatGPT model",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if apiKey == "" {
			fmt.Println(red("Error: API key is required"))
			return
		}
		if prompt == "" {
			fmt.Println(red("Error: A prompt is required"))
			return
		}

		generateText(defaultModel, prompt, maxTokens, temperature)
	},
}

func generateText(model string, prompt string, maxTokens int, temperature float64) {
	
	inputMessage := fmt.Sprintf(`{"role": "user", "content": "%s"}`, prompt)
	payload := strings.NewReader(
		fmt.Sprintf(`{
			"model": "%s",
			"messages": [{"role": "system", "content": "You are a helpful assistant."}, %s],
			"max_tokens": %d,
			"temperature": %f
		}`, model, inputMessage, maxTokens, temperature))

	chatGPTResponse, err := apiRequest(inputMessage, payload)
	if err != nil {
		fmt.Println(red("Error: Failed to parse JSON response"))
	}
	if len(chatGPTResponse.Choices) > 0 {
		assistantMessage := chatGPTResponse.Choices[0].Message
		fmt.Println(yellow("Generated text:"))
		fmt.Println(green(assistantMessage.Content))
	} else {
		fmt.Println(red("Error: No generated text found"))
	}

}

func init() {
	generateCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "The input prompt for text generation")
	generateCmd.Flags().IntVarP(&maxTokens, "max-tokens", "m", 250, "The maximum number of tokens to generate")
	generateCmd.Flags().Float64P("temperature", "t", 0.9, "The temperature of the model")
	generateCmd.MarkFlagRequired("prompt")
	rootCmd.AddCommand(generateCmd)
}



