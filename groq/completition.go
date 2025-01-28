package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiUrl = "https://api.groq.com/openai/v1/chat/completions"

type model string

const (
	Model_Llama3_8b_8192          model = "llama3-8b-8192"
	Model_Llama_3_1_70b_versatile model = "llama-3.1-70b-versatile"
	Model_Llama_3_3_70b_versatile model = "llama-3.3-70b-versatile"
)
const (
	ChatMessageRoleSystem = "system"
	ChatMessageRoleUser   = "user"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	PromptTime       float64 `json:"prompt_time"`
	CompletionTokens int     `json:"completion_tokens"`
	CompletionTime   float64 `json:"completion_time"`
	TotalTokens      int     `json:"total_tokens"`
	TotalTime        float64 `json:"total_time"`
}

type XGroq struct {
	ID string `json:"id"`
}

type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
	XGroq             XGroq    `json:"x_groq"`
}

type ChatRequest struct {
	Messages  []Message `json:"messages"`
	Model     model     `json:"model"`
	MaxTokens int       `json:"max_tokens"`
}

type GroqClient struct {
	apiKey string
}

func NewGroqClient(apiKey string) *GroqClient {
	return &GroqClient{
		apiKey: apiKey,
	}
}

func (gc GroqClient) sendChatRequest(jsonData []byte) (string, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.apiKey))
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

func (gc GroqClient) ChatCompletition(chatRequest ChatRequest) (ChatCompletionResponse, error) {
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

	if len(chatResponse.Choices) == 0 {
		return ChatCompletionResponse{}, fmt.Errorf("no choices in response")
	}

	return chatResponse, nil
}
