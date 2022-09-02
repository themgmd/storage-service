package filetype

import "github.com/onemgvv/storage-service/internal/domain"

var (
	images = []string{".png", ".webp", ".svg", ".jpg", ".jpeg", ".bmp"}
	videos = []string{".mp4", ".wav", ".mov"}
	audios = []string{".mp3"}
	docs = []string{".doc", ".docx", ".pptx", ".xlsx", ".csv"}
	text = []string{".txt"}

	types = map[domain.FileType][]string{
		domain.Image: images,
		domain.Video: videos,
		domain.Audio: audios,
		domain.DOCS: docs,
		domain.Text: text,
	}
)

// Detect file type by extension and whitelist
func DetectType(ext string) domain.FileType {
	for k, v := range types {
		for _, v2 := range v {
			if v2 == ext {
				return k
			}
		}
	}

	return ""
}