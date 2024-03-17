package config

import (
	"encoding/json"
	"fmt"
	"go-stable-diffusion-sandbox/pkg/stablediffusion"
	"io"

	"os"
)

// LoadConfigToTxt2ImgParams は指定された設定ファイルから設定を読み込みます。
func LoadConfigToTxt2ImgParams(configPath string) (*stablediffusion.Txt2ImgRequestParams, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var params stablediffusion.Txt2ImgRequestParams
	if err := json.Unmarshal(bytes, &params); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &params, nil
}

func LoadConfigToImg2ImgParams(configPath string) (*stablediffusion.Img2imgRequestParams, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var params stablediffusion.Img2imgRequestParams
	if err := json.Unmarshal(bytes, &params); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &params, nil
}
