package openai

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
