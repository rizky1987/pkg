package error_generator

import (
	"net/http"
)

var httpStatusMessagesID = map[int]string{
	http.StatusBadRequest:          "Permintaan tidak valid",
	http.StatusUnauthorized:        "Sesi berakhir atau otentikasi gagal",
	http.StatusForbidden:           "Anda tidak memiliki akses ke sumber daya ini",
	http.StatusNotFound:            "Data atau halaman tidak ditemukan",
	http.StatusMethodNotAllowed:    "Metode HTTP tidak diizinkan",
	http.StatusConflict:            "Data sudah ada atau terjadi konflik",
	http.StatusUnprocessableEntity: "Validasi data gagal",
	http.StatusTooManyRequests:     "Terlalu banyak permintaan, silakan coba beberapa saat lagi",
	http.StatusInternalServerError: "Terjadi kesalahan pada sistem server",
	http.StatusBadGateway:          "Gagal terhubung ke layanan hulu (Bad Gateway)",
	http.StatusServiceUnavailable:  "Layanan sedang tidak dapat diakses",
	http.StatusGatewayTimeout:      "Waktu koneksi ke server habis",
}

func GetHTTPErrorMessageID(statusCode int) string {
	if msg, exists := httpStatusMessagesID[statusCode]; exists {
		return msg
	}
	return "Terjadi kesalahan yang tidak diketahui"
}
