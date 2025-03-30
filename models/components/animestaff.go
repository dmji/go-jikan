package components

// AnimeStaffPerson - Person details
type AnimeStaffPerson struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string       `json:"url,omitempty"`
	Images *PeopleImages `json:"images,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
}

func (o *AnimeStaffPerson) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeStaffPerson) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeStaffPerson) GetImages() *PeopleImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeStaffPerson) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AnimeStaffData struct {
	// Person details
	Person *AnimeStaffPerson `json:"person,omitempty"`
	// Staff Positions
	Positions []string `json:"positions,omitempty"`
}

func (o *AnimeStaffData) GetPerson() *AnimeStaffPerson {
	if o == nil {
		return nil
	}
	return o.Person
}

func (o *AnimeStaffData) GetPositions() []string {
	if o == nil {
		return nil
	}
	return o.Positions
}

// AnimeStaff - Anime Staff Resource
type AnimeStaff struct {
	Data []AnimeStaffData `json:"data,omitempty"`
}

func (o *AnimeStaff) GetData() []AnimeStaffData {
	if o == nil {
		return nil
	}
	return o.Data
}
