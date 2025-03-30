package components

type AnimeSearchItems struct {
	Count   *int64 `json:"count,omitempty"`
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}

func (o *AnimeSearchItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *AnimeSearchItems) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *AnimeSearchItems) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type AnimeSearchPagination struct {
	LastVisiblePage *int64            `json:"last_visible_page,omitempty"`
	HasNextPage     *bool             `json:"has_next_page,omitempty"`
	Items           *AnimeSearchItems `json:"items,omitempty"`
}

func (o *AnimeSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *AnimeSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

func (o *AnimeSearchPagination) GetItems() *AnimeSearchItems {
	if o == nil {
		return nil
	}
	return o.Items
}

// AnimeSearch - Anime Collection Resource
type AnimeSearch struct {
	Data       []Anime                `json:"data,omitempty"`
	Pagination *AnimeSearchPagination `json:"pagination,omitempty"`
}

func (o *AnimeSearch) GetData() []Anime {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnimeSearch) GetPagination() *AnimeSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
