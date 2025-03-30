package components

// CommonImagesJpg - Available images in JPG
type CommonImagesJpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
}

func (o *CommonImagesJpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

type CommonImages struct {
	// Available images in JPG
	Jpg *CommonImagesJpg `json:"jpg,omitempty"`
}

func (o *CommonImages) GetJpg() *CommonImagesJpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}
