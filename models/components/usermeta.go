package components

type UserMeta struct {
	// MyAnimeList Username
	Username *string `json:"username,omitempty"`
	// MyAnimeList Profile URL
	URL    *string     `json:"url,omitempty"`
	Images *UserImages `json:"images,omitempty"`
}

func (o *UserMeta) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UserMeta) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserMeta) GetImages() *UserImages {
	if o == nil {
		return nil
	}
	return o.Images
}
