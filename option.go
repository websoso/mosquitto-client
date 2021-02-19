package mc

type Option struct {
	broker     string
	clientId   string
	username   string
	password   string
	clientName string
	capacity   int // 单个消息容量
}

type Handler func(payload string)

func NewOption() *Option {

	return &Option{
		capacity: MESSAGE_MAX_LENGTH,
	}

}

func (o *Option) SetBroker(broker string) {
	o.broker = broker
}

func (o *Option) SetClientId(clientId string) {
	o.clientId = clientId
}

func (o *Option) SetUsername(username string) {
	o.username = username
}

func (o *Option) SetPassword(password string) {
	o.password = password
}

func (o *Option) SetClientName(clientName string) {
	o.clientName = clientName
}

func (o *Option) SetCapacity(capacity int) {
	o.capacity = capacity
}
