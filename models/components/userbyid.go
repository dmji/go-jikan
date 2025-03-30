package components

// UserByID - User Meta By ID
type UserByID struct {
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// MyAnimeList Username
	Username *string `json:"username,omitempty"`
}

func (o *UserByID) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserByID) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
