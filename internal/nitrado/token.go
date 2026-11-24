package nitrado

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
)

type tokenResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token struct {
			Scopes []string `json:"scopes"`
		} `json:"token"`
	} `json:"data"`
}

func ValidateToken(ctx context.Context, token string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://oauth.nitrado.net/token", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid token")
	}

	var payload tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return err
	}

	if payload.Status != "success" || !slices.Contains(payload.Data.Token.Scopes, "service") {
		return errors.New("token must be valid and include the \"service\" scope")
	}

	return nil
}
