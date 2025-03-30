package components

// UserProfileFullAnime - Anime Statistics
type UserProfileFullAnime struct {
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

func (o *UserProfileFullAnime) GetDaysWatched() *float32 {
	if o == nil {
		return nil
	}
	return o.DaysWatched
}

func (o *UserProfileFullAnime) GetMeanScore() *float32 {
	if o == nil {
		return nil
	}
	return o.MeanScore
}

func (o *UserProfileFullAnime) GetWatching() *int64 {
	if o == nil {
		return nil
	}
	return o.Watching
}

func (o *UserProfileFullAnime) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *UserProfileFullAnime) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *UserProfileFullAnime) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *UserProfileFullAnime) GetPlanToWatch() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToWatch
}

func (o *UserProfileFullAnime) GetTotalEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

func (o *UserProfileFullAnime) GetRewatched() *int64 {
	if o == nil {
		return nil
	}
	return o.Rewatched
}

func (o *UserProfileFullAnime) GetEpisodesWatched() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesWatched
}

// UserProfileFullManga - Manga Statistics
type UserProfileFullManga struct {
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

func (o *UserProfileFullManga) GetDaysRead() *float32 {
	if o == nil {
		return nil
	}
	return o.DaysRead
}

func (o *UserProfileFullManga) GetMeanScore() *float32 {
	if o == nil {
		return nil
	}
	return o.MeanScore
}

func (o *UserProfileFullManga) GetReading() *int64 {
	if o == nil {
		return nil
	}
	return o.Reading
}

func (o *UserProfileFullManga) GetCompleted() *int64 {
	if o == nil {
		return nil
	}
	return o.Completed
}

func (o *UserProfileFullManga) GetOnHold() *int64 {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *UserProfileFullManga) GetDropped() *int64 {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *UserProfileFullManga) GetPlanToRead() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanToRead
}

func (o *UserProfileFullManga) GetTotalEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

func (o *UserProfileFullManga) GetReread() *int64 {
	if o == nil {
		return nil
	}
	return o.Reread
}

func (o *UserProfileFullManga) GetChaptersRead() *int64 {
	if o == nil {
		return nil
	}
	return o.ChaptersRead
}

func (o *UserProfileFullManga) GetVolumesRead() *int64 {
	if o == nil {
		return nil
	}
	return o.VolumesRead
}

type Statistics struct {
	// Anime Statistics
	Anime *UserProfileFullAnime `json:"anime,omitempty"`
	// Manga Statistics
	Manga *UserProfileFullManga `json:"manga,omitempty"`
}

func (o *Statistics) GetAnime() *UserProfileFullAnime {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *Statistics) GetManga() *UserProfileFullManga {
	if o == nil {
		return nil
	}
	return o.Manga
}

type UserProfileFullExternal struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *UserProfileFullExternal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UserProfileFullExternal) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// UserProfileFull - Transform the resource into an array.
type UserProfileFull struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList Username
	Username *string `json:"username,omitempty"`
	// MyAnimeList URL
	URL    *string     `json:"url,omitempty"`
	Images *UserImages `json:"images,omitempty"`
	// Last Online Date ISO8601
	LastOnline *string `json:"last_online,omitempty"`
	// User Gender
	Gender *string `json:"gender,omitempty"`
	// Birthday Date ISO8601
	Birthday *string `json:"birthday,omitempty"`
	// Location
	Location *string `json:"location,omitempty"`
	// Joined Date ISO8601
	Joined     *string                   `json:"joined,omitempty"`
	Statistics *Statistics               `json:"statistics,omitempty"`
	External   []UserProfileFullExternal `json:"external,omitempty"`
}

func (o *UserProfileFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *UserProfileFull) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UserProfileFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UserProfileFull) GetImages() *UserImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *UserProfileFull) GetLastOnline() *string {
	if o == nil {
		return nil
	}
	return o.LastOnline
}

func (o *UserProfileFull) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UserProfileFull) GetBirthday() *string {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *UserProfileFull) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *UserProfileFull) GetJoined() *string {
	if o == nil {
		return nil
	}
	return o.Joined
}

func (o *UserProfileFull) GetStatistics() *Statistics {
	if o == nil {
		return nil
	}
	return o.Statistics
}

func (o *UserProfileFull) GetExternal() []UserProfileFullExternal {
	if o == nil {
		return nil
	}
	return o.External
}
