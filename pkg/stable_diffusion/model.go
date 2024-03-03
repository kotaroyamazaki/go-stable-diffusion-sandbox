package stablediffusion

type RequestParams struct {
	Key               string  `json:"key"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    *string `json:"negative_prompt,omitempty"`
	Width             int     `json:"width,omitempty"`
	Height            int     `json:"height,omitempty"`
	Samples           int     `json:"samples,omitempty"`
	NumInferenceSteps int     `json:"num_inference_steps,omitempty"`
	SafetyChecker     string  `json:"safety_checker,omitempty"`
	EnhancePrompt     string  `json:"enhance_prompt,omitempty"`
	Seed              *int    `json:"seed,omitempty"`
	GuidanceScale     float64 `json:"guidance_scale,omitempty"`
	MultiLingual      string  `json:"multi_lingual,omitempty"`
	Panorama          string  `json:"panorama,omitempty"`
	SelfAttention     string  `json:"self_attention,omitempty"`
	Upscale           string  `json:"upscale,omitempty"`
	EmbeddingsModel   *string `json:"embeddings_model,omitempty"`
	Webhook           *string `json:"webhook,omitempty"`
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
