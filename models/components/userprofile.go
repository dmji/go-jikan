package components

type UserProfile struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList Username
	Username *string `json:"username,omitempty"`
	// MyAnimeList URL
	URL    *string     `json:"url,omitempty"`
	Images *UserImages `json:"images,omitempty"`
	// Last Online Date ISO8601
	LastOnline *string `json:"last_online,omitempty"`
	// User Gender
	Gender *string `json:"gender,omitempty"`
	// Birthday Date ISO8601
	Birthday *string `json:"birthday,omitempty"`
	// Location
	Location *string `json:"location,omitempty"`
	// Joined Date ISO8601
	Joined *string `json:"joined,omitempty"`
}

func (o *UserProfile) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *UserProfile) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UserProfile) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserProfile) GetImages() *UserImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *UserProfile) GetLastOnline() *string {
	if o == nil {
		return nil
	}
	return o.LastOnline
}

func (o *UserProfile) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UserProfile) GetBirthday() *string {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *UserProfile) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *UserProfile) GetJoined() *string {
	if o == nil {
		return nil
	}
	return o.Joined
}
