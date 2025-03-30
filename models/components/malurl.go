package components

// MalURL - Parsed URL Data
type MalURL struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// Type of resource
	Type *string `json:"type,omitempty"`
	// Resource Name/Title
	Name *string `json:"name,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
}

func (o *MalURL) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *MalURL) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *MalURL) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MalURL) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
