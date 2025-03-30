package components

type WatchPromosData struct {
	Entry   *AnimeMeta `json:"entry,omitempty"`
	Trailer []Trailer  `json:"trailer,omitempty"`
}

func (o *WatchPromosData) GetEntry() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *WatchPromosData) GetTrailer() []Trailer {
	if o == nil {
		return nil
	}
	return o.Trailer
}

type WatchPromosPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *WatchPromosPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *WatchPromosPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// WatchPromos - Watch Promos
type WatchPromos struct {
	// Promo Title
	Title      *string                `json:"title,omitempty"`
	Data       []WatchPromosData      `json:"data,omitempty"`
	Pagination *WatchPromosPagination `json:"pagination,omitempty"`
}

func (o *WatchPromos) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *WatchPromos) GetData() []WatchPromosData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WatchPromos) GetPagination() *WatchPromosPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
