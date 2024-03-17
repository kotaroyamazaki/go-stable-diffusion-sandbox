package main

import (
	"fmt"
	"go-stable-diffusion-sandbox/config"
	"go-stable-diffusion-sandbox/pkg/discord"
	"go-stable-diffusion-sandbox/pkg/stablediffusion"

	"os"
)

func main() {
	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("STABLE_DIFFUSION_API_KEY")
	if apiKey == "" {
		fmt.Println("STABLE_DIFFUSION_API_KEY is not set")
		return
	}
	client := stablediffusion.New(apiKey)

	// discordに画像を送信する
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	if discordToken == "" {
		fmt.Println("DISCORD_BOT_TOKEN is not set")
		return
	}
	discordChannelID := os.Getenv("DISCORD_CHANNEL_ID")
	if discordChannelID == "" {
		fmt.Println("DISCORD_CHANNEL_ID is not set")
		return
	}
	discordClient, err := discord.New(discordToken)
	if err != nil {
		fmt.Println("Error creating Discord client:", err)
		return
	}

	configPath := "config.json"
	params, err := config.LoadConfigToTxt2ImgParams(configPath)
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

	var negativePrompt string
	if params.NegativePrompt != nil {
		negativePrompt = *params.NegativePrompt
	}
	if err := discordClient.SendMessage(discordChannelID, fmt.Sprintf("- Prompt: `%s`\n- Negative Prompt: `%s`\n- Seed: `%d`\n- Model: `%s`", params.Prompt, negativePrompt, response.Meta.Seed, response.Meta.Model)); err != nil {
		fmt.Println("Error sending message to Discord:", err)
	}

	//var wg sync.WaitGroup
	for _, imageURL := range response.Output {
		fmt.Println("Downloading", imageURL)
		if err := discordClient.SendMessage(discordChannelID, imageURL); err != nil {
			fmt.Println("Error sending message to Discord:", err)
		}
		// 	wg.Add(1)
		// 	go func(URL string, idx int) { // ゴルーチン内でURLを参照するために引数で渡す
		// 		defer wg.Done()

		// 		parsedURL, err := url.Parse(URL)
		// 		if err != nil {
		// 			fmt.Println("URL parsing error:", err)
		// 			return
		// 		}
		// 		// パスからファイルの拡張子を取得
		// 		ext := path.Ext(parsedURL.Path)

		// 		// 画像をダウンロード
		// 		// ファイル名はタイムスタンプなどを使ってユニークなものにする
		// 		fileName := fmt.Sprintf("%d-%d%s", time.Now().Unix(), idx, ext)
		// 		file, err := os.Create("outputs/" + fileName)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		defer file.Close()

		// 		resp, err := http.Get(URL)
		// 		if err != nil {
		// 			fmt.Printf("Error: %v\n", err)
		// 			return
		// 		}
		// 		defer resp.Body.Close()

		// 		io.Copy(file, resp.Body)
		// 	}(imageURL, index)
	}
	// wg.Wait()
}
