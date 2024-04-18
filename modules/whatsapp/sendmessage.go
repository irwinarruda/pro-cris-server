package whatsapp

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
	"github.com/irwinarruda/pro-cris-server/libs/wa"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func SendMessage(body *wa.ReqMessage) error {
	env := configs.GetEnv()
	req, err := prohttp.NewRequest(prohttp.Config[wa.ReqMessage]{
		Url:    fmt.Sprintf("%v/%v/messages", env.WhatsAppUrl, env.WhatsAppPhoneId),
		Method: http.MethodPost,
		Body:   body,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", env.WhatsAppAuthToken),
			"Content-Type":  "application/json",
		},
	})
	utils.AssertErr(err)

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	utils.AssertErr(err)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Error with the status: %v", res.StatusCode))
	}
	return nil
}
