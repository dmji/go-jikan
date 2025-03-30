package components

// MangaReviewsReactions - User reaction count on the review
type MangaReviewsReactions struct {
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

func (o *MangaReviewsReactions) GetOverall() *int64 {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *MangaReviewsReactions) GetNice() *int64 {
	if o == nil {
		return nil
	}
	return o.Nice
}

func (o *MangaReviewsReactions) GetLoveIt() *int64 {
	if o == nil {
		return nil
	}
	return o.LoveIt
}

func (o *MangaReviewsReactions) GetFunny() *int64 {
	if o == nil {
		return nil
	}
	return o.Funny
}

func (o *MangaReviewsReactions) GetConfusing() *int64 {
	if o == nil {
		return nil
	}
	return o.Confusing
}

func (o *MangaReviewsReactions) GetInformative() *int64 {
	if o == nil {
		return nil
	}
	return o.Informative
}

func (o *MangaReviewsReactions) GetWellWritten() *int64 {
	if o == nil {
		return nil
	}
	return o.WellWritten
}

func (o *MangaReviewsReactions) GetCreative() *int64 {
	if o == nil {
		return nil
	}
	return o.Creative
}

type MangaReviewsData struct {
	User *UserMeta `json:"user,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList review URL
	URL *string `json:"url,omitempty"`
	// Entry type
	Type *string `json:"type,omitempty"`
	// User reaction count on the review
	Reactions *MangaReviewsReactions `json:"reactions,omitempty"`
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
}

func (o *MangaReviewsData) GetUser() *UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *MangaReviewsData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *MangaReviewsData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *MangaReviewsData) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *MangaReviewsData) GetReactions() *MangaReviewsReactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *MangaReviewsData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *MangaReviewsData) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *MangaReviewsData) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *MangaReviewsData) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *MangaReviewsData) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *MangaReviewsData) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

type MangaReviewsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *MangaReviewsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *MangaReviewsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// MangaReviews - Manga Reviews Resource
type MangaReviews struct {
	Data       []MangaReviewsData      `json:"data,omitempty"`
	Pagination *MangaReviewsPagination `json:"pagination,omitempty"`
}

func (o *MangaReviews) GetData() []MangaReviewsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *MangaReviews) GetPagination() *MangaReviewsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
