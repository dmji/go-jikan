package components

type ClubsSearchPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *ClubsSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *ClubsSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// ClubsSearch - Clubs Search Resource
type ClubsSearch struct {
	Data       []Club                 `json:"data,omitempty"`
	Pagination *ClubsSearchPagination `json:"pagination,omitempty"`
}

func (o *ClubsSearch) GetData() []Club {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ClubsSearch) GetPagination() *ClubsSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
