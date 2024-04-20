package whatsapp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/prohttp"
)

const (
	MessagingProduct = "whatsapp"
	RecipientType    = "individual"
	Object           = "whatsapp_business_account"
)

type MessageType = string

const (
	TypeText     MessageType = "text"
	TypeTemplate MessageType = "template"
)

type ResText struct {
	Body string
}

type ResMessage struct {
	From      string      `json:"from"`
	Timestamp string      `json:"timestamp"`
	Type      MessageType `json:"type"`
	Text      *ResText    `json:"text"`
}

type ResValue struct {
	Messages []ResMessage `json:"messages"`
}

type ResChange struct {
	Value ResValue `json:"value"`
}

type ResEntry struct {
	Id      string      `json:"id"`
	Changes []ResChange `json:"changes"`
}

type ResInfo struct {
	Entry []ResEntry `json:"entry"`
}

type ReqText struct {
	Body       string `json:"body"`
	PreviewUrl bool   `json:"preview_url"`
}

type ReqMessage struct {
	MessagingProduct string      `json:"messaging_product"`
	RecipientType    string      `json:"recipient_type"`
	To               string      `json:"to"`
	Type             MessageType `json:"type"`
	Text             *ReqText    `json:"text"`
}

func NewReqTextMessage(to string, message string) ReqMessage {
	return ReqMessage{
		MessagingProduct: MessagingProduct,
		RecipientType:    RecipientType,
		To:               to,
		Type:             TypeText,
		Text: &ReqText{
			Body:       message,
			PreviewUrl: false,
		},
	}
}

func GetResMessageGin(c *gin.Context) (ResMessage, error) {
	resBody := ResInfo{}
	if err := c.Bind(&resBody); err != nil {
		return ResMessage{}, err
	}
	return ParseResInfo(&resBody)
}

func ParseResInfo(resData *ResInfo) (ResMessage, error) {
	value := (*resData).Entry[0].Changes[0].Value
	if value.Messages == nil || len(value.Messages) == 0 {
		return ResMessage{}, errors.New("[wa]: no messages found")
	}
	message := value.Messages[0]
	return message, nil
}

type Client struct {
	Url       string
	PhoneId   string
	AuthToken string
	Challenge string
}

func (w Client) ValidateWAWebhookGin(c *gin.Context) (challenge string, ok bool) {
	token := c.Query("hub.verify_token")
	challenge = c.Query("hub.challenge")
	if token == w.Challenge {
		return challenge, true
	}
	return "", false
}

func (w Client) SendMessage(body *ReqMessage) error {
	res, err := prohttp.DoRequest[interface{}](prohttp.RequestConfig[ReqMessage]{
		Url:    fmt.Sprintf("%v/%v/messages", w.Url, w.PhoneId),
		Method: http.MethodPost,
		Body:   body,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", w.AuthToken),
		},
	})
	if err != nil {
		return err
	}
	if !res.IsOk() {
		return errors.New(res.RawBody())
	}
	return nil
}
