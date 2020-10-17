package pretty

import "fmt"
import "encoding/json"

func JSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func JSONString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s\n", b)
}

func Json(v interface{}) {
	fmt.Printf("%s\n", jstr(v))
}

func JsonString(v interface{}) string {
	return fmt.Sprintf("%s\n", jstr(v))
}

func jstr(v interface{}) string {
	var (
		b   []byte
		err error
	)

	switch v.(type) {
	case *json.RawMessage:
		b, err = v.(*json.RawMessage).MarshalJSON()
	case json.RawMessage:
		b, err = v.(json.RawMessage).MarshalJSON()
	default:
		b, err = json.Marshal(v)
	}

	if err != nil {
		panic(err)
	}

	return string(b)
}
