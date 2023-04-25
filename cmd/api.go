package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

func apiRequest(inputMessage string, payload io.Reader) (ChatGPTResponse, error) {
	client := &http.Client{}
	url := "https://api.openai.com/v1/chat/completions"

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(red("Error: Failed to send request to ChatGPT API"))
		return ChatGPTResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf(red("Error: ChatGPT API return ed a non-200 status code: %d\n"), resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(red("Error: Failed to read response from ChatGPT API"))
		return ChatGPTResponse{}, err
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(body, &chatGPTResponse)
	if err != nil {
		fmt.Println(red("Error: Failed to parse JSON response"))
		return ChatGPTResponse{}, err
	}

	return chatGPTResponse, nil

}