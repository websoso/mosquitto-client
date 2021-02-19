package mc

type Wrapper struct {
	messageList []*Message
}

func NewWrapper(length int) *Wrapper {
	wrapper := &Wrapper{}
	for i := 0; i < length; i++ {
		wrapper.messageList = append(wrapper.messageList, nil)
	}
	return wrapper
}

func (w *Wrapper) Add(message *Message) (completed bool) {
	w.messageList[message.Order-1] = message
	completed = true
	for _, v := range w.messageList {
		if v == nil {
			completed = false
		}
	}
	return completed
}
