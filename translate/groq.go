package translate

import (
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
			Content: "Responde únicamente con la traducción en español",
		},
		{
			Role:    groq.ChatMessageRoleUser,
			Content: text,
		},
	}

	chatRequest := groq.ChatRequest{
		Messages:  messages,
		Model:     groq.Model_Llama_3_1_70b_versatile,
		MaxTokens: 100,
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
			Content: "Responde únicamente con la traducción en inglés",
		},
		{
			Role:    groq.ChatMessageRoleUser,
			Content: text,
		},
	}

	chatRequest := groq.ChatRequest{
		Messages:  messages,
		Model:     groq.Model_Llama_3_1_70b_versatile,
		MaxTokens: 100,
	}

	response, err := gc.client.ChatCompletition(chatRequest)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}
