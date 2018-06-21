package message

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	uuidv4, _ = uuid.NewV4()
	msg := Message{uuidv4, uuidv4, []byte("Hello, World!")}
	emsg := &Message{}
	data, _ := json.Marshal(msg)
	emsg.Unmarshal(data)
	if !reflect.DeepEqual(msg, *emsg) {
		t.Errorf("messages should be equal: %s and %s: %s and %s", msg, emsg)
	}
}
