package components

import (
	"encoding/json"
	"fmt"
)

// Anime Type
type Type string

const (
	TypeTv      Type = "TV"
	TypeOva     Type = "OVA"
	TypeMovie   Type = "Movie"
	TypeSpecial Type = "Special"
	TypeOna     Type = "ONA"
	TypeMusic   Type = "Music"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
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
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// Status - Airing status
type Status string

const (
	StatusFinishedAiring  Status = "Finished Airing"
	StatusCurrentlyAiring Status = "Currently Airing"
	StatusNotYetAired     Status = "Not yet aired"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
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
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// Rating - Anime audience rating
type Rating string

const (
	RatingGAllAges                    Rating = "G - All Ages"
	RatingPgChildren                  Rating = "PG - Children"
	RatingPg13Teens13OrOlder          Rating = "PG-13 - Teens 13 or older"
	RatingR17PlusViolenceAndProfanity Rating = "R - 17+ (violence & profanity)"
	RatingRPlusMildNudity             Rating = "R+ - Mild Nudity"
	RatingRxHentai                    Rating = "Rx - Hentai"
)

func (e Rating) ToPointer() *Rating {
	return &e
}

func (e *Rating) UnmarshalJSON(data []byte) error {
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
		*e = Rating(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Rating: %v", v)
	}
}

// Season
type Season string

const (
	SeasonSummer Season = "summer"
	SeasonWinter Season = "winter"
	SeasonSpring Season = "spring"
	SeasonFall   Season = "fall"
)

func (e Season) ToPointer() *Season {
	return &e
}

func (e *Season) UnmarshalJSON(data []byte) error {
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
		*e = Season(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Season: %v", v)
	}
}

type Relations struct {
	// Relation type
	Relation *string  `json:"relation,omitempty"`
	Entry    []MalURL `json:"entry,omitempty"`
}

func (o *Relations) GetRelation() *string {
	if o == nil {
		return nil
	}
	return o.Relation
}

func (o *Relations) GetEntry() []MalURL {
	if o == nil {
		return nil
	}
	return o.Entry
}

type Theme struct {
	Openings []string `json:"openings,omitempty"`
	Endings  []string `json:"endings,omitempty"`
}

func (o *Theme) GetOpenings() []string {
	if o == nil {
		return nil
	}
	return o.Openings
}

func (o *Theme) GetEndings() []string {
	if o == nil {
		return nil
	}
	return o.Endings
}

type External struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *External) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *External) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type Streaming struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *Streaming) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Streaming) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// AnimeFull - Full anime Resource
type AnimeFull struct {
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
	Type *Type `json:"type,omitempty"`
	// Original Material/Source adapted from
	Source *string `json:"source,omitempty"`
	// Episode count
	Episodes *int64 `json:"episodes,omitempty"`
	// Airing status
	Status *Status `json:"status,omitempty"`
	// Airing boolean
	Airing *bool `json:"airing,omitempty"`
	// Date range
	Aired *Daterange `json:"aired,omitempty"`
	// Parsed raw duration
	Duration *string `json:"duration,omitempty"`
	// Anime audience rating
	Rating *Rating `json:"rating,omitempty"`
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
	Season *Season `json:"season,omitempty"`
	// Year
	Year *int64 `json:"year,omitempty"`
	// Broadcast Details
	Broadcast      *Broadcast  `json:"broadcast,omitempty"`
	Producers      []MalURL    `json:"producers,omitempty"`
	Licensors      []MalURL    `json:"licensors,omitempty"`
	Studios        []MalURL    `json:"studios,omitempty"`
	Genres         []MalURL    `json:"genres,omitempty"`
	ExplicitGenres []MalURL    `json:"explicit_genres,omitempty"`
	Themes         []MalURL    `json:"themes,omitempty"`
	Demographics   []MalURL    `json:"demographics,omitempty"`
	Relations      []Relations `json:"relations,omitempty"`
	Theme          *Theme      `json:"theme,omitempty"`
	External       []External  `json:"external,omitempty"`
	Streaming      []Streaming `json:"streaming,omitempty"`
}

func (o *AnimeFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeFull) GetImages() *AnimeImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeFull) GetTrailer() *TrailerBase {
	if o == nil {
		return nil
	}
	return o.Trailer
}

func (o *AnimeFull) GetApproved() *bool {
	if o == nil {
		return nil
	}
	return o.Approved
}

func (o *AnimeFull) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *AnimeFull) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AnimeFull) GetTitleEnglish() *string {
	if o == nil {
		return nil
	}
	return o.TitleEnglish
}

func (o *AnimeFull) GetTitleJapanese() *string {
	if o == nil {
		return nil
	}
	return o.TitleJapanese
}

func (o *AnimeFull) GetTitleSynonyms() []string {
	if o == nil {
		return nil
	}
	return o.TitleSynonyms
}

func (o *AnimeFull) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AnimeFull) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *AnimeFull) GetEpisodes() *int64 {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *AnimeFull) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AnimeFull) GetAiring() *bool {
	if o == nil {
		return nil
	}
	return o.Airing
}

func (o *AnimeFull) GetAired() *Daterange {
	if o == nil {
		return nil
	}
	return o.Aired
}

func (o *AnimeFull) GetDuration() *string {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *AnimeFull) GetRating() *Rating {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *AnimeFull) GetScore() *float32 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *AnimeFull) GetScoredBy() *int64 {
	if o == nil {
		return nil
	}
	return o.ScoredBy
}

func (o *AnimeFull) GetRank() *int64 {
	if o == nil {
		return nil
	}
	return o.Rank
}

func (o *AnimeFull) GetPopularity() *int64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *AnimeFull) GetMembers() *int64 {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *AnimeFull) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *AnimeFull) GetSynopsis() *string {
	if o == nil {
		return nil
	}
	return o.Synopsis
}

func (o *AnimeFull) GetBackground() *string {
	if o == nil {
		return nil
	}
	return o.Background
}

func (o *AnimeFull) GetSeason() *Season {
	if o == nil {
		return nil
	}
	return o.Season
}

func (o *AnimeFull) GetYear() *int64 {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *AnimeFull) GetBroadcast() *Broadcast {
	if o == nil {
		return nil
	}
	return o.Broadcast
}

func (o *AnimeFull) GetProducers() []MalURL {
	if o == nil {
		return nil
	}
	return o.Producers
}

func (o *AnimeFull) GetLicensors() []MalURL {
	if o == nil {
		return nil
	}
	return o.Licensors
}

func (o *AnimeFull) GetStudios() []MalURL {
	if o == nil {
		return nil
	}
	return o.Studios
}

func (o *AnimeFull) GetGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *AnimeFull) GetExplicitGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.ExplicitGenres
}

func (o *AnimeFull) GetThemes() []MalURL {
	if o == nil {
		return nil
	}
	return o.Themes
}

func (o *AnimeFull) GetDemographics() []MalURL {
	if o == nil {
		return nil
	}
	return o.Demographics
}

func (o *AnimeFull) GetRelations() []Relations {
	if o == nil {
		return nil
	}
	return o.Relations
}

func (o *AnimeFull) GetTheme() *Theme {
	if o == nil {
		return nil
	}
	return o.Theme
}

func (o *AnimeFull) GetExternal() []External {
	if o == nil {
		return nil
	}
	return o.External
}

func (o *AnimeFull) GetStreaming() []Streaming {
	if o == nil {
		return nil
	}
	return o.Streaming
}
