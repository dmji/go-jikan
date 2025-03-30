package components

type UsersSearchData struct {
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// MyAnimeList Username
	Username *string     `json:"username,omitempty"`
	Images   *UserImages `json:"images,omitempty"`
	// Last Online Date ISO8601
	LastOnline *string `json:"last_online,omitempty"`
}

func (o *UsersSearchData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UsersSearchData) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UsersSearchData) GetImages() *UserImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *UsersSearchData) GetLastOnline() *string {
	if o == nil {
		return nil
	}
	return o.LastOnline
}

type UsersSearchPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *UsersSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *UsersSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// UsersSearch - User Results
type UsersSearch struct {
	Data       []UsersSearchData      `json:"data,omitempty"`
	Pagination *UsersSearchPagination `json:"pagination,omitempty"`
}

func (o *UsersSearch) GetData() []UsersSearchData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UsersSearch) GetPagination() *UsersSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
