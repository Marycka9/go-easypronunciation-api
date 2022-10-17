package client

import (
	"encoding/json"
	"github.com/marycka9/go-easypronunciation-api/entities"
	"net/http"
)

type Client struct {
	Client      *http.Client
	AccessToken string `json:"access_token"`
}

func NewClient(accessToken string) *Client {
	return &Client{
		Client:      http.DefaultClient,
		AccessToken: accessToken,
	}
}

func (c *Client) Close() {
	c.Close()
}

func (c *Client) PhoneticTranslator(lang, phrase string, base64 bool) (*entities.PhoneticTranslatorResponse, error) {
	translator := entities.NewPhoneticTranslatorRequest(phrase, lang, base64)

	url, err := translator.GetUrl(c.AccessToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var response *entities.PhoneticTranslatorResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return response, nil
}
