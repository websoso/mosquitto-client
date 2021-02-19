package mc

type Message struct {
	Code    string `json:"code"`
	Count   int    `json:"count"`
	Order   int    `json:"order"`
	Payload string `json:"payload"`
}

func NewMessage(code string) *Message {

	return &Message{Code: code}

}
