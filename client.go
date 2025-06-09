package novitus_gosdk

import (
	"fmt"
	"resty.dev/v3"
)

type NovitusClient struct {
	host  string
	token string
}

func NewNovitusClient(host string) (*NovitusClient, error) {
	client := &NovitusClient{host: host}
	token, err := client.ObtainToken()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain token: %w", err)
	}
	client.token = token.Token
	return client, nil
}

func (n *NovitusClient) ObtainToken() (TokenResponse, error) {
	client := resty.New()
	defer client.Close()
	var tokenResponse TokenResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&tokenResponse).SetError(&errorResponse).Get(n.host + "/api/v1/token")
	if err != nil {
		return TokenResponse{}, fmt.Errorf("failed to obtain token: %w", err)
	}
	if res.IsError() {
		return TokenResponse{}, fmt.Errorf("error obtaining token: %s", errorResponse.Exception.Description)
	}
	return tokenResponse, nil
}

func (n *NovitusClient) RefreshToken() error {
	client := resty.New()
	defer client.Close()
	var tokenResponse TokenResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&tokenResponse).SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Patch(n.host + "/api/v1/token")
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error refreshing token: %s", errorResponse.Exception.Description)
	}
	n.token = tokenResponse.Token
	return nil

}
