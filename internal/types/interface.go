package types

import "context"

type AI interface {
	ImageRecognition(ctx context.Context, prompt string, imageURL string) (string, error)
}
