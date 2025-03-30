package components

type AnimeMeta struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *AnimeImages `json:"images,omitempty"`
	// Entry title
	Title *string `json:"title,omitempty"`
}

func (o *AnimeMeta) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeMeta) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeMeta) GetImages() *AnimeImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeMeta) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
