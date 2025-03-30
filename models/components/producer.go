package components

// Producer - Producers Resource
type Producer struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// All titles
	Titles []Title       `json:"titles,omitempty"`
	Images *CommonImages `json:"images,omitempty"`
	// Producers's member favorites count
	Favorites *int64 `json:"favorites,omitempty"`
	// Producers's anime count
	Count *int64 `json:"count,omitempty"`
	// Established Date ISO8601
	Established *string `json:"established,omitempty"`
	// About the Producer
	About *string `json:"about,omitempty"`
}

func (o *Producer) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Producer) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Producer) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *Producer) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Producer) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *Producer) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Producer) GetEstablished() *string {
	if o == nil {
		return nil
	}
	return o.Established
}

func (o *Producer) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}
