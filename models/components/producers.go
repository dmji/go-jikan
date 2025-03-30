package components

type ProducersPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *ProducersPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *ProducersPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// Producers Collection Resource
type Producers struct {
	Data       []Producer           `json:"data,omitempty"`
	Pagination *ProducersPagination `json:"pagination,omitempty"`
}

func (o *Producers) GetData() []Producer {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Producers) GetPagination() *ProducersPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
