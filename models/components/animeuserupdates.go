package components

type AnimeUserupdatesData struct {
	User *UserMeta `json:"user,omitempty"`
	// User Score
	Score *int64 `json:"score,omitempty"`
	// User list status
	Status *string `json:"status,omitempty"`
	// Number of episodes seen
	EpisodesSeen *int64 `json:"episodes_seen,omitempty"`
	// Total number of episodes
	EpisodesTotal *int64 `json:"episodes_total,omitempty"`
	// Last updated date ISO8601
	Date *string `json:"date,omitempty"`
}

func (o *AnimeUserupdatesData) GetUser() *UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *AnimeUserupdatesData) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *AnimeUserupdatesData) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AnimeUserupdatesData) GetEpisodesSeen() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesSeen
}

func (o *AnimeUserupdatesData) GetEpisodesTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesTotal
}

func (o *AnimeUserupdatesData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type AnimeUserupdatesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *AnimeUserupdatesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *AnimeUserupdatesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// AnimeUserupdates - Anime User Updates Resource
type AnimeUserupdates struct {
	Data       []AnimeUserupdatesData      `json:"data,omitempty"`
	Pagination *AnimeUserupdatesPagination `json:"pagination,omitempty"`
}

func (o *AnimeUserupdates) GetData() []AnimeUserupdatesData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnimeUserupdates) GetPagination() *AnimeUserupdatesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
