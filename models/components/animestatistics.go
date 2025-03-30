package components

type Scores struct {
	// Scoring value
	Score *int64 `json:"score,omitempty"`
	// Number of votes for this score
	Votes *int64 `json:"votes,omitempty"`
	// Percentage of votes for this score
	Percentage *float32 `json:"percentage,omitempty"`
}

func (o *Scores) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Scores) GetVotes() *int64 {
	if o == nil {
		return nil
	}
	return o.Votes
}

func (o *Scores) GetPercentage() *float32 {
	if o == nil {
		return nil
	}
	return o.Percentage
}

type AnimeStatisticsData struct {
	// Number of users watching the resource
	Watching *int64 `json:"watching,omitempty"`
	// Number of users who have completed the resource
	Completed *int64 `json:"completed,omitempty"`
	// Number of users who have put the resource on hold
	OnHold *int64 `json:"on_hold,omitempty"`
	// Number of users who have dropped the resource
	Dropped *int64 `json:"dropped,omitempty"`
	// Number of users who have planned to watch the resource
	PlanToWatch *int64 `json:"plan_to_watch,omitempty"`
	// Total number of users who have the resource added to their lists
	Total  *int64   `json:"total,omitempty"`
	Scores []Scores `json:"scores,omitempty"`
}

func (o *AnimeStatisticsData) GetWatching() *int64 {
	if o == nil {
		return nil
	}
	return o.Watching
}

func (o *AnimeStatisticsData) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *AnimeStatisticsData) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *AnimeStatisticsData) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *AnimeStatisticsData) GetPlanToWatch() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToWatch
}

func (o *AnimeStatisticsData) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *AnimeStatisticsData) GetScores() []Scores {
	if o == nil {
		return nil
	}
	return o.Scores
}

// AnimeStatistics - Anime Statistics Resource
type AnimeStatistics struct {
	Data *AnimeStatisticsData `json:"data,omitempty"`
}

func (o *AnimeStatistics) GetData() *AnimeStatisticsData {
	if o == nil {
		return nil
	}
	return o.Data
}
