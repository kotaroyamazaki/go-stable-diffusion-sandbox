package main

import (
	"fmt"
	"go-stable-diffusion-sandbox/config"
	"go-stable-diffusion-sandbox/pkg/stablediffusion"
	"sync"

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

	configPath := "config.json"
	params, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}
	// 読み込んだ設定を使用する
	fmt.Printf("Loaded RequestParams: %+v\n", params)

	response, err := client.Text2Image(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(response.Output) == 0 {
		fmt.Println("No image URL in response")
		return
	}

	var wg sync.WaitGroup
	for _, imageURL := range response.Output {
		wg.Add(1)
		go func(URL string) { // ゴルーチン内でURLを参照するために引数で渡す
			defer wg.Done()

			parsedURL, err := url.Parse(URL)
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

			resp, err := http.Get(URL)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			io.Copy(file, resp.Body)
		}(imageURL)
	}
	wg.Wait()
}
