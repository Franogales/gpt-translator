package localmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiUrl = "http://localhost:5000/generate"

type model string

const (
	ChatMessageRoleSystem = "system"
	ChatMessageRoleUser   = "user"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type LocalModelClient struct {
}

type ChatCompletionResponse struct {
	Response string `json:"response"`
}

func NewLocalModelClient() *LocalModelClient {
	return &LocalModelClient{}
}

func (gc LocalModelClient) sendChatRequest(jsonData []byte) (string, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}

	return string(body), nil
}

func (gc LocalModelClient) ChatCompletition(chatRequest ChatRequest) (ChatCompletionResponse, error) {
	jsonData, err := json.Marshal(chatRequest)
	if err != nil {
		return ChatCompletionResponse{}, fmt.Errorf("error marshalling JSON: %v", err)
	}

	response, err := gc.sendChatRequest(jsonData)
	if err != nil {
		return ChatCompletionResponse{}, err
	}

	var chatResponse ChatCompletionResponse
	err = json.Unmarshal([]byte(response), &chatResponse)
	if err != nil {
		return ChatCompletionResponse{}, fmt.Errorf("error unmarshalling response JSON: %v", err)
	}

	return chatResponse, nil
}
