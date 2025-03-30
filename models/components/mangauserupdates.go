package components

type MangaUserupdatesData struct {
	User *UserMeta `json:"user,omitempty"`
	// User Score
	Score *int64 `json:"score,omitempty"`
	// User list status
	Status *string `json:"status,omitempty"`
	// Number of volumes read
	VolumesRead *int64 `json:"volumes_read,omitempty"`
	// Total number of volumes
	VolumesTotal *int64 `json:"volumes_total,omitempty"`
	// Number of chapters read
	ChaptersRead *int64 `json:"chapters_read,omitempty"`
	// Total number of chapters
	ChaptersTotal *int64 `json:"chapters_total,omitempty"`
	// Last updated date ISO8601
	Date *string `json:"date,omitempty"`
}

func (o *MangaUserupdatesData) GetUser() *UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *MangaUserupdatesData) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *MangaUserupdatesData) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *MangaUserupdatesData) GetVolumesRead() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesRead
}

func (o *MangaUserupdatesData) GetVolumesTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesTotal
}

func (o *MangaUserupdatesData) GetChaptersRead() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersRead
}

func (o *MangaUserupdatesData) GetChaptersTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersTotal
}

func (o *MangaUserupdatesData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type MangaUserupdatesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *MangaUserupdatesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *MangaUserupdatesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// MangaUserupdates - Manga User Updates Resource
type MangaUserupdates struct {
	Data       []MangaUserupdatesData      `json:"data,omitempty"`
	Pagination *MangaUserupdatesPagination `json:"pagination,omitempty"`
}

func (o *MangaUserupdates) GetData() []MangaUserupdatesData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *MangaUserupdates) GetPagination() *MangaUserupdatesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
