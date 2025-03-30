package components

import (
	"encoding/json"
	"fmt"
)

// AnimeType - Anime Type
type AnimeType string

const (
	AnimeTypeTv      AnimeType = "TV"
	AnimeTypeOva     AnimeType = "OVA"
	AnimeTypeMovie   AnimeType = "Movie"
	AnimeTypeSpecial AnimeType = "Special"
	AnimeTypeOna     AnimeType = "ONA"
	AnimeTypeMusic   AnimeType = "Music"
)

func (e AnimeType) ToPointer() *AnimeType {
	return &e
}

func (e *AnimeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TV":
		fallthrough
	case "OVA":
		fallthrough
	case "Movie":
		fallthrough
	case "Special":
		fallthrough
	case "ONA":
		fallthrough
	case "Music":
		*e = AnimeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeType: %v", v)
	}
}

// AnimeStatus - Airing status
type AnimeStatus string

const (
	AnimeStatusFinishedAiring  AnimeStatus = "Finished Airing"
	AnimeStatusCurrentlyAiring AnimeStatus = "Currently Airing"
	AnimeStatusNotYetAired     AnimeStatus = "Not yet aired"
)

func (e AnimeStatus) ToPointer() *AnimeStatus {
	return &e
}

func (e *AnimeStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Finished Airing":
		fallthrough
	case "Currently Airing":
		fallthrough
	case "Not yet aired":
		*e = AnimeStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeStatus: %v", v)
	}
}

// AnimeRating - Anime audience rating
type AnimeRating string

const (
	AnimeRatingGAllAges                    AnimeRating = "G - All Ages"
	AnimeRatingPgChildren                  AnimeRating = "PG - Children"
	AnimeRatingPg13Teens13OrOlder          AnimeRating = "PG-13 - Teens 13 or older"
	AnimeRatingR17PlusViolenceAndProfanity AnimeRating = "R - 17+ (violence & profanity)"
	AnimeRatingRPlusMildNudity             AnimeRating = "R+ - Mild Nudity"
	AnimeRatingRxHentai                    AnimeRating = "Rx - Hentai"
)

func (e AnimeRating) ToPointer() *AnimeRating {
	return &e
}

func (e *AnimeRating) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "G - All Ages":
		fallthrough
	case "PG - Children":
		fallthrough
	case "PG-13 - Teens 13 or older":
		fallthrough
	case "R - 17+ (violence & profanity)":
		fallthrough
	case "R+ - Mild Nudity":
		fallthrough
	case "Rx - Hentai":
		*e = AnimeRating(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeRating: %v", v)
	}
}

// AnimeSeason - Season
type AnimeSeason string

const (
	AnimeSeasonSummer AnimeSeason = "summer"
	AnimeSeasonWinter AnimeSeason = "winter"
	AnimeSeasonSpring AnimeSeason = "spring"
	AnimeSeasonFall   AnimeSeason = "fall"
)

func (e AnimeSeason) ToPointer() *AnimeSeason {
	return &e
}

func (e *AnimeSeason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "summer":
		fallthrough
	case "winter":
		fallthrough
	case "spring":
		fallthrough
	case "fall":
		*e = AnimeSeason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeSeason: %v", v)
	}
}

// Anime Resource
type Anime struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *AnimeImages `json:"images,omitempty"`
	// Youtube Details
	Trailer *TrailerBase `json:"trailer,omitempty"`
	// Whether the entry is pending approval on MAL or not
	Approved *bool `json:"approved,omitempty"`
	// All titles
	Titles []Title `json:"titles,omitempty"`
	// Title
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Title *string `json:"title,omitempty"`
	// English Title
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	TitleEnglish *string `json:"title_english,omitempty"`
	// Japanese Title
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	TitleJapanese *string `json:"title_japanese,omitempty"`
	// Other Titles
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	TitleSynonyms []string `json:"title_synonyms,omitempty"`
	// Anime Type
	Type *AnimeType `json:"type,omitempty"`
	// Original Material/Source adapted from
	Source *string `json:"source,omitempty"`
	// Episode count
	Episodes *int64 `json:"episodes,omitempty"`
	// Airing status
	Status *AnimeStatus `json:"status,omitempty"`
	// Airing boolean
	Airing *bool `json:"airing,omitempty"`
	// Date range
	Aired *Daterange `json:"aired,omitempty"`
	// Parsed raw duration
	Duration *string `json:"duration,omitempty"`
	// Anime audience rating
	Rating *AnimeRating `json:"rating,omitempty"`
	// Score
	Score *float32 `json:"score,omitempty"`
	// Number of users
	ScoredBy *int64 `json:"scored_by,omitempty"`
	// Ranking
	Rank *int64 `json:"rank,omitempty"`
	// Popularity
	Popularity *int64 `json:"popularity,omitempty"`
	// Number of users who have added this entry to their list
	Members *int64 `json:"members,omitempty"`
	// Number of users who have favorited this entry
	Favorites *int64 `json:"favorites,omitempty"`
	// Synopsis
	Synopsis *string `json:"synopsis,omitempty"`
	// Background
	Background *string `json:"background,omitempty"`
	// Season
	Season *AnimeSeason `json:"season,omitempty"`
	// Year
	Year *int64 `json:"year,omitempty"`
	// Broadcast Details
	Broadcast      *Broadcast `json:"broadcast,omitempty"`
	Producers      []MalURL   `json:"producers,omitempty"`
	Licensors      []MalURL   `json:"licensors,omitempty"`
	Studios        []MalURL   `json:"studios,omitempty"`
	Genres         []MalURL   `json:"genres,omitempty"`
	ExplicitGenres []MalURL   `json:"explicit_genres,omitempty"`
	Themes         []MalURL   `json:"themes,omitempty"`
	Demographics   []MalURL   `json:"demographics,omitempty"`
}

func (o *Anime) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Anime) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Anime) GetImages() *AnimeImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Anime) GetTrailer() *TrailerBase {
	if o == nil {
		return nil
	}
	return o.Trailer
}

func (o *Anime) GetApproved() *bool {
	if o == nil {
		return nil
	}
	return o.Approved
}

func (o *Anime) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *Anime) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Anime) GetTitleEnglish() *string {
	if o == nil {
		return nil
	}
	return o.TitleEnglish
}

func (o *Anime) GetTitleJapanese() *string {
	if o == nil {
		return nil
	}
	return o.TitleJapanese
}

func (o *Anime) GetTitleSynonyms() []string {
	if o == nil {
		return nil
	}
	return o.TitleSynonyms
}

func (o *Anime) GetType() *AnimeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Anime) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *Anime) GetEpisodes() *int64 {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *Anime) GetStatus() *AnimeStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Anime) GetAiring() *bool {
	if o == nil {
		return nil
	}
	return o.Airing
}

func (o *Anime) GetAired() *Daterange {
	if o == nil {
		return nil
	}
	return o.Aired
}

func (o *Anime) GetDuration() *string {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *Anime) GetRating() *AnimeRating {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *Anime) GetScore() *float32 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Anime) GetScoredBy() *int64 {
	if o == nil {
		return nil
	}
	return o.ScoredBy
}

func (o *Anime) GetRank() *int64 {
	if o == nil {
		return nil
	}
	return o.Rank
}

func (o *Anime) GetPopularity() *int64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *Anime) GetMembers() *int64 {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *Anime) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *Anime) GetSynopsis() *string {
	if o == nil {
		return nil
	}
	return o.Synopsis
}

func (o *Anime) GetBackground() *string {
	if o == nil {
		return nil
	}
	return o.Background
}

func (o *Anime) GetSeason() *AnimeSeason {
	if o == nil {
		return nil
	}
	return o.Season
}

func (o *Anime) GetYear() *int64 {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *Anime) GetBroadcast() *Broadcast {
	if o == nil {
		return nil
	}
	return o.Broadcast
}

func (o *Anime) GetProducers() []MalURL {
	if o == nil {
		return nil
	}
	return o.Producers
}

func (o *Anime) GetLicensors() []MalURL {
	if o == nil {
		return nil
	}
	return o.Licensors
}

func (o *Anime) GetStudios() []MalURL {
	if o == nil {
		return nil
	}
	return o.Studios
}

func (o *Anime) GetGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *Anime) GetExplicitGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.ExplicitGenres
}

func (o *Anime) GetThemes() []MalURL {
	if o == nil {
		return nil
	}
	return o.Themes
}

func (o *Anime) GetDemographics() []MalURL {
	if o == nil {
		return nil
	}
	return o.Demographics
}
