package imgBase64

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"os"
)

func ImgToBase64(filePath string) (string, error) {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		return jpgToBase64(bytes)
	case "image/png":
		return pngToBase64(bytes)
	case "image/gif":
		return gifToBase64(bytes)
	default:
		return "", errors.New("Unsupported file type")
	}
}

func toBase64(imgBytes []byte) string {
	return base64.StdEncoding.EncodeToString(imgBytes)
}

func pngToBase64(imgBytes []byte) (string, error) {
	base64String := toBase64(imgBytes)
	return "data:image/png;base64," + base64String, nil
}

func jpgToBase64(imgBytes []byte) (string, error) {
	base64String := toBase64(imgBytes)
	return "data:image/jpg;base64," + base64String, nil
}

func gifToBase64(imgBytes []byte) (string, error) {
	base64String := toBase64(imgBytes)
	return "data:image/gif;base64," + base64String, nil
}
