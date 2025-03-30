package components

import (
	"encoding/json"
	"fmt"
)

// MangaFullType - Manga Type
type MangaFullType string

const (
	MangaFullTypeManga      MangaFullType = "Manga"
	MangaFullTypeNovel      MangaFullType = "Novel"
	MangaFullTypeLightNovel MangaFullType = "Light Novel"
	MangaFullTypeOneShot    MangaFullType = "One-shot"
	MangaFullTypeDoujinshi  MangaFullType = "Doujinshi"
	MangaFullTypeManhua     MangaFullType = "Manhua"
	MangaFullTypeManhwa     MangaFullType = "Manhwa"
	MangaFullTypeOel        MangaFullType = "OEL"
)

func (e MangaFullType) ToPointer() *MangaFullType {
	return &e
}

func (e *MangaFullType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Manga":
		fallthrough
	case "Novel":
		fallthrough
	case "Light Novel":
		fallthrough
	case "One-shot":
		fallthrough
	case "Doujinshi":
		fallthrough
	case "Manhua":
		fallthrough
	case "Manhwa":
		fallthrough
	case "OEL":
		*e = MangaFullType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaFullType: %v", v)
	}
}

// MangaFullStatus - Publishing status
type MangaFullStatus string

const (
	MangaFullStatusFinished        MangaFullStatus = "Finished"
	MangaFullStatusPublishing      MangaFullStatus = "Publishing"
	MangaFullStatusOnHiatus        MangaFullStatus = "On Hiatus"
	MangaFullStatusDiscontinued    MangaFullStatus = "Discontinued"
	MangaFullStatusNotYetPublished MangaFullStatus = "Not yet published"
)

func (e MangaFullStatus) ToPointer() *MangaFullStatus {
	return &e
}

func (e *MangaFullStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Finished":
		fallthrough
	case "Publishing":
		fallthrough
	case "On Hiatus":
		fallthrough
	case "Discontinued":
		fallthrough
	case "Not yet published":
		*e = MangaFullStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaFullStatus: %v", v)
	}
}

type MangaFullRelations struct {
	// Relation type
	Relation *string  `json:"relation,omitempty"`
	Entry    []MalURL `json:"entry,omitempty"`
}

func (o *MangaFullRelations) GetRelation() *string {
	if o == nil {
		return nil
	}
	return o.Relation
}

func (o *MangaFullRelations) GetEntry() []MalURL {
	if o == nil {
		return nil
	}
	return o.Entry
}

type MangaFullExternal struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *MangaFullExternal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MangaFullExternal) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// MangaFull - Manga Resource
type MangaFull struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL    *string      `json:"url,omitempty"`
	Images *MangaImages `json:"images,omitempty"`
	// Whether the entry is pending approval on MAL or not
	Approved *bool `json:"approved,omitempty"`
	// All Titles
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
	// Manga Type
	Type *MangaFullType `json:"type,omitempty"`
	// Chapter count
	Chapters *int64 `json:"chapters,omitempty"`
	// Volume count
	Volumes *int64 `json:"volumes,omitempty"`
	// Publishing status
	Status *MangaFullStatus `json:"status,omitempty"`
	// Publishing boolean
	Publishing *bool `json:"publishing,omitempty"`
	// Date range
	Published *Daterange `json:"published,omitempty"`
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
	Background     *string              `json:"background,omitempty"`
	Authors        []MalURL             `json:"authors,omitempty"`
	Serializations []MalURL             `json:"serializations,omitempty"`
	Genres         []MalURL             `json:"genres,omitempty"`
	ExplicitGenres []MalURL             `json:"explicit_genres,omitempty"`
	Themes         []MalURL             `json:"themes,omitempty"`
	Demographics   []MalURL             `json:"demographics,omitempty"`
	Relations      []MangaFullRelations `json:"relations,omitempty"`
	External       []MangaFullExternal  `json:"external,omitempty"`
}

func (o *MangaFull) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *MangaFull) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *MangaFull) GetImages() *MangaImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *MangaFull) GetApproved() *bool {
	if o == nil {
		return nil
	}
	return o.Approved
}

func (o *MangaFull) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *MangaFull) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MangaFull) GetTitleEnglish() *string {
	if o == nil {
		return nil
	}
	return o.TitleEnglish
}

func (o *MangaFull) GetTitleJapanese() *string {
	if o == nil {
		return nil
	}
	return o.TitleJapanese
}

func (o *MangaFull) GetTitleSynonyms() []string {
	if o == nil {
		return nil
	}
	return o.TitleSynonyms
}

func (o *MangaFull) GetType() *MangaFullType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *MangaFull) GetChapters() *int64 {
	if o == nil {
		return nil
	}
	return o.Chapters
}

func (o *MangaFull) GetVolumes() *int64 {
	if o == nil {
		return nil
	}
	return o.Volumes
}

func (o *MangaFull) GetStatus() *MangaFullStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *MangaFull) GetPublishing() *bool {
	if o == nil {
		return nil
	}
	return o.Publishing
}

func (o *MangaFull) GetPublished() *Daterange {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *MangaFull) GetScore() *float32 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *MangaFull) GetScoredBy() *int64 {
	if o == nil {
		return nil
	}
	return o.ScoredBy
}

func (o *MangaFull) GetRank() *int64 {
	if o == nil {
		return nil
	}
	return o.Rank
}

func (o *MangaFull) GetPopularity() *int64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *MangaFull) GetMembers() *int64 {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *MangaFull) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *MangaFull) GetSynopsis() *string {
	if o == nil {
		return nil
	}
	return o.Synopsis
}

func (o *MangaFull) GetBackground() *string {
	if o == nil {
		return nil
	}
	return o.Background
}

func (o *MangaFull) GetAuthors() []MalURL {
	if o == nil {
		return nil
	}
	return o.Authors
}

func (o *MangaFull) GetSerializations() []MalURL {
	if o == nil {
		return nil
	}
	return o.Serializations
}

func (o *MangaFull) GetGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *MangaFull) GetExplicitGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.ExplicitGenres
}

func (o *MangaFull) GetThemes() []MalURL {
	if o == nil {
		return nil
	}
	return o.Themes
}

func (o *MangaFull) GetDemographics() []MalURL {
	if o == nil {
		return nil
	}
	return o.Demographics
}

func (o *MangaFull) GetRelations() []MangaFullRelations {
	if o == nil {
		return nil
	}
	return o.Relations
}

func (o *MangaFull) GetExternal() []MangaFullExternal {
	if o == nil {
		return nil
	}
	return o.External
}
