package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

func ToSnakeCaseOrSlug(input string) string {
	// 1. Trim spasi di awal dan akhir string
	cleanInput := strings.TrimSpace(input)
	if cleanInput == "" {
		return ""
	}

	// 2. Ubah ke lowercase
	lower := strings.ToLower(cleanInput)

	// 3. Ganti spasi atau karakter khusus berturut-turut menjadi satu underscore "_"
	snake := nonAlphanumericRegex.ReplaceAllString(lower, "_")

	// 4. Hapus underscore di awal atau akhir string jika ada
	return strings.Trim(snake, "_")
}

func StringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("gagal konversi '%s' ke int: %w", input, err)
	}
	return num, nil
}

func IsNullOrWhitespaceString(input *string) bool {

	result := strings.TrimSpace(*input)

	if result == "" {
		return true
	}

	return false
}

const RequestPhoneNumberCancelButtonPrefix = "\u200B" // warna trasparan jadi ga kelihat di HP sebagai penanda untuk button Cancel

func IsHasRequestPhoneNumberCancelButtonPrefix(input string) bool {
	if strings.HasPrefix(input, RequestPhoneNumberCancelButtonPrefix+RequestPhoneNumberCancelButtonPrefix) {
		return true
	}

	return false
}

func IsIndonesianLanguageButtonPrefix(input, language string) bool {
	if strings.HasPrefix(input, fmt.Sprintf("%s %s", RequestPhoneNumberCancelButtonPrefix, language)) {
		return true
	}

	return false
}
