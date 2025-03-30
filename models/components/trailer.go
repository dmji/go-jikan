package components

type Images struct {
	// Default Image Size URL (120x90)
	ImageURL *string `json:"image_url,omitempty"`
	// Small Image Size URL (640x480)
	SmallImageURL *string `json:"small_image_url,omitempty"`
	// Medium Image Size URL (320x180)
	MediumImageURL *string `json:"medium_image_url,omitempty"`
	// Large Image Size URL (480x360)
	LargeImageURL *string `json:"large_image_url,omitempty"`
	// Maximum Image Size URL (1280x720)
	MaximumImageURL *string `json:"maximum_image_url,omitempty"`
}

func (o *Images) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *Images) GetSmallImageURL() *string {
	if o == nil {
		return nil
	}
	return o.SmallImageURL
}

func (o *Images) GetMediumImageURL() *string {
	if o == nil {
		return nil
	}
	return o.MediumImageURL
}

func (o *Images) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

func (o *Images) GetMaximumImageURL() *string {
	if o == nil {
		return nil
	}
	return o.MaximumImageURL
}

// Trailer - Youtube Details
type Trailer struct {
	// YouTube ID
	YoutubeID *string `json:"youtube_id,omitempty"`
	// YouTube URL
	URL *string `json:"url,omitempty"`
	// Parsed Embed URL
	EmbedURL *string `json:"embed_url,omitempty"`
	Images   *Images `json:"images,omitempty"`
}

func (o *Trailer) GetYoutubeID() *string {
	if o == nil {
		return nil
	}
	return o.YoutubeID
}

func (o *Trailer) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Trailer) GetEmbedURL() *string {
	if o == nil {
		return nil
	}
	return o.EmbedURL
}

func (o *Trailer) GetImages() *Images {
	if o == nil {
		return nil
	}
	return o.Images
}
