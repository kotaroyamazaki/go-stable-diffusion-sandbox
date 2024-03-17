package stablediffusion

type Txt2ImgRequestParams struct {
	Key string `json:"key"`
	// Text prompt with description of the things you want in the image to be generated.
	Prompt string `json:"prompt"`
	// Items you don't want in the image.
	NegativePrompt *string `json:"negative_prompt,omitempty"`
	// Max Height: Width: 1024x1024.
	Width int `json:"width,omitempty"`
	// Max Height: Width: 1024x1024.
	Height int `json:"height,omitempty"`
	// Number of images to be returned in response. The maximum value is 4.
	Samples int `json:"samples,omitempty"`
	// Number of denoising steps. Available values: 21, 31, 41, 51.
	NumInferenceSteps int `json:"num_inference_steps,omitempty"`
	// A checker for NSFW images. If such an image is detected, it will be replaced by a blank image.
	SafetyChecker string `json:"safety_checker,omitempty"`
	// Enhance prompts for better results; default: yes, options: yes/no.
	EnhancePrompt string `json:"enhance_prompt,omitempty"`
	// Seed is used to reproduce results, same seed will give you same image in return again. Pass null for a random number.
	Seed *int `json:"seed,omitempty"`
	// Scale for classifier-free guidance (minimum: 1; maximum: 20).
	GuidanceScale float64 `json:"guidance_scale,omitempty"`
	// Allow multi lingual prompt to generate images. Use "no" for the default English.
	MultiLingual string `json:"multi_lingual,omitempty"`
	// Set this parameter to "yes" to generate a panorama image.
	Panorama string `json:"panorama,omitempty"`
	// If you want a high quality image, set this parameter to "yes". In this case the image generation will take more time.
	SelfAttention string `json:"self_attention,omitempty"`
	// Set this parameter to "yes" if you want to upscale the given image resolution two times (2x). If the requested resolution is 512 x 512 px, the generated image will be 1024 x 1024 px.
	Upscale string `json:"upscale,omitempty"`
	// This is used to pass an embeddings model (embeddings_model_id).
	EmbeddingsModel *string `json:"embeddings_model,omitempty"`
	// Set an URL to get a POST API call once the image generation is complete.
	Webhook *string `json:"webhook,omitempty"`
	// This ID is returned in the response to the webhook API call. This will be used to identify the webhook request.
	TrackID *string `json:"track_id,omitempty"`
}

type Img2imgRequestParams struct {
	Key               string  `json:"key"`                           // Your API Key used for request authorization.
	Prompt            string  `json:"prompt"`                        // Text prompt with description of the things you want in the image to be generated.
	NegativePrompt    *string `json:"negative_prompt,omitempty"`     // Items you don't want in the image.
	InitImage         string  `json:"init_image"`                    // Link to the Initial Image.
	Width             int     `json:"width,omitempty"`               // Max Height: Width: 1024x1024.
	Height            int     `json:"height,omitempty"`              // Number of images to be returned in response. The maximum value is 4.
	Samples           int     `json:"samples,omitempty"`             // Number of denoising steps. Available values: 21, 31, 41, 51.
	NumInferenceSteps int     `json:"num_inference_steps,omitempty"` // A checker for NSFW images. If such an image is detected, it will be replaced by a blank image.
	SafettyChecker    string  `json:"safety_checker,omitempty"`      // Enhance prompts for better results; default: yes, options: yes/no.
	EnhancePrompt     string  `json:"enhance_prompt,omitempty"`      // Scale for classifier-free guidance (minimum: 1; maximum: 20).
	GuidanceScale     float64 `json:"guidance_scale,omitempty"`      // Prompt strength when using init image. 1.0 corresponds to full destruction of information in the init image.
	Strength          float64 `json:"strength,omitempty"`            // Seed is used to reproduce results, same seed will give you same image in return again. Pass null for a random number.
	Seed              *int    `json:"seed,omitempty"`                // Get response as base64 string, pass init_image as base64 string, to get base64 response. default: "no", options: yes/no
	Base64            string  `json:"base64,omitempty"`              // Set an URL to get a POST API call once the image generation is complete.
	Webhook           *string `json:"webhook,omitempty"`             // This ID is returned in the response to the webhook API call. This will be used to identify the webhook request.
	TrackID           *string `json:"track_id,omitempty"`
}

type ResponseModel struct {
	Status         string   `json:"status"`
	GenerationTime float64  `json:"generationTime"`
	ID             int      `json:"id"`
	Output         []string `json:"output"`
	Meta           MetaData `json:"meta"`
}

// MetaData はAPIレスポンスのmetaフィールドのデータを保持する構造体です。
type MetaData struct {
	H                      int     `json:"H"`
	W                      int     `json:"W"`
	EnableAttentionSlicing string  `json:"enable_attention_slicing"`
	FilePrefix             string  `json:"file_prefix"`
	GuidanceScale          float64 `json:"guidance_scale"`
	Model                  string  `json:"model"`
	NSamples               int     `json:"n_samples"`
	NegativePrompt         string  `json:"negative_prompt"`
	Outdir                 string  `json:"outdir"`
	Prompt                 string  `json:"prompt"`
	Revision               string  `json:"revision"`
	SafetyChecker          string  `json:"safetychecker"`
	Seed                   int64   `json:"seed"`
	Steps                  int     `json:"steps"`
	Vae                    string  `json:"vae"`
}
