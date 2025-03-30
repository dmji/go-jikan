package components

import (
	"encoding/json"
	"fmt"
)

// MangaType - Manga Type
type MangaType string

const (
	MangaTypeManga      MangaType = "Manga"
	MangaTypeNovel      MangaType = "Novel"
	MangaTypeLightNovel MangaType = "Light Novel"
	MangaTypeOneShot    MangaType = "One-shot"
	MangaTypeDoujinshi  MangaType = "Doujinshi"
	MangaTypeManhua     MangaType = "Manhua"
	MangaTypeManhwa     MangaType = "Manhwa"
	MangaTypeOel        MangaType = "OEL"
)

func (e MangaType) ToPointer() *MangaType {
	return &e
}

func (e *MangaType) UnmarshalJSON(data []byte) error {
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
		*e = MangaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaType: %v", v)
	}
}

// MangaStatus - Publishing status
type MangaStatus string

const (
	MangaStatusFinished        MangaStatus = "Finished"
	MangaStatusPublishing      MangaStatus = "Publishing"
	MangaStatusOnHiatus        MangaStatus = "On Hiatus"
	MangaStatusDiscontinued    MangaStatus = "Discontinued"
	MangaStatusNotYetPublished MangaStatus = "Not yet published"
)

func (e MangaStatus) ToPointer() *MangaStatus {
	return &e
}

func (e *MangaStatus) UnmarshalJSON(data []byte) error {
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
		*e = MangaStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaStatus: %v", v)
	}
}

// Manga Resource
type Manga struct {
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
	// Manga Type
	Type *MangaType `json:"type,omitempty"`
	// Chapter count
	Chapters *int64 `json:"chapters,omitempty"`
	// Volume count
	Volumes *int64 `json:"volumes,omitempty"`
	// Publishing status
	Status *MangaStatus `json:"status,omitempty"`
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
	Background     *string  `json:"background,omitempty"`
	Authors        []MalURL `json:"authors,omitempty"`
	Serializations []MalURL `json:"serializations,omitempty"`
	Genres         []MalURL `json:"genres,omitempty"`
	ExplicitGenres []MalURL `json:"explicit_genres,omitempty"`
	Themes         []MalURL `json:"themes,omitempty"`
	Demographics   []MalURL `json:"demographics,omitempty"`
}

func (o *Manga) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Manga) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Manga) GetImages() *MangaImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Manga) GetApproved() *bool {
	if o == nil {
		return nil
	}
	return o.Approved
}

func (o *Manga) GetTitles() []Title {
	if o == nil {
		return nil
	}
	return o.Titles
}

func (o *Manga) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Manga) GetTitleEnglish() *string {
	if o == nil {
		return nil
	}
	return o.TitleEnglish
}

func (o *Manga) GetTitleJapanese() *string {
	if o == nil {
		return nil
	}
	return o.TitleJapanese
}

func (o *Manga) GetType() *MangaType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Manga) GetChapters() *int64 {
	if o == nil {
		return nil
	}
	return o.Chapters
}

func (o *Manga) GetVolumes() *int64 {
	if o == nil {
		return nil
	}
	return o.Volumes
}

func (o *Manga) GetStatus() *MangaStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Manga) GetPublishing() *bool {
	if o == nil {
		return nil
	}
	return o.Publishing
}

func (o *Manga) GetPublished() *Daterange {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *Manga) GetScore() *float32 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *Manga) GetScoredBy() *int64 {
	if o == nil {
		return nil
	}
	return o.ScoredBy
}

func (o *Manga) GetRank() *int64 {
	if o == nil {
		return nil
	}
	return o.Rank
}

func (o *Manga) GetPopularity() *int64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *Manga) GetMembers() *int64 {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *Manga) GetFavorites() *int64 {
	if o == nil {
		return nil
	}
	return o.Favorites
}

func (o *Manga) GetSynopsis() *string {
	if o == nil {
		return nil
	}
	return o.Synopsis
}

func (o *Manga) GetBackground() *string {
	if o == nil {
		return nil
	}
	return o.Background
}

func (o *Manga) GetAuthors() []MalURL {
	if o == nil {
		return nil
	}
	return o.Authors
}

func (o *Manga) GetSerializations() []MalURL {
	if o == nil {
		return nil
	}
	return o.Serializations
}

func (o *Manga) GetGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *Manga) GetExplicitGenres() []MalURL {
	if o == nil {
		return nil
	}
	return o.ExplicitGenres
}

func (o *Manga) GetThemes() []MalURL {
	if o == nil {
		return nil
	}
	return o.Themes
}

func (o *Manga) GetDemographics() []MalURL {
	if o == nil {
		return nil
	}
	return o.Demographics
}
