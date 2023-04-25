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

	"github.com/spf13/cobra"
)

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



var prompt string
var defaultModel = "gpt-3.5-turbo"

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a response from the ChatGPT model",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		if apiKey == "" {
			fmt.Println("Error: API key is required")
			return
		}
		if prompt == "" {
			fmt.Println("Error: A prompt is required")
			return
		}

		generateText(prompt, defaultModel)
	},
}

func generateText(prompt string, model string) {
	client := &http.Client{}

	url := "https://api.openai.com/v1/chat/completions"

	inputMessage := fmt.Sprintf(`{"role": "user", "content": "%s"}`, prompt)
	payload := strings.NewReader(
		fmt.Sprintf(`{
			"model": "%s",
			"messages": [{"role": "system", "content": "You are a helpful assistant."}, %s],
			"max_tokens": 100,
			"temperature": 0.9
		}`, model, inputMessage))

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: Failed to send request to ChatGPT API")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: ChatGPT API returned a non-200 status code: %d\n", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Failed to read response from ChatGPT API")
		return
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(body, &chatGPTResponse)
	if err != nil {
		fmt.Println("Error: Failed to parse JSON response")
		return
	}

	if len(chatGPTResponse.Choices) > 0 {
		assistantMessage := chatGPTResponse.Choices[0].Message
		fmt.Println("Generated text:")
		fmt.Println(assistantMessage.Content)
	} else {
		fmt.Println("Error: No generated text found")
	}
	
	
}


func init() {
	generateCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "The input prompt for text generation")
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
