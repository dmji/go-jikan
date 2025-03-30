package components

// Jpg - Available images in JPG
type Jpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL JPG
	SmallImageURL *string `json:"small_image_url,omitempty"`
	// Image URL JPG
	LargeImageURL *string `json:"large_image_url,omitempty"`
}

func (o *Jpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *Jpg) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

func (o *Jpg) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

// Webp - Available images in WEBP
type Webp struct {
	// Image URL WEBP
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image URL WEBP
	SmallImageURL *string `json:"small_image_url,omitempty"`
	// Image URL WEBP
	LargeImageURL *string `json:"large_image_url,omitempty"`
}

func (o *Webp) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *Webp) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

func (o *Webp) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

type AnimeImages struct {
	// Available images in JPG
	Jpg *Jpg `json:"jpg,omitempty"`
	// Available images in WEBP
	Webp *Webp `json:"webp,omitempty"`
}

func (o *AnimeImages) GetJpg() *Jpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}

func (o *AnimeImages) GetWebp() *Webp {
	if o == nil {
		return nil
	}
	return o.Webp
}
