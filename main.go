package main

import (
	"fmt"
	stablediffusion "go-stable-diffusion-sandbox/pkg/stable_diffusion"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

func main() {
	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("STABLE_DIFFUSION_API_KEY")
	if apiKey == "" {
		fmt.Println("STABLE_DIFFUSION_API_KEY is not set")
		return
	}
	client := stablediffusion.New(apiKey)

	params := &stablediffusion.RequestParams{
		Prompt:            "ultra realistic close up portrait ((beautiful pale cyberpunk female with heavy black eyeliner))",
		Width:             512,
		Height:            512,
		Samples:           1,
		NumInferenceSteps: 20,
		SafetyChecker:     "no",
		EnhancePrompt:     "yes",
		GuidanceScale:     7.5,
		MultiLingual:      "no",
	}

	response, err := client.Text2Image(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)

	// response output のURLから画像をダウンロードする
	// コードは以下
	//

	if len(response.Output) == 0 {
		fmt.Println("No image URL in response")
		return
	}

	for _, v := range response.Output {
		imageURL := v
		//		imageURL  の拡張子を取得
		// URL解析
		parsedURL, err := url.Parse(imageURL)
		if err != nil {
			fmt.Println("URL parsing error:", err)
			return
		}

		// パスからファイルの拡張子を取得
		ext := path.Ext(parsedURL.Path)

		// 画像をダウンロード
		// ファイル名はタイムスタンプなどを使ってユニークなものにする
		fileName := fmt.Sprintf("%d%s", time.Now().Unix(), ext)
		file, err := os.Create("outputs/" + fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		resp, err := http.Get(imageURL)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer resp.Body.Close()

		io.Copy(file, resp.Body)
	}
}
