package components

// Reactions - User reaction count on the review
type Reactions struct {
	// Overall reaction count
	Overall *int64 `json:"overall,omitempty"`
	// Nice reaction count
	Nice *int64 `json:"nice,omitempty"`
	// Love it reaction count
	LoveIt *int64 `json:"love_it,omitempty"`
	// Funny reaction count
	Funny *int64 `json:"funny,omitempty"`
	// Confusing reaction count
	Confusing *int64 `json:"confusing,omitempty"`
	// Informative reaction count
	Informative *int64 `json:"informative,omitempty"`
	// Well written reaction count
	WellWritten *int64 `json:"well_written,omitempty"`
	// Creative reaction count
	Creative *int64 `json:"creative,omitempty"`
}

func (o *Reactions) GetOverall() *int64 {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *Reactions) GetNice() *int64 {
	if o == nil {
		return nil
	}
	return o.Nice
}

func (o *Reactions) GetLoveIt() *int64 {
	if o == nil {
		return nil
	}
	return o.LoveIt
}

func (o *Reactions) GetFunny() *int64 {
	if o == nil {
		return nil
	}
	return o.Funny
}

func (o *Reactions) GetConfusing() *int64 {
	if o == nil {
		return nil
	}
	return o.Confusing
}

func (o *Reactions) GetInformative() *int64 {
	if o == nil {
		return nil
	}
	return o.Informative
}

func (o *Reactions) GetWellWritten() *int64 {
	if o == nil {
		return nil
	}
	return o.WellWritten
}

func (o *Reactions) GetCreative() *int64 {
	if o == nil {
		return nil
	}
	return o.Creative
}

type AnimeReviewsData struct {
	User *UserMeta `json:"user,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList review URL
	URL *string `json:"url,omitempty"`
	// Entry type
	Type *string `json:"type,omitempty"`
	// User reaction count on the review
	Reactions *Reactions `json:"reactions,omitempty"`
	// Review created date ISO8601
	Date *string `json:"date,omitempty"`
	// Review content
	Review *string `json:"review,omitempty"`
	// Number of user votes on the Review
	Score *int64 `json:"score,omitempty"`
	// Review tags
	Tags []string `json:"tags,omitempty"`
	// The review contains spoiler
	IsSpoiler *bool `json:"is_spoiler,omitempty"`
	// The review was made before the entry was completed
	IsPreliminary *bool `json:"is_preliminary,omitempty"`
	// Number of episodes watched
	EpisodesWatched *int64 `json:"episodes_watched,omitempty"`
}

func (o *AnimeReviewsData) GetUser() *UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *AnimeReviewsData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeReviewsData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeReviewsData) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AnimeReviewsData) GetReactions() *Reactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *AnimeReviewsData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *AnimeReviewsData) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *AnimeReviewsData) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *AnimeReviewsData) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AnimeReviewsData) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *AnimeReviewsData) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

func (o *AnimeReviewsData) GetEpisodesWatched() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesWatched
}

type AnimeReviewsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *AnimeReviewsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *AnimeReviewsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// AnimeReviews - Anime Reviews Resource
type AnimeReviews struct {
	Data       []AnimeReviewsData      `json:"data,omitempty"`
	Pagination *AnimeReviewsPagination `json:"pagination,omitempty"`
}

func (o *AnimeReviews) GetData() []AnimeReviewsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnimeReviews) GetPagination() *AnimeReviewsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
