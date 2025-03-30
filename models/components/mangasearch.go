package components

type MangaSearchItems struct {
	Count   *int64 `json:"count,omitempty"`
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}

func (o *MangaSearchItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *MangaSearchItems) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MangaSearchItems) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type MangaSearchPagination struct {
	LastVisiblePage *int64            `json:"last_visible_page,omitempty"`
	HasNextPage     *bool             `json:"has_next_page,omitempty"`
	Items           *MangaSearchItems `json:"items,omitempty"`
}

func (o *MangaSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *MangaSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

func (o *MangaSearchPagination) GetItems() *MangaSearchItems {
	if o == nil {
		return nil
	}
	return o.Items
}

// MangaSearch - Manga Search Resource
type MangaSearch struct {
	Data       []Manga                `json:"data,omitempty"`
	Pagination *MangaSearchPagination `json:"pagination,omitempty"`
}

func (o *MangaSearch) GetData() []Manga {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *MangaSearch) GetPagination() *MangaSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
