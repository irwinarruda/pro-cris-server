package openai

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
)

type AiModel = string
type ChatRole = string

const (
	AiModelGpt35Turbo AiModel = "gpt-3.5-turbo"
	AiModelGpt4Turbo  AiModel = "gpt-4-turbo"
)

const (
	ChatRoleSystem ChatRole = "system"
	ChatRoleUser   ChatRole = "user"
)

type Message struct {
	Role    ChatRole `json:"role"`
	Content string   `json:"content"`
}

type ReqChat struct {
	Model    AiModel   `json:"model"`
	Messages []Message `json:"messages"`
}

type ResChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ResUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ResChat struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             AiModel     `json:"model"`
	SystemFingerprint string      `json:"system_fingerprint"`
	Choices           []ResChoice `json:"choices"`
	Usage             ResUsage    `json:"usage"`
}

func NewReqChat(model AiModel, messages []Message) ReqChat {
	return ReqChat{
		Model:    model,
		Messages: messages,
	}
}

func NewMessage(role ChatRole, content string) Message {
	return Message{
		Role:    role,
		Content: content,
	}
}

type Client struct {
	Url       string
	AuthToken string
	Model     AiModel
	System    []string
}

func (c *Client) SendPrompt(prompt string) (ResChat, error) {
	messages := []Message{}
	for _, message := range c.System {
		messages = append(messages, NewMessage(ChatRoleSystem, message))
	}
	messages = append(messages, NewMessage(ChatRoleUser, prompt))
	chat := NewReqChat(c.Model, messages)
	fmt.Println(chat)
	res, err := prohttp.DoRequest[ResChat](prohttp.RequestConfig[ReqChat]{
		Url:    fmt.Sprintf("%v/v1/chat/completions", c.Url),
		Method: http.MethodPost,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", c.AuthToken),
		},
		Body: &chat,
	})
	body := ResChat{}
	if err != nil {
		return body, err
	}
	if !res.IsOk() || !res.ParseBody(&body) {
		return body, errors.New(res.RawBody())
	}
	return body, nil
}
