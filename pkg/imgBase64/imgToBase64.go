package imgBase64

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
)

func ImgToBase64(filePath string) (string, error) {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)
  var base64String string
	switch mimeType {
	case "image/jpeg":
    base64String, err =  jpgToBase64(bytes)
    if err != nil {
      return "", err
    }
	case "image/png":
    base64String, err = pngToBase64(bytes)
    if err != nil {
      return "", err
    }
	case "image/gif":
    base64String, err =  gifToBase64(bytes)
    if err != nil {
      return "", err
    }
	default:
		return "", errors.New("Unsupported file type")
	}
  //copy string to clipboard
  clipboard.WriteAll(base64String)
  memorySize := len(base64String)
  fmt.Printf("Copied to clipboard with siz: %v kb", (memorySize * 8) / 1024)
  return base64String, nil
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
