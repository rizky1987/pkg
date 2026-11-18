package utils

import (
	"encoding/json"
	"fmt"
)

// ToPrettyJSONString mengonversi struct/interface menjadi JSON string berformat (Pretty Print)
// Sangat berguna untuk kebutuhan logging atau debugging di terminal
func ToPrettyJSONString(v interface{}) string {
	if v == nil {
		return ""
	}

	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(bytes)
}
