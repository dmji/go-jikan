package components

type PeopleSearchItems struct {
	Count   *int64 `json:"count,omitempty"`
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}

func (o *PeopleSearchItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *PeopleSearchItems) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *PeopleSearchItems) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type PeopleSearchPagination struct {
	LastVisiblePage *int64             `json:"last_visible_page,omitempty"`
	HasNextPage     *bool              `json:"has_next_page,omitempty"`
	Items           *PeopleSearchItems `json:"items,omitempty"`
}

func (o *PeopleSearchPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *PeopleSearchPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

func (o *PeopleSearchPagination) GetItems() *PeopleSearchItems {
	if o == nil {
		return nil
	}
	return o.Items
}

// PeopleSearch - People Search
type PeopleSearch struct {
	Data       []Person                `json:"data,omitempty"`
	Pagination *PeopleSearchPagination `json:"pagination,omitempty"`
}

func (o *PeopleSearch) GetData() []Person {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PeopleSearch) GetPagination() *PeopleSearchPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
