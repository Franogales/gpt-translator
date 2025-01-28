package translate

import (
	"fmt"

	"github.com/Franogales/gpt-translator/groq"
)

type GroqChat struct {
	client *groq.GroqClient
}

func NewGroqChat(apiKey string) Translator {
	client := groq.NewGroqClient(apiKey)
	return &GroqChat{client}
}

func (gc GroqChat) TranslateES(text string) (string, error) {
	messages := []groq.Message{
		{
			Role:    groq.ChatMessageRoleSystem,
			Content: "you are a translation assistant. Your task is to translate any provided text into Spanish. Do not provide explanations or engage in conversations. Only respond with the accurate translation.",
		},
		{
			Role:    groq.ChatMessageRoleUser,
			Content: text,
		},
	}

	chatRequest := groq.ChatRequest{
		Messages:  messages,
		Model:     groq.Model_Llama_3_3_70b_versatile,
		MaxTokens: 10000,
	}

	response, err := gc.client.ChatCompletition(chatRequest)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}

func (gc GroqChat) TranslateEN(text string) (string, error) {
	messages := []groq.Message{
		{
			Role:    groq.ChatMessageRoleSystem,
			Content: "You are a translation assistant. Your task is to translate any provided text into English. Do not provide explanations or engage in conversations. Only respond with the accurate translation.",
		},
		{
			Role:    groq.ChatMessageRoleUser,
			Content: fmt.Sprintf("traduce el siguiente texto al inglés: %s", text),
		},
	}

	chatRequest := groq.ChatRequest{
		Messages:  messages,
		Model:     groq.Model_Llama_3_3_70b_versatile,
		MaxTokens: 10000,
	}

	response, err := gc.client.ChatCompletition(chatRequest)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}
