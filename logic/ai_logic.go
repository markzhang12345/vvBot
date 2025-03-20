package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/kpmark/vvbot/config"
)

type AIImageMatcher struct {
	APIKey     string
	URL        string
	Model      string
	IsAISearch bool
	ImgDir     string
	Filenames  []string
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type RequestBody struct {
	Model            string    `json:"model"`
	Stream           bool      `json:"stream"`
	MaxTokens        int       `json:"max_tokens"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	TopK             int       `json:"top_k"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	N                int       `json:"n"`
	Messages         []Message `json:"messages"`
}

func NewAIImageMatcher(filenames []string) *AIImageMatcher {
	return &AIImageMatcher{
		APIKey:     config.GlobalConfig.AI.APIKey,
		URL:        config.GlobalConfig.AI.URL,
		Model:      config.GlobalConfig.AI.Model,
		Filenames:  filenames,
		IsAISearch: config.GlobalConfig.AI.IsAISearch,
		ImgDir:     "./vvsource",
	}
}

func (m *AIImageMatcher) MatchImageByKeyword(keyword string) string {
	if strings.TrimSpace(keyword) == "" {
		return searchImageByKeyword("", m.Filenames)
	}

	prompt := fmt.Sprintf(`你是一个图片文件选择系统。根据用户搜索关键词"%s"，提取其中最可能的关键信息，从以下文件名列表中选择一个最匹配的图片。
只返回完整准确的文件名，不包含任何路径、引号或附加文本。如果找不到匹配项，选择一个随机的文件名。`, keyword)

	result, err := m.queryAI(prompt)
	if err != nil || !isValidFilename(result, m.Filenames) {
		// AI失败时，回退到原始函数
		filepath := searchImageByKeyword(keyword, m.Filenames)
		// 记录回退选择的日志
		fmt.Println("AI failed, fallback to original search:", err)
		return filepath
	}

	return filepath.Join(m.ImgDir, result)
}

func (m *AIImageMatcher) queryAI(prompt string) (string, error) {
	url := m.URL

	prompt = strings.ReplaceAll(prompt, "\"", "'")

	req := RequestBody{
		Model:            m.Model,
		Stream:           false,
		MaxTokens:        512,
		Temperature:      0.7,
		TopP:             0.7,
		TopK:             50,
		FrequencyPenalty: 0.5,
		N:                1,
		Messages: []Message{
			{
				Content: "你是一个图片文件选择系统。根据用户搜索关键词,提取其中最可能的关键信息，从以下文件名列表中选择一个最匹配的图片。只返回完整准确的文件名，不包含任何路径、引号或附加文本。如果找不到匹配项，选择一个随机的文件名。文件列表如下：\n" + strings.Join(m.Filenames, ", "),
				Role:    "system",
			},
			{
				Content: prompt,
				Role:    "user",
			},
		},
	}

	requestJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(requestJSON))
	if err != nil {
		return "", err
	}

	httpReq.Header.Add("Authorization", "Bearer "+m.APIKey)
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	// 按照新的响应结构解析
	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("无效的响应格式：没有choices字段或为空")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("无效的响应格式：choices[0]不是对象")
	}

	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("无效的响应格式：message字段不存在或类型错误")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("无效的响应格式：content字段不存在或不是字符串")
	}

	return strings.TrimSpace(content), nil
}

// 检查AI返回的文件名是否在有效列表中
func isValidFilename(filename string, validFilenames []string) bool {
	// 清理文件名，移除可能的引号等
	clean := strings.Trim(filename, "\"' \t\n")

	for _, valid := range validFilenames {
		if clean == valid {
			return true
		}
	}

	return false
}

// 提供聊天服务
func (m *AIImageMatcher) Chat(prompt string) (string, error) {
	url := m.URL

	req := RequestBody{
		Model:            m.Model,
		Stream:           false,
		MaxTokens:        512,
		Temperature:      0.7,
		TopP:             0.7,
		TopK:             50,
		FrequencyPenalty: 0.5,
		N:                1,
		Messages: []Message{
			{
				Content: "你是一个资深的coder, 你的每个回答都需要编写一个简单的程序，使用的编程语言仅限Go和Rust, 你非常幽默，会将代码巧妙的融合在回答中，你的回答必须带有 中心对饭，食的套址 或者 中心乐天真，大唬神而唱",
				Role:    "system",
			},
			{
				Content: prompt,
				Role:    "user",
			},
		},
	}

	requestJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(requestJSON))
	if err != nil {
		return "", err
	}

	httpReq.Header.Add("Authorization", "Bearer "+m.APIKey)
	httpReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	// 按照新的响应结构解析
	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("无效的响应格式：没有choices字段或为空")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("无效的响应格式：choices[0]不是对象")
	}

	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("无效的响应格式：message字段不存在或类型错误")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("无效的响应格式：content字段不存在或不是字符串")
	}

	return strings.TrimSpace(content), nil
}
