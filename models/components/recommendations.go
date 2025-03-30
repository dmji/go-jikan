package components

import (
	"errors"
	"fmt"

	"github.com/dmji/jikan/internal/utils"
)

type RecommendationsEntryType string

const (
	RecommendationsEntryTypeAnimeMeta RecommendationsEntryType = "anime_meta"
	RecommendationsEntryTypeMangaMeta RecommendationsEntryType = "manga_meta"
)

type RecommendationsEntry struct {
	AnimeMeta *AnimeMeta `queryParam:"inline"`
	MangaMeta *MangaMeta `queryParam:"inline"`

	Type RecommendationsEntryType
}

func CreateRecommendationsEntryAnimeMeta(animeMeta AnimeMeta) RecommendationsEntry {
	typ := RecommendationsEntryTypeAnimeMeta

	return RecommendationsEntry{
		AnimeMeta: &animeMeta,
		Type:      typ,
	}
}

func CreateRecommendationsEntryMangaMeta(mangaMeta MangaMeta) RecommendationsEntry {
	typ := RecommendationsEntryTypeMangaMeta

	return RecommendationsEntry{
		MangaMeta: &mangaMeta,
		Type:      typ,
	}
}

func (u *RecommendationsEntry) UnmarshalJSON(data []byte) error {
	var animeMeta AnimeMeta = AnimeMeta{}
	if err := utils.UnmarshalJSON(data, &animeMeta, "", true, true); err == nil {
		u.AnimeMeta = &animeMeta
		u.Type = RecommendationsEntryTypeAnimeMeta
		return nil
	}

	var mangaMeta MangaMeta = MangaMeta{}
	if err := utils.UnmarshalJSON(data, &mangaMeta, "", true, true); err == nil {
		u.MangaMeta = &mangaMeta
		u.Type = RecommendationsEntryTypeMangaMeta
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RecommendationsEntry", string(data))
}

func (u RecommendationsEntry) MarshalJSON() ([]byte, error) {
	if u.AnimeMeta != nil {
		return utils.MarshalJSON(u.AnimeMeta, "", true)
	}

	if u.MangaMeta != nil {
		return utils.MarshalJSON(u.MangaMeta, "", true)
	}

	return nil, errors.New("could not marshal union type RecommendationsEntry: all fields are null")
}

type RecommendationsData struct {
	// MAL IDs of recommendations is both of the MAL ID's with a `-` delimiter
	MalID *string `json:"mal_id,omitempty"`
	// Array of 2 entries that are being recommended to each other
	Entry []RecommendationsEntry `json:"entry,omitempty"`
	// Recommendation context provided by the user
	Content *string `json:"content,omitempty"`
	// User Meta By ID
	User *UserByID `json:"user,omitempty"`
}

func (o *RecommendationsData) GetMalID() *string {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *RecommendationsData) GetEntry() []RecommendationsEntry {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *RecommendationsData) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *RecommendationsData) GetUser() *UserByID {
	if o == nil {
		return nil
	}
	return o.User
}

type RecommendationsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *RecommendationsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *RecommendationsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// Recommendations
type Recommendations struct {
	Data       []RecommendationsData      `json:"data,omitempty"`
	Pagination *RecommendationsPagination `json:"pagination,omitempty"`
}

func (o *Recommendations) GetData() []RecommendationsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Recommendations) GetPagination() *RecommendationsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
