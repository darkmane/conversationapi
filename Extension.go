package conversationapi

import "encoding/json"

type Extension struct {
	Type   string
	Values map[string]interface{}
}

func (ext *Extension) UnmarshalJSON(b []byte) error {
	value_map := make(map[string]interface{})
	err := json.Unmarshal(b, value_map)
	if err != nil {
		return err
	}
	ext.Type = value_map["@type"].(string)
	delete(value_map, "@type")
	ext.Values = value_map
	return nil
}
