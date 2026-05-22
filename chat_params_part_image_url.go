package gollum

type ImageURLPart struct {
	URL string
}

func (i ImageURLPart) Type() string {
	return PartTypeImageURL
}

func (i ImageURLPart) Text() string {
	return ""
}

func NewImageURLPart(url string) ImageURLPart {
	return ImageURLPart{
		URL: url,
	}
}
