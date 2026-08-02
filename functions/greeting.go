package functions

import (
	"context"
	"fmt"
)

func Greeting(ctx context.Context, name string) (string, error) {
	if len(name) < 7 {
		return fmt.Sprintf("Hello %s! Your name is to short\n", name), nil
	}
	return fmt.Sprintf("Hi %s", name), nil
}
