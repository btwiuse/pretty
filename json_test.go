package pretty

import "testing"

//import j "encoding/json"
import jsoniter "github.com/json-iterator/go"

func TestJsonString(t *testing.T) {
	b := []byte("魑魅魍魉")
	for i := 0; i <= len(b); i++ {
		var raw jsoniter.RawMessage = b[:i]
		t.Logf("%q", JsonString(raw))
		t.Logf("%q", b[:i])
	}
}
