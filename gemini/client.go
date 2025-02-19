package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
}

func NewGeminiClient(apiKey string) GeminiClient {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	return GeminiClient{client}
}

func (gc GeminiClient) Close() {
	gc.client.Close()
}

func (gc *GeminiClient) ChatCompletition(systemPrompt, userPrompt string) (string, error) {
	ctx := context.Background()

	// [START text_gen_text_only_prompt]
	model := gc.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = genai.NewUserContent(genai.Text(systemPrompt))
	resp, err := model.GenerateContent(ctx, genai.Text(userPrompt))
	if err != nil {
		log.Fatal(err)
	}

	return extractResponse(resp)
}

func extractResponse(resp *genai.GenerateContentResponse) (string, error) {
	var responseText string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {

				responseText = fmt.Sprintf("%s", part)
			}
		}
	}

	return responseText, nil
}
