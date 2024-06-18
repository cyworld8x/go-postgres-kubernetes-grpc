package prompt

import "context"

type UseCase interface {
	SinglePrompt(ctx context.Context, prompt string) (string, error)
}
