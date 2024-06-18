package prompt

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type service struct {
	llm *openai.LLM
}

func NewPromptHandler(llm *openai.LLM) UseCase {
	return &service{
		llm: llm,
	}
}

func (s *service) SinglePrompt(ctx context.Context, prompt string) (string, error) {
	log.Info().Msgf("Prompt: %s", prompt)
	return llms.GenerateFromSinglePrompt(ctx, s.llm, prompt)
}
