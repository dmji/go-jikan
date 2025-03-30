package components

type CharacterMeta struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string          `json:"url,omitempty"`
	Images *CharacterImages `json:"images,omitempty"`
	// Entry name
	Name *string `json:"name,omitempty"`
}

func (o *CharacterMeta) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *CharacterMeta) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *CharacterMeta) GetImages() *CharacterImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *CharacterMeta) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
