package lablinkv4client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/minerofish/lablink-v4-client-go/option"
)

type AuthenticatedClient interface {
	Do() *Client
}

type authenticatedClient struct {
	client         *Client
	baseURL        url.URL
	username       string
	password       string
	token          AuthToken
	tokenExpiresAt time.Time
}

func NewAuthenticatedClient(baseURL string, username string, password string) (AuthenticatedClient, error) {
	urlParsed, err := url.Parse(baseURL)
	if err != nil {
		return &authenticatedClient{}, err
	}

	a := authenticatedClient{
		username: username,
		password: password,
		baseURL:  *urlParsed,
	}

	err = a.refreshToken()
	if err != nil {
		return &authenticatedClient{}, err
	}

	lablinkClient := NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken(a.token.AccessToken),
	)
	a.client = &lablinkClient

	return &a, nil
}

func (a *authenticatedClient) Do() *Client {
	_ = a.refreshToken()
	return a.client
}

func (a *authenticatedClient) refreshToken() error {
	if !time.Now().After(a.tokenExpiresAt.Add(-10 * time.Second)) {
		return nil
	}

	form := url.Values{}
	form.Set("username", a.username)
	form.Set("password", a.password)

	req, err := http.NewRequest("POST", a.baseURL.JoinPath("/v4/login").String(), bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal to map[string]any
	var authToken AuthToken
	if err := json.Unmarshal(bodyBytes, &authToken); err != nil {
		return err
	}

	a.token = authToken
	a.tokenExpiresAt = time.Now().Add(time.Duration(authToken.ExpiresIn) * time.Second)
	return nil
}
