package tools

import "encoding/json"

func Bind(data interface{}, bindTo interface{}) error {
	bytes, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &bindTo)
}
