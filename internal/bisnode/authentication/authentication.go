package authentication

import (
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationClient struct {
	endpoint string
}

func New(endpoint string) *AuthenticationClient {
	return &AuthenticationClient{endpoint: endpoint}
}

func (ac *AuthenticationClient) GetToken(scope string) (*http.Response, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", scope)

	req, err := http.NewRequest(http.MethodPost, ac.endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("BCI_CLIENT_ID", "BCI_CLIENT_SECRET")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
