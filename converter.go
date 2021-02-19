package mc

import (
	"github.com/go-basic/uuid"
	cmap "github.com/orcaman/concurrent-map"
	"math"
	"strings"
)

const MESSAGE_MAX_LENGTH = 1024 * 3

type Converter struct {
	receivedMap cmap.ConcurrentMap
	fallback    Handler
	capacity    int
}

func NewConverter(capacity int, fallback Handler) *Converter {
	if capacity <= 100 {
		capacity = 100
	}
	if capacity >= MESSAGE_MAX_LENGTH {
		capacity = MESSAGE_MAX_LENGTH
	}
	return &Converter{receivedMap: cmap.New(), fallback: fallback, capacity: capacity}

}

func (c *Converter) encode(payload string) *Wrapper {
	length := len([]rune(payload))
	var count = 0
	// 将载荷分割后放入切片
	var payloadList []string
	if length <= c.capacity {
		count = 1
		payloadList = append(payloadList, payload)
	} else {
		count = int(math.Ceil(float64(length) / float64(c.capacity)))
		for i := 1; i <= count; i++ {
			from := (i - 1) * c.capacity
			to := i * c.capacity
			if to > length {
				to = length
			}
			payloadList = append(payloadList, string([]rune(payload)[from:to]))
		}
	}
	// 生成 wrapper
	var code = UUID()
	wrapper := NewWrapper(len(payloadList))
	for i := 1; i <= len(payloadList); i++ {
		message := NewMessage(code)
		message.Count = len(payloadList)
		message.Order = i
		message.Payload = payloadList[i-1]
		wrapper.messageList[i-1] = message
	}
	return wrapper
}

func (c *Converter) decode(m *Message) {
	var completed bool
	var wrapper *Wrapper
	wrapperObj, ok := c.receivedMap.Get(m.Code)
	if !ok {
		wrapper = NewWrapper(m.Count)
		completed = wrapper.Add(m)
		c.receivedMap.Set(m.Code, wrapper)
	} else {
		wrapper = wrapperObj.(*Wrapper)
		completed = wrapper.Add(m)
	}
	if completed {
		c.complete(m.Code)
	}
}

func (c *Converter) complete(key string) {
	wrapperObj, ok := c.receivedMap.Get(key)
	if ok {
		message := ""
		wrapper := wrapperObj.(*Wrapper)
		for _, v := range wrapper.messageList {
			message += v.Payload
		}
		c.fallback(message)
	}
}

func UUID() string {
	str := uuid.New()
	str = strings.Replace(str, "-", "", -1)
	str = strings.ToUpper(str)
	return str
}
