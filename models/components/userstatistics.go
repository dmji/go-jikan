package components

// UserStatisticsAnime - Anime Statistics
type UserStatisticsAnime struct {
	// Number of days spent watching Anime
	DaysWatched *float32 `json:"days_watched,omitempty"`
	// Mean Score
	MeanScore *float32 `json:"mean_score,omitempty"`
	// Anime Watching
	Watching *int64 `json:"watching,omitempty"`
	// Anime Completed
	Completed *int64 `json:"completed,omitempty"`
	// Anime On-Hold
	OnHold *int64 `json:"on_hold,omitempty"`
	// Anime Dropped
	Dropped *int64 `json:"dropped,omitempty"`
	// Anime Planned to Watch
	PlanToWatch *int64 `json:"plan_to_watch,omitempty"`
	// Total Anime entries on User list
	TotalEntries *int64 `json:"total_entries,omitempty"`
	// Anime re-watched
	Rewatched *int64 `json:"rewatched,omitempty"`
	// Number of Anime Episodes Watched
	EpisodesWatched *int64 `json:"episodes_watched,omitempty"`
}

func (o *UserStatisticsAnime) GetDaysWatched() *float32 {
	if o == nil {
		return nil
	}
	return o.DaysWatched
}

func (o *UserStatisticsAnime) GetMeanScore() *float32 {
	if o == nil {
		return nil
	}
	return o.MeanScore
}

func (o *UserStatisticsAnime) GetWatching() *int64 {
	if o == nil {
		return nil
	}
	return o.Watching
}

func (o *UserStatisticsAnime) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *UserStatisticsAnime) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *UserStatisticsAnime) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *UserStatisticsAnime) GetPlanToWatch() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToWatch
}

func (o *UserStatisticsAnime) GetTotalEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

func (o *UserStatisticsAnime) GetRewatched() *int64 {
	if o == nil {
		return nil
	}
	return o.Rewatched
}

func (o *UserStatisticsAnime) GetEpisodesWatched() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesWatched
}

// UserStatisticsManga - Manga Statistics
type UserStatisticsManga struct {
	// Number of days spent reading Manga
	DaysRead *float32 `json:"days_read,omitempty"`
	// Mean Score
	MeanScore *float32 `json:"mean_score,omitempty"`
	// Manga Reading
	Reading *int64 `json:"reading,omitempty"`
	// Manga Completed
	Completed *int64 `json:"completed,omitempty"`
	// Manga On-Hold
	OnHold *int64 `json:"on_hold,omitempty"`
	// Manga Dropped
	Dropped *int64 `json:"dropped,omitempty"`
	// Manga Planned to Read
	PlanToRead *int64 `json:"plan_to_read,omitempty"`
	// Total Manga entries on User list
	TotalEntries *int64 `json:"total_entries,omitempty"`
	// Manga re-read
	Reread *int64 `json:"reread,omitempty"`
	// Number of Manga Chapters Read
	ChaptersRead *int64 `json:"chapters_read,omitempty"`
	// Number of Manga Volumes Read
	VolumesRead *int64 `json:"volumes_read,omitempty"`
}

func (o *UserStatisticsManga) GetDaysRead() *float32 {
	if o == nil {
		return nil
	}
	return o.DaysRead
}

func (o *UserStatisticsManga) GetMeanScore() *float32 {
	if o == nil {
		return nil
	}
	return o.MeanScore
}

func (o *UserStatisticsManga) GetReading() *int64 {
	if o == nil {
		return nil
	}
	return o.Reading
}

func (o *UserStatisticsManga) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *UserStatisticsManga) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *UserStatisticsManga) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *UserStatisticsManga) GetPlanToRead() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToRead
}

func (o *UserStatisticsManga) GetTotalEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

func (o *UserStatisticsManga) GetReread() *int64 {
	if o == nil {
		return nil
	}
	return o.Reread
}

func (o *UserStatisticsManga) GetChaptersRead() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersRead
}

func (o *UserStatisticsManga) GetVolumesRead() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesRead
}

type UserStatisticsData struct {
	// Anime Statistics
	Anime *UserStatisticsAnime `json:"anime,omitempty"`
	// Manga Statistics
	Manga *UserStatisticsManga `json:"manga,omitempty"`
}

func (o *UserStatisticsData) GetAnime() *UserStatisticsAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *UserStatisticsData) GetManga() *UserStatisticsManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

type UserStatistics struct {
	Data *UserStatisticsData `json:"data,omitempty"`
}

func (o *UserStatistics) GetData() *UserStatisticsData {
	if o == nil {
		return nil
	}
	return o.Data
}
