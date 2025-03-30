package components

type Items struct {
	Count   *int64 `json:"count,omitempty"`
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"per_page,omitempty"`
}

func (o *Items) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Items) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *Items) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type SchedulesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
	Items           *Items `json:"items,omitempty"`
}

func (o *SchedulesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *SchedulesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

func (o *SchedulesPagination) GetItems() *Items {
	if o == nil {
		return nil
	}
	return o.Items
}

// Schedules - Anime resources currently airing
type Schedules struct {
	Data       []Anime              `json:"data,omitempty"`
	Pagination *SchedulesPagination `json:"pagination,omitempty"`
}

func (o *Schedules) GetData() []Anime {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Schedules) GetPagination() *SchedulesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
