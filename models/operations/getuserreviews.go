package operations

import (
	"errors"
	"fmt"

	"github.com/dmji/go-jikan/internal/utils"
	"github.com/dmji/go-jikan/models/components"
)

type GetUserReviewsRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	Page     *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetUserReviewsRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserReviewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// GetUserReviewsDataUsersReactions - User reaction count on the review
type GetUserReviewsDataUsersReactions struct {
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

func (o *GetUserReviewsDataUsersReactions) GetOverall() *int64 {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *GetUserReviewsDataUsersReactions) GetNice() *int64 {
	if o == nil {
		return nil
	}
	return o.Nice
}

func (o *GetUserReviewsDataUsersReactions) GetLoveIt() *int64 {
	if o == nil {
		return nil
	}
	return o.LoveIt
}

func (o *GetUserReviewsDataUsersReactions) GetFunny() *int64 {
	if o == nil {
		return nil
	}
	return o.Funny
}

func (o *GetUserReviewsDataUsersReactions) GetConfusing() *int64 {
	if o == nil {
		return nil
	}
	return o.Confusing
}

func (o *GetUserReviewsDataUsersReactions) GetInformative() *int64 {
	if o == nil {
		return nil
	}
	return o.Informative
}

func (o *GetUserReviewsDataUsersReactions) GetWellWritten() *int64 {
	if o == nil {
		return nil
	}
	return o.WellWritten
}

func (o *GetUserReviewsDataUsersReactions) GetCreative() *int64 {
	if o == nil {
		return nil
	}
	return o.Creative
}

type Data2 struct {
	User  *components.UserMeta  `json:"user,omitempty"`
	Manga *components.MangaMeta `json:"manga,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList review URL
	URL *string `json:"url,omitempty"`
	// Entry type
	Type *string `json:"type,omitempty"`
	// User reaction count on the review
	Reactions *GetUserReviewsDataUsersReactions `json:"reactions,omitempty"`
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

func (o *Data2) GetUser() *components.UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Data2) GetManga() *components.MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *Data2) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Data2) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Data2) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Data2) GetReactions() *GetUserReviewsDataUsersReactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *Data2) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Data2) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *Data2) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Data2) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Data2) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *Data2) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

// GetUserReviewsDataReactions - User reaction count on the review
type GetUserReviewsDataReactions struct {
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

func (o *GetUserReviewsDataReactions) GetOverall() *int64 {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *GetUserReviewsDataReactions) GetNice() *int64 {
	if o == nil {
		return nil
	}
	return o.Nice
}

func (o *GetUserReviewsDataReactions) GetLoveIt() *int64 {
	if o == nil {
		return nil
	}
	return o.LoveIt
}

func (o *GetUserReviewsDataReactions) GetFunny() *int64 {
	if o == nil {
		return nil
	}
	return o.Funny
}

func (o *GetUserReviewsDataReactions) GetConfusing() *int64 {
	if o == nil {
		return nil
	}
	return o.Confusing
}

func (o *GetUserReviewsDataReactions) GetInformative() *int64 {
	if o == nil {
		return nil
	}
	return o.Informative
}

func (o *GetUserReviewsDataReactions) GetWellWritten() *int64 {
	if o == nil {
		return nil
	}
	return o.WellWritten
}

func (o *GetUserReviewsDataReactions) GetCreative() *int64 {
	if o == nil {
		return nil
	}
	return o.Creative
}

type Data1 struct {
	User  *components.UserMeta  `json:"user,omitempty"`
	Anime *components.AnimeMeta `json:"anime,omitempty"`
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList review URL
	URL *string `json:"url,omitempty"`
	// Entry type
	Type *string `json:"type,omitempty"`
	// User reaction count on the review
	Reactions *GetUserReviewsDataReactions `json:"reactions,omitempty"`
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

func (o *Data1) GetUser() *components.UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Data1) GetAnime() *components.AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *Data1) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Data1) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Data1) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Data1) GetReactions() *GetUserReviewsDataReactions {
	if o == nil {
		return nil
	}
	return o.Reactions
}

func (o *Data1) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Data1) GetReview() *string {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *Data1) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Data1) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Data1) GetIsSpoiler() *bool {
	if o == nil {
		return nil
	}
	return o.IsSpoiler
}

func (o *Data1) GetIsPreliminary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPreliminary
}

func (o *Data1) GetEpisodesWatched() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodesWatched
}

type GetUserReviewsUsersDataType string

const (
	GetUserReviewsUsersDataTypeData1 GetUserReviewsUsersDataType = "data_1"
	GetUserReviewsUsersDataTypeData2 GetUserReviewsUsersDataType = "data_2"
)

type GetUserReviewsUsersData struct {
	Data1 *Data1 `queryParam:"inline"`
	Data2 *Data2 `queryParam:"inline"`

	Type GetUserReviewsUsersDataType
}

func CreateGetUserReviewsUsersDataData1(data1 Data1) GetUserReviewsUsersData {
	typ := GetUserReviewsUsersDataTypeData1

	return GetUserReviewsUsersData{
		Data1: &data1,
		Type:  typ,
	}
}

func CreateGetUserReviewsUsersDataData2(data2 Data2) GetUserReviewsUsersData {
	typ := GetUserReviewsUsersDataTypeData2

	return GetUserReviewsUsersData{
		Data2: &data2,
		Type:  typ,
	}
}

func (u *GetUserReviewsUsersData) UnmarshalJSON(data []byte) error {
	var data2 Data2 = Data2{}
	if err := utils.UnmarshalJSON(data, &data2, "", true, true); err == nil {
		u.Data2 = &data2
		u.Type = GetUserReviewsUsersDataTypeData2
		return nil
	}

	var data1 Data1 = Data1{}
	if err := utils.UnmarshalJSON(data, &data1, "", true, true); err == nil {
		u.Data1 = &data1
		u.Type = GetUserReviewsUsersDataTypeData1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetUserReviewsUsersData", string(data))
}

func (u GetUserReviewsUsersData) MarshalJSON() ([]byte, error) {
	if u.Data1 != nil {
		return utils.MarshalJSON(u.Data1, "", true)
	}

	if u.Data2 != nil {
		return utils.MarshalJSON(u.Data2, "", true)
	}

	return nil, errors.New("could not marshal union type GetUserReviewsUsersData: all fields are null")
}

type GetUserReviewsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *GetUserReviewsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *GetUserReviewsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

type GetUserReviewsData struct {
	Data       []GetUserReviewsUsersData `json:"data,omitempty"`
	Pagination *GetUserReviewsPagination `json:"pagination,omitempty"`
}

func (o *GetUserReviewsData) GetData() []GetUserReviewsUsersData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetUserReviewsData) GetPagination() *GetUserReviewsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

// GetUserReviewsResponseBody - Returns user reviews
type GetUserReviewsResponseBody struct {
	Data *GetUserReviewsData `json:"data,omitempty"`
}

func (o *GetUserReviewsResponseBody) GetData() *GetUserReviewsData {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserReviewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user reviews
	Object *GetUserReviewsResponseBody
}

func (o *GetUserReviewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserReviewsResponse) GetObject() *GetUserReviewsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
