# go-stable-diffusion-sandbox

## 概要

このプロジェクトは、Stable Diffusion API を利用して画像生成リクエストを行う Go 言語のクライアントです。設定ファイルからリクエストパラメータを読み込み、指定されたプロンプトに基づいて画像を生成します。

## 設定ファイル

設定ファイル config.json は、画像生成リクエストのパラメータを指定します。以下はそのフォーマットの例です。

```json
{
  "Prompt": "ultra realistic close up portrait ((beautiful pale cyberpunk female with heavy black eyeliner))",
  "Width": 512,
  "Height": 512,
  "Samples": 1,
  "NumInferenceSteps": 20,
  "SafetyChecker": "no",
  "EnhancePrompt": "yes",
  "GuidanceScale": 7.5,
  "MultiLingual": "no"
}
```

## 実行方法

プロジェクトのルートディレクトリで以下のコマンドを実行して、アプリケーションをビルドし実行します。

```
go run main.do
```
