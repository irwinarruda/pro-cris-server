package wa

import "errors"

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

type ResData struct {
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

type ReqData = ReqMessage

func NewReqTextMessage(to string, message *ReqText) (*ReqData, error) {
	if message == nil {
		return nil, errors.New("[WA]: You must pass a ReqText as an argument")
	}
	return &ReqMessage{
		MessagingProduct: MessagingProduct,
		RecipientType:    RecipientType,
		To:               to,
		Type:             TypeText,
		Text:             message,
	}, nil
}

func GetResMessage(resData *ResData) *ResMessage {
	value := (*resData).Entry[0].Changes[0].Value
	if value.Messages == nil {
		return nil
	}
	message := value.Messages[0]
	return &message
}
