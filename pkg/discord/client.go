package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Client 構造体はDiscordクライアントの設定を保持します
type Client struct {
	Token   string
	Session *discordgo.Session
}

// New は新しいDiscordClientを初期化して返します
func New(token string) (*Client, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Client{
		Token:   token,
		Session: session,
	}, nil
}

// Start はDiscordサーバーへの接続を開始します
func (client *Client) Start() error {
	client.Session.AddHandler(messageCreate)
	err := client.Session.Open()
	if err != nil {
		return err
	}
	fmt.Println("Botは正常に動作しています。Ctrl+Cで終了します。")
	return nil
}

// messageCreate は、新しいメッセージがサーバーに投稿された時に呼び出される関数です
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "こんにちは！")
	}
}

// SendMessage は指定されたチャンネルにメッセージを送信します
func (client *Client) SendMessage(channelID, message string) error {
	_, err := client.Session.ChannelMessageSend(channelID, message)
	return err
}
