package functions

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Greeting(ctx context.Context, name string) (string, error) {
	log.Ctx(ctx).Info().Msg("generating greeting")

	result, err := validateName(ctx, name)
	if err != nil {
		return "", err
	}

	log.Ctx(ctx).Info().Str("name", name).Msg("greeting ready")
	return result, nil
}

func validateName(ctx context.Context, name string) (string, error) {
	log.Ctx(ctx).Info().Msg("validating name")

	if len(name) < 7 {
		return fmt.Sprintf("Hello %s! Your name is to short\n", name), nil
	}
	return fmt.Sprintf("Hi %s", name), nil
}
