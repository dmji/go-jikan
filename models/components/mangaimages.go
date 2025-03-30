package components

// MangaImagesJpg - Available images in JPG
type MangaImagesJpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL JPG
	SmallImageURL *string `json:"small_image_url,omitempty"`
	// Image URL JPG
	LargeImageURL *string `json:"large_image_url,omitempty"`
}

func (o *MangaImagesJpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *MangaImagesJpg) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

func (o *MangaImagesJpg) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

// MangaImagesWebp - Available images in WEBP
type MangaImagesWebp struct {
	// Image URL WEBP
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL WEBP
	SmallImageURL *string `json:"small_image_url,omitempty"`
	// Image URL WEBP
	LargeImageURL *string `json:"large_image_url,omitempty"`
}

func (o *MangaImagesWebp) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *MangaImagesWebp) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

func (o *MangaImagesWebp) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

type MangaImages struct {
	// Available images in JPG
	Jpg *MangaImagesJpg `json:"jpg,omitempty"`
	// Available images in WEBP
	Webp *MangaImagesWebp `json:"webp,omitempty"`
}

func (o *MangaImages) GetJpg() *MangaImagesJpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}

func (o *MangaImages) GetWebp() *MangaImagesWebp {
	if o == nil {
		return nil
	}
	return o.Webp
}
