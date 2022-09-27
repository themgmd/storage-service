package resize

import (
	"image"
	"os"

	res "github.com/nfnt/resize"
)

func ChangeSize(filePath string, width, height uint) (image.Image, error) {
	// Open file from path
	openedFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	// Convert file to image structure
	img, _, err := image.Decode(openedFile)
	if err != nil {
		return nil, err
	}

	return res.Resize(width, height, img, res.Lanczos2), nil
}
