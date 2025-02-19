package translate

import (
	"fmt"

	"github.com/Franogales/gpt-translator/gemini"
)

type GeminiChat struct {
	client gemini.GeminiClient
}

func NewGeminiChat(apikey string) Translator {
	client := gemini.NewGeminiClient(apikey)
	return &GeminiChat{client}
}

func (gc *GeminiChat) TranslateES(text string) (string, error) {
	response, err := gc.client.ChatCompletition(PromptSystemEnglishToSpanish, fmt.Sprintf(PromptUserEnglishToSpanish, text))
	if err != nil {
		return "", err
	}
	return response, nil
}

func (gc *GeminiChat) TranslateEN(text string) (string, error) {
	response, err := gc.client.ChatCompletition(PromptUserSpanishToEnglish, fmt.Sprintf(PromptUserSpanishToEnglish, text))
	if err != nil {
		return "", err
	}
	return response, nil
}
