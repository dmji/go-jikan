package components

type AnimeVideosEpisodesData struct {
	// MyAnimeList ID or Episode Number
	MalID *int64 `json:"mal_id,omitempty"`
	// Episode Title
	Title *string `json:"title,omitempty"`
	// Episode Subtitle
	Episode *string `json:"episode,omitempty"`
	// Episode Page URL
	URL    *string       `json:"url,omitempty"`
	Images *CommonImages `json:"images,omitempty"`
}

func (o *AnimeVideosEpisodesData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeVideosEpisodesData) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AnimeVideosEpisodesData) GetEpisode() *string {
	if o == nil {
		return nil
	}
	return o.Episode
}

func (o *AnimeVideosEpisodesData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeVideosEpisodesData) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

type AnimeVideosEpisodesPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *AnimeVideosEpisodesPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *AnimeVideosEpisodesPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// AnimeVideosEpisodes - Anime Videos Episodes Resource
type AnimeVideosEpisodes struct {
	Data       []AnimeVideosEpisodesData      `json:"data,omitempty"`
	Pagination *AnimeVideosEpisodesPagination `json:"pagination,omitempty"`
}

func (o *AnimeVideosEpisodes) GetData() []AnimeVideosEpisodesData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnimeVideosEpisodes) GetPagination() *AnimeVideosEpisodesPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
