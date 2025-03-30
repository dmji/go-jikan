package components

type MagazinesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *MagazinesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *MagazinesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// Magazines - Magazine Collection Resource
type Magazines struct {
	Data       []Magazine           `json:"data,omitempty"`
	Pagination *MagazinesPagination `json:"pagination,omitempty"`
}

func (o *Magazines) GetData() []Magazine {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Magazines) GetPagination() *MagazinesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
