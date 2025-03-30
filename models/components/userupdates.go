package components

type UserUpdatesAnime struct {
	Entry         *AnimeMeta `json:"entry,omitempty"`
	Score         *int64     `json:"score,omitempty"`
	Status        *string    `json:"status,omitempty"`
	EpisodesSeen  *int64     `json:"episodes_seen,omitempty"`
	EpisodesTotal *int64     `json:"episodes_total,omitempty"`
	// ISO8601 format
	Date *string `json:"date,omitempty"`
}

func (o *UserUpdatesAnime) GetEntry() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *UserUpdatesAnime) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *UserUpdatesAnime) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UserUpdatesAnime) GetEpisodesSeen() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesSeen
}

func (o *UserUpdatesAnime) GetEpisodesTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesTotal
}

func (o *UserUpdatesAnime) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type UserUpdatesManga struct {
	Entry         *MangaMeta `json:"entry,omitempty"`
	Score         *int64     `json:"score,omitempty"`
	Status        *string    `json:"status,omitempty"`
	ChaptersRead  *int64     `json:"chapters_read,omitempty"`
	ChaptersTotal *int64     `json:"chapters_total,omitempty"`
	VolumesRead   *int64     `json:"volumes_read,omitempty"`
	VolumesTotal  *int64     `json:"volumes_total,omitempty"`
	// ISO8601 format
	Date *string `json:"date,omitempty"`
}

func (o *UserUpdatesManga) GetEntry() *MangaMeta {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *UserUpdatesManga) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *UserUpdatesManga) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UserUpdatesManga) GetChaptersRead() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersRead
}

func (o *UserUpdatesManga) GetChaptersTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersTotal
}

func (o *UserUpdatesManga) GetVolumesRead() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesRead
}

func (o *UserUpdatesManga) GetVolumesTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesTotal
}

func (o *UserUpdatesManga) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type UserUpdatesData struct {
	// Last updated Anime
	Anime []UserUpdatesAnime `json:"anime,omitempty"`
	// Last updated Manga
	Manga []UserUpdatesManga `json:"manga,omitempty"`
}

func (o *UserUpdatesData) GetAnime() []UserUpdatesAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *UserUpdatesData) GetManga() []UserUpdatesManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

type UserUpdates struct {
	Data *UserUpdatesData `json:"data,omitempty"`
}

func (o *UserUpdates) GetData() *UserUpdatesData {
	if o == nil {
		return nil
	}
	return o.Data
}
