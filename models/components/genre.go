package components

// Genre Resource
type Genre struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// Genre Name
	Name *string `json:"name,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Genre's entry count
	Count *int64 `json:"count,omitempty"`
}

func (o *Genre) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Genre) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Genre) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Genre) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}
