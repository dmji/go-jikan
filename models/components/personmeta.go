package components

type PersonMeta struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string       `json:"url,omitempty"`
	Images *PeopleImages `json:"images,omitempty"`
	// Entry name
	Name *string `json:"name,omitempty"`
}

func (o *PersonMeta) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *PersonMeta) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PersonMeta) GetImages() *PeopleImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *PersonMeta) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
