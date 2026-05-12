package part

type ImagePart struct {
	URL      string
	MimeType string
}

func (i ImagePart) Type() string {
	return "image"
}

func NewImagePart(url, mimeType string) ImagePart {
	return ImagePart{
		URL:      url,
		MimeType: mimeType,
	}
}
