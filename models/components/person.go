package components

// Person Resource
type Person struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Person's website URL
	WebsiteURL *string       `json:"website_url,omitempty"`
	Images     *PeopleImages `json:"images,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Given Name
	GivenName *string `json:"given_name,omitempty"`
	// Family Name
	FamilyName *string `json:"family_name,omitempty"`
	// Other Names
	AlternateNames []string `json:"alternate_names,omitempty"`
	// Birthday Date ISO8601
	Birthday *string `json:"birthday,omitempty"`
	// Number of users who have favorited this entry
	Favorites *int64 `json:"favorites,omitempty"`
	// Biography
	About *string `json:"about,omitempty"`
}

func (o *Person) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Person) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Person) GetWebsiteURL() *string {
	if o == nil {
		return nil
	}
	return o.WebsiteURL
}

func (o *Person) GetImages() *PeopleImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Person) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Person) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *Person) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *Person) GetAlternateNames() []string {
	if o == nil {
		return nil
	}
	return o.AlternateNames
}

func (o *Person) GetBirthday() *string {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *Person) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *Person) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}
