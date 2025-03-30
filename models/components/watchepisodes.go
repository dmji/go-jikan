package components

type WatchEpisodesEpisodes struct {
	// MyAnimeList ID
	MalID *string `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Episode Title
	Title *string `json:"title,omitempty"`
	// For MyAnimeList Premium Users
	Premium *bool `json:"premium,omitempty"`
}

func (o *WatchEpisodesEpisodes) GetMalID() *string {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *WatchEpisodesEpisodes) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *WatchEpisodesEpisodes) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *WatchEpisodesEpisodes) GetPremium() *bool {
	if o == nil {
		return nil
	}
	return o.Premium
}

type WatchEpisodesData struct {
	Entry *AnimeMeta `json:"entry,omitempty"`
	// Recent Episodes (max 2 listed)
	Episodes []WatchEpisodesEpisodes `json:"episodes,omitempty"`
	// Region Locked Episode
	RegionLocked *bool `json:"region_locked,omitempty"`
}

func (o *WatchEpisodesData) GetEntry() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *WatchEpisodesData) GetEpisodes() []WatchEpisodesEpisodes {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *WatchEpisodesData) GetRegionLocked() *bool {
	if o == nil {
		return nil
	}
	return o.RegionLocked
}

type WatchEpisodesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *WatchEpisodesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *WatchEpisodesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// WatchEpisodes - Watch Episodes
type WatchEpisodes struct {
	Data       []WatchEpisodesData      `json:"data,omitempty"`
	Pagination *WatchEpisodesPagination `json:"pagination,omitempty"`
}

func (o *WatchEpisodes) GetData() []WatchEpisodesData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *WatchEpisodes) GetPagination() *WatchEpisodesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
