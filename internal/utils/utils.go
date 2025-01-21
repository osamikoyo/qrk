package utils

import (
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GetQR(link string, w *http.ResponseWriter) barcode.Barcode {
	qrCode, err := qr.Encode(link, qr.L, qr.Auto)
	if err != nil {
		http.Error(*w, "Failed to generate QR code", http.StatusInternalServerError)
		return nil
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		http.Error(*w, "Failed to scale QR code", http.StatusInternalServerError)
		return nil
	}

	return qrCode
}