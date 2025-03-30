package components

// Magazine Resource
type Magazine struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// Magazine Name
	Name *string `json:"name,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Magazine's manga count
	Count *int64 `json:"count,omitempty"`
}

func (o *Magazine) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Magazine) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Magazine) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Magazine) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}
