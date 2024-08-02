package translate

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type GPTChat struct {
	client *openai.Client
}

func NewGPTChat(token string) Translator {
	client := openai.NewClient(token)
	return &GPTChat{client}
}

func (gpt GPTChat) TranslateEN(text string) (string, error) {
	resp, err := gpt.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens: 200,
			Model:     openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Responde únicamente con la traducción en inglés",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (gpt GPTChat) TranslateES(text string) (string, error) {
	resp, err := gpt.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens: 200,
			Model:     openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Translate to spanish: " + text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
