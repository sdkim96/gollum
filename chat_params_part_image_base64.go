package gollum

type ImageBase64Part struct {
	Base64   string
	MimeType string
}

func (i ImageBase64Part) Type() string {
	return PartTypeImageBase64
}

func (i ImageBase64Part) Text() string {
	return ""
}

func NewImageBase64Part(base64, mimeType string) ImageBase64Part {
	return ImageBase64Part{
		Base64:   base64,
		MimeType: mimeType,
	}
}
