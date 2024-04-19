package whatsapp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
	"github.com/irwinarruda/pro-cris-server/libs/wa"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func SendMessage(body *wa.ReqMessage) error {
	env := configs.GetEnv()
	res, err := prohttp.DoRequest[interface{}](prohttp.RequestConfig[wa.ReqMessage]{
		Url:    fmt.Sprintf("%v/%v/messages", env.WhatsAppUrl, env.WhatsAppPhoneId),
		Method: http.MethodPost,
		Body:   body,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", env.WhatsAppAuthToken),
			"Content-Type":  "application/json",
		},
	})
	utils.AssertErr(err)

	if !res.IsOk() {
		fmt.Println(res.RawBody())
		return errors.New("Could not send message")
	}
	return nil
}
