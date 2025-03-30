package components

// AnimeEpisode - Anime Episode Resource
type AnimeEpisode struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Title Japanese
	TitleJapanese *string `json:"title_japanese,omitempty"`
	// title_romanji
	TitleRomanji *string `json:"title_romanji,omitempty"`
	// Episode duration in seconds
	Duration *int64 `json:"duration,omitempty"`
	// Aired Date ISO8601
	Aired *string `json:"aired,omitempty"`
	// Filler episode
	Filler *bool `json:"filler,omitempty"`
	// Recap episode
	Recap *bool `json:"recap,omitempty"`
	// Episode Synopsis
	Synopsis *string `json:"synopsis,omitempty"`
}

func (o *AnimeEpisode) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeEpisode) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeEpisode) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AnimeEpisode) GetTitleJapanese() *string {
	if o == nil {
		return nil
	}
	return o.TitleJapanese
}

func (o *AnimeEpisode) GetTitleRomanji() *string {
	if o == nil {
		return nil
	}
	return o.TitleRomanji
}

func (o *AnimeEpisode) GetDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *AnimeEpisode) GetAired() *string {
	if o == nil {
		return nil
	}
	return o.Aired
}

func (o *AnimeEpisode) GetFiller() *bool {
	if o == nil {
		return nil
	}
	return o.Filler
}

func (o *AnimeEpisode) GetRecap() *bool {
	if o == nil {
		return nil
	}
	return o.Recap
}

func (o *AnimeEpisode) GetSynopsis() *string {
	if o == nil {
		return nil
	}
	return o.Synopsis
}
