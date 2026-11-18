package utils

import "os"

// ReplacePlaceholders mengganti placeholder ${key} atau $key berdasarkan map replacements
func ReplacePlaceholders(template string, data map[string]string) string {
	return os.Expand(template, func(key string) string {
		if val, ok := data[key]; ok {
			return val
		}
		return "${" + key + "}" // Kembalikan format awal jika key tidak ditemukan
	})
}
