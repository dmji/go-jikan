package components

type MangaStatisticsScores struct {
	// Scoring value
	Score *int64 `json:"score,omitempty"`
	// Number of votes for this score
	Votes *int64 `json:"votes,omitempty"`
	// Percentage of votes for this score
	Percentage *float32 `json:"percentage,omitempty"`
}

func (o *MangaStatisticsScores) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *MangaStatisticsScores) GetVotes() *int64 {
	if o == nil {
		return nil
	}
	return o.Votes
}

func (o *MangaStatisticsScores) GetPercentage() *float32 {
	if o == nil {
		return nil
	}
	return o.Percentage
}

type MangaStatisticsData struct {
	// Number of users reading the resource
	Reading *int64 `json:"reading,omitempty"`
	// Number of users who have completed the resource
	Completed *int64 `json:"completed,omitempty"`
	// Number of users who have put the resource on hold
	OnHold *int64 `json:"on_hold,omitempty"`
	// Number of users who have dropped the resource
	Dropped *int64 `json:"dropped,omitempty"`
	// Number of users who have planned to read the resource
	PlanToRead *int64 `json:"plan_to_read,omitempty"`
	// Total number of users who have the resource added to their lists
	Total  *int64                  `json:"total,omitempty"`
	Scores []MangaStatisticsScores `json:"scores,omitempty"`
}

func (o *MangaStatisticsData) GetReading() *int64 {
	if o == nil {
		return nil
	}
	return o.Reading
}

func (o *MangaStatisticsData) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *MangaStatisticsData) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *MangaStatisticsData) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *MangaStatisticsData) GetPlanToRead() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToRead
}

func (o *MangaStatisticsData) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MangaStatisticsData) GetScores() []MangaStatisticsScores {
	if o == nil {
		return nil
	}
	return o.Scores
}

// MangaStatistics - Manga Statistics Resource
type MangaStatistics struct {
	Data *MangaStatisticsData `json:"data,omitempty"`
}

func (o *MangaStatistics) GetData() *MangaStatisticsData {
	if o == nil {
		return nil
	}
	return o.Data
}
