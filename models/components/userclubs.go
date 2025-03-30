package components

type UserClubsData struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// Club Name
	Name *string `json:"name,omitempty"`
	// Club URL
	URL *string `json:"url,omitempty"`
}

func (o *UserClubsData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *UserClubsData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UserClubsData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type UserClubsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *UserClubsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *UserClubsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// UserClubs - User Clubs
type UserClubs struct {
	Data       []UserClubsData      `json:"data,omitempty"`
	Pagination *UserClubsPagination `json:"pagination,omitempty"`
}

func (o *UserClubs) GetData() []UserClubsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UserClubs) GetPagination() *UserClubsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
