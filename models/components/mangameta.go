package components

type MangaMeta struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *MangaImages `json:"images,omitempty"`
	// Entry title
	Title *string `json:"title,omitempty"`
}

func (o *MangaMeta) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *MangaMeta) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *MangaMeta) GetImages() *MangaImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *MangaMeta) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
