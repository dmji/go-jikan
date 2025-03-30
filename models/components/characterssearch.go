package components

type CharactersSearchItems struct {
	Count   *int64 `json:"count,omitempty"`
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}

func (o *CharactersSearchItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CharactersSearchItems) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *CharactersSearchItems) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type CharactersSearchPagination struct {
	LastVisiblePage *int64                 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool                  `json:"has_next_page,omitempty"`
	Items           *CharactersSearchItems `json:"items,omitempty"`
}

func (o *CharactersSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *CharactersSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

func (o *CharactersSearchPagination) GetItems() *CharactersSearchItems {
	if o == nil {
		return nil
	}
	return o.Items
}

// CharactersSearch - Characters Search Resource
type CharactersSearch struct {
	Data       []Character                 `json:"data,omitempty"`
	Pagination *CharactersSearchPagination `json:"pagination,omitempty"`
}

func (o *CharactersSearch) GetData() []Character {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *CharactersSearch) GetPagination() *CharactersSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
