package nitrado

import (
	"context"
	"fmt"
)

func ValidateToken(ctx context.Context, token string) error {
	fmt.Println("token:", token)
	return nil
}
