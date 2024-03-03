package stablediffusion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	Text2Image(params *RequestParams) (*ResponseModel, error)
}

type client struct {
	APIKey  string
	BaseURL string
}

func New(apiKey string) Client {
	return &client{
		APIKey:  apiKey,
		BaseURL: "https://stablediffusionapi.com/api",
	}
}

func (c *client) Text2Image(params *RequestParams) (*ResponseModel, error) {
	endpoint := "/v3/text2img"
	params.Key = c.APIKey // APIキーをパラメータに設定

	payloadBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s%s", c.BaseURL, endpoint), "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	print(body)

	var responseModel ResponseModel
	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
