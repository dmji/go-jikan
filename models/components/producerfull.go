package components

type ProducerFullExternal struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *ProducerFullExternal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProducerFullExternal) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// ProducerFull - Producers Resource
type ProducerFull struct {
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
	About    *string                `json:"about,omitempty"`
	External []ProducerFullExternal `json:"external,omitempty"`
}

func (o *ProducerFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *ProducerFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *ProducerFull) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *ProducerFull) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *ProducerFull) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *ProducerFull) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ProducerFull) GetEstablished() *string {
	if o == nil {
		return nil
	}
	return o.Established
}

func (o *ProducerFull) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}

func (o *ProducerFull) GetExternal() []ProducerFullExternal {
	if o == nil {
		return nil
	}
	return o.External
}
