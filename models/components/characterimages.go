package components

// CharacterImagesJpg - Available images in JPG
type CharacterImagesJpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL JPG
	SmallImageURL *string `json:"small_image_url,omitempty"`
}

func (o *CharacterImagesJpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *CharacterImagesJpg) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

// CharacterImagesWebp - Available images in WEBP
type CharacterImagesWebp struct {
	// Image URL WEBP
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL WEBP
	SmallImageURL *string `json:"small_image_url,omitempty"`
}

func (o *CharacterImagesWebp) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *CharacterImagesWebp) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

type CharacterImages struct {
	// Available images in JPG
	Jpg *CharacterImagesJpg `json:"jpg,omitempty"`
	// Available images in WEBP
	Webp *CharacterImagesWebp `json:"webp,omitempty"`
}

func (o *CharacterImages) GetJpg() *CharacterImagesJpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}

func (o *CharacterImages) GetWebp() *CharacterImagesWebp {
	if o == nil {
		return nil
	}
	return o.Webp
}
