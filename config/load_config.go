package config

import (
	"encoding/json"
	"fmt"
	"go-stable-diffusion-sandbox/pkg/stablediffusion"
	"io"

	"os"
)

// LoadConfig は指定された設定ファイルから設定を読み込みます。
func LoadConfig(configPath string) (*stablediffusion.RequestParams, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var params stablediffusion.RequestParams
	if err := json.Unmarshal(bytes, &params); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &params, nil
}
