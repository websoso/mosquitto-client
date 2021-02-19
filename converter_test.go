package mc

import (
	"encoding/json"
	"testing"
)

func TestConverter_encode(t *testing.T) {

	converter := NewConverter(3, fallback)
	converter.encode("ABCDE")

}

func Test_decode(t *testing.T) {
	payload1 := "{\"code\":\"528056E749CA4288FAC798BCD06EC6C1\",\"count\":3,\"order\":1,\"payload\":\"梁召峰\"}"
	payload2 := "{\"code\":\"528056E749CA4288FAC798BCD06EC6C1\",\"count\":3,\"order\":2,\"payload\":\"是BB\"}"
	payload3 := "{\"code\":\"528056E749CA4288FAC798BCD06EC6C1\",\"count\":3,\"order\":3,\"payload\":\"CCC\"}"
	converter := NewConverter(3, fallback)

	var message1 *Message
	json.Unmarshal([]byte(payload1), &message1)
	converter.decode(message1)

	var message2 *Message
	json.Unmarshal([]byte(payload2), &message2)
	converter.decode(message2)

	var message3 *Message
	json.Unmarshal([]byte(payload3), &message3)
	converter.decode(message3)
}
