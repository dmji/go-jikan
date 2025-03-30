package operations

import (
	"errors"
	"fmt"

	"github.com/dmji/go-jikan/internal/utils"
	"github.com/dmji/go-jikan/models/components"
)

type GetTopReviewsRequest struct {
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// The type of reviews to filter by. Defaults to anime.
	Type *components.TopReviewsTypeEnum `queryParam:"style=form,explode=true,name=type"`
	// Whether the results include preliminary reviews or not. Defaults to true.
	Preliminary *bool `queryParam:"style=form,explode=true,name=preliminary"`
	// Whether the results include reviews with spoilers or not. Defaults to true.
	Spoilers *bool `queryParam:"style=form,explode=true,name=spoilers"`
}

func (o *GetTopReviewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetTopReviewsRequest) GetType() *components.TopReviewsTypeEnum {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetTopReviewsRequest) GetPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.Preliminary
}

func (o *GetTopReviewsRequest) GetSpoilers() *bool {
	if o == nil {
		return nil
	}
	return o.Spoilers
}

// DataReactions - User reaction count on the review
type DataReactions struct {
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

func (o *DataReactions) GetOverall() *int64 {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *DataReactions) GetNice() *int64 {
	if o == nil {
		return nil
	}
	return o.Nice
}

func (o *DataReactions) GetLoveIt() *int64 {
	if o == nil {
		return nil
	}
	return o.LoveIt
}

func (o *DataReactions) GetFunny() *int64 {
	if o == nil {
		return nil
	}
	return o.Funny
}

func (o *DataReactions) GetConfusing() *int64 {
	if o == nil {
		return nil
	}
	return o.Confusing
}

func (o *DataReactions) GetInformative() *int64 {
	if o == nil {
		return nil
	}
	return o.Informative
}

func (o *DataReactions) GetWellWritten() *int64 {
	if o == nil {
		return nil
	}
	return o.WellWritten
}

func (o *DataReactions) GetCreative() *int64 {
	if o == nil {
		return nil
	}
	return o.Creative
}

type Two struct {
	User  *components.UserMeta  `json:"user,omitempty"`
	Manga *components.MangaMeta `json:"manga,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList review URL
	URL *string `json:"url,omitempty"`
	// Entry type
	Type *string `json:"type,omitempty"`
	// User reaction count on the review
	Reactions *DataReactions `json:"reactions,omitempty"`
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

func (o *Two) GetUser() *components.UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Two) GetManga() *components.MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *Two) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Two) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Two) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Two) GetReactions() *DataReactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *Two) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Two) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *Two) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Two) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Two) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *Two) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

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

type One struct {
	User  *components.UserMeta  `json:"user,omitempty"`
	Anime *components.AnimeMeta `json:"anime,omitempty"`
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

func (o *One) GetUser() *components.UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *One) GetAnime() *components.AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *One) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *One) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *One) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *One) GetReactions() *Reactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *One) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *One) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *One) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *One) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *One) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *One) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

func (o *One) GetEpisodesWatched() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesWatched
}

type GetTopReviewsTopDataType string

const (
	GetTopReviewsTopDataTypeOne GetTopReviewsTopDataType = "1"
	GetTopReviewsTopDataTypeTwo GetTopReviewsTopDataType = "2"
)

type GetTopReviewsTopData struct {
	One *One `queryParam:"inline"`
	Two *Two `queryParam:"inline"`

	Type GetTopReviewsTopDataType
}

func CreateGetTopReviewsTopDataOne(one One) GetTopReviewsTopData {
	typ := GetTopReviewsTopDataTypeOne

	return GetTopReviewsTopData{
		One:  &one,
		Type: typ,
	}
}

func CreateGetTopReviewsTopDataTwo(two Two) GetTopReviewsTopData {
	typ := GetTopReviewsTopDataTypeTwo

	return GetTopReviewsTopData{
		Two:  &two,
		Type: typ,
	}
}

func (u *GetTopReviewsTopData) UnmarshalJSON(data []byte) error {
	var two Two = Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = GetTopReviewsTopDataTypeTwo
		return nil
	}

	var one One = One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = GetTopReviewsTopDataTypeOne
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetTopReviewsTopData", string(data))
}

func (u GetTopReviewsTopData) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type GetTopReviewsTopData: all fields are null")
}

type GetTopReviewsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *GetTopReviewsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *GetTopReviewsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

type GetTopReviewsData struct {
	Data       []GetTopReviewsTopData   `json:"data,omitempty"`
	Pagination *GetTopReviewsPagination `json:"pagination,omitempty"`
}

func (o *GetTopReviewsData) GetData() []GetTopReviewsTopData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetTopReviewsData) GetPagination() *GetTopReviewsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

// GetTopReviewsResponseBody - Returns top reviews
type GetTopReviewsResponseBody struct {
	Data *GetTopReviewsData `json:"data,omitempty"`
}

func (o *GetTopReviewsResponseBody) GetData() *GetTopReviewsData {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetTopReviewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns top reviews
	Object *GetTopReviewsResponseBody
}

func (o *GetTopReviewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopReviewsResponse) GetObject() *GetTopReviewsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
