/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Colors for the CLI
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

// Debug variable
var debug bool = false

// GPT variables
var defaultModel = "gpt-3.5-turbo"
var prompt string
var maxTokens int
var temperature float64
type ChatGPTResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

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
		if debug == true {
			fmt.Println(yellow("Debug mode enabled"))
			fmt.Printf("Prompt: %s\n", prompt)
			fmt.Printf("Model: %s\n", defaultModel)
			fmt.Printf("Max tokens: %d\n", maxTokens)
			return
		}

		generateText(prompt, defaultModel, maxTokens, temperature)
	},
}

func generateText(prompt string, model string, maxTokens int, temperature float64) {
	
	client := &http.Client{}

	url := "https://api.openai.com/v1/chat/completions"

	inputMessage := fmt.Sprintf(`{"role": "user", "content": "%s"}`, prompt)
	payload := strings.NewReader(
		fmt.Sprintf(`{
			"model": "%s",
			"messages": [{"role": "system", "content": "You are a helpful assistant."}, %s],
			"max_tokens": %d,
			"temperature": %f
		}`, model, inputMessage, maxTokens, temperature))

	if debug == true {
		fmt.Println("Request body:")
		fmt.Println(payload)
	}
	
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(red("Error: Failed to send request to ChatGPT API"))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf(red("Error: ChatGPT API returned a non-200 status code: %d\n"), resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(red("Error: Failed to read response from ChatGPT API"))
		return
	}

	if debug == true {
		responseString := string(body)
		fmt.Println("Response body:")
		fmt.Println(responseString)
		return
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(body, &chatGPTResponse)
	if err != nil {
		fmt.Println(red("Error: Failed to parse JSON response"))
		return
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
	generateCmd.Flags().Float64P("temperature", "t", 0.5, "The sampling temperature")
	// TODO: fix debug flag
	generateCmd.Flags().BoolP("debug", "d", true, "Enable debug mode")
	generateCmd.MarkFlagRequired("prompt")
	rootCmd.AddCommand(generateCmd)
}
