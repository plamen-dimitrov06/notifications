package notifications

type Message struct {
	Recipient	string `json:"recipient"`
	Content		string `json:"content"`
}

func NewMessage() Message { return Message{} }
