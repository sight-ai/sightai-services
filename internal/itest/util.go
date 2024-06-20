package itest

import (
	"context"
	openapi "github.com/capybaralabs-xyz/sightai-services/internal/itest/client"
)

func GetContextWithUserToken(ctx context.Context, jwt string) context.Context {
	newCtx := context.WithValue(ctx, openapi.ContextAPIKeys, map[string]openapi.APIKey{
		"userJwtToken": {
			jwt,
			"",
		},
	})

	return newCtx
}
