package translate

import (
	"fmt"

	"github.com/Franogales/gpt-translator/localmodel"
)

type LocalChat struct {
	client *localmodel.LocalModelClient
}

func NewLocalChat() Translator {
	client := localmodel.NewLocalModelClient()
	return &LocalChat{client}
}

func (gc LocalChat) TranslateES(text string) (string, error) {
	messages := []localmodel.Message{
		{
			Role:    localmodel.ChatMessageRoleSystem,
			Content: PromptSystemEnglishToSpanish,
		},
		{
			Role:    localmodel.ChatMessageRoleUser,
			Content: fmt.Sprintf(PromptUserEnglishToSpanish, text),
		},
	}

	chatRequest := localmodel.ChatRequest{
		Messages:  messages,
		MaxTokens: 10000,
	}

	response, err := gc.client.ChatCompletition(chatRequest)
	if err != nil {
		return "", err
	}
	return response.Response, nil
}

func (gc LocalChat) TranslateEN(text string) (string, error) {
	messages := []localmodel.Message{
		{
			Role:    localmodel.ChatMessageRoleSystem,
			Content: PromptUserSpanishToEnglish,
		},
		{
			Role:    localmodel.ChatMessageRoleUser,
			Content: fmt.Sprintf(PromptUserSpanishToEnglish, text),
		},
	}

	chatRequest := localmodel.ChatRequest{
		Messages:  messages,
		MaxTokens: 10000,
	}

	response, err := gc.client.ChatCompletition(chatRequest)
	if err != nil {
		return "", err
	}
	return response.Response, nil
}
