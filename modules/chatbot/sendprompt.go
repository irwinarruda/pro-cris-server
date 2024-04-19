package chatbot

import (
	"fmt"
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/openai"
	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func SendPrompt(text string) openai.ResChat {
	env := configs.GetEnv()
	messages := []openai.Message{
		openai.NewMessage(openai.ChatRoleSystem, "You are a personal trainter with focus on food science"),
		openai.NewMessage(openai.ChatRoleSystem, "Respond with medium to small messages"),
		openai.NewMessage(openai.ChatRoleUser, text),
	}
	chat := openai.NewReqChat(openai.AiModelGpt4Turbo, messages)
	res, err := prohttp.DoRequest[openai.ResChat](prohttp.RequestConfig[openai.ReqChat]{
		Url:    fmt.Sprintf("%v/v1/chat/completions", env.OpenAiUrl),
		Method: http.MethodPost,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %v", env.OpenAiAuthToken),
		},
		Body: &chat,
	})
	utils.AssertErr(err)

	body := openai.ResChat{}
	if !res.IsOk() || !res.ParseBody(&body) {
		fmt.Println(res.Body)
	}
	return body
}
