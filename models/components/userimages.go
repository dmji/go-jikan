package components

// UserImagesJpg - Available images in JPG
type UserImagesJpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
}

func (o *UserImagesJpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

// UserImagesWebp - Available images in WEBP
type UserImagesWebp struct {
	// Image URL WEBP
	ImageURL *string `json:"image_url,omitempty"`
}

func (o *UserImagesWebp) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

type UserImages struct {
	// Available images in JPG
	Jpg *UserImagesJpg `json:"jpg,omitempty"`
	// Available images in WEBP
	Webp *UserImagesWebp `json:"webp,omitempty"`
}

func (o *UserImages) GetJpg() *UserImagesJpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}

func (o *UserImages) GetWebp() *UserImagesWebp {
	if o == nil {
		return nil
	}
	return o.Webp
}
