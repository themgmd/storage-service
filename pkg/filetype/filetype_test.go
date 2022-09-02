package filetype

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/onemgvv/storage-service/internal/domain"
)

func TestDetectType_Images(t *testing.T) {
	var imageExt = []string{".png", ".webp", ".svg", ".jpg", ".jpeg", ".bmp"}
	for _, v := range imageExt {
		res := DetectType(v)
		assert.Equal(t, res, domain.Image, "Must be an image")
	}
}

func TestDetectType_Videos(t *testing.T) {
	var videoExt = []string{".mp4", ".wav", ".mov"}
	for _, v := range videoExt {
		res := DetectType(v)
		assert.Equal(t, res, domain.Video, "Must be an video")
	}
}

func TestDetectType_Audio(t *testing.T) {
	var audioExt = []string{".mp3"}
	for _, v := range audioExt {
		res := DetectType(v)
		assert.Equal(t, res, domain.Audio, "Must be an audio")
	}
}