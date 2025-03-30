package components

import (
	"errors"
	"fmt"

	"github.com/dmji/jikan/internal/utils"
)

type EntryType string

const (
	EntryTypeAnimeMeta EntryType = "anime_meta"
	EntryTypeMangaMeta EntryType = "manga_meta"
)

type Entry struct {
	AnimeMeta *AnimeMeta `queryParam:"inline"`
	MangaMeta *MangaMeta `queryParam:"inline"`

	Type EntryType
}

func CreateEntryAnimeMeta(animeMeta AnimeMeta) Entry {
	typ := EntryTypeAnimeMeta

	return Entry{
		AnimeMeta: &animeMeta,
		Type:      typ,
	}
}

func CreateEntryMangaMeta(mangaMeta MangaMeta) Entry {
	typ := EntryTypeMangaMeta

	return Entry{
		MangaMeta: &mangaMeta,
		Type:      typ,
	}
}

func (u *Entry) UnmarshalJSON(data []byte) error {
	var animeMeta AnimeMeta = AnimeMeta{}
	if err := utils.UnmarshalJSON(data, &animeMeta, "", true, true); err == nil {
		u.AnimeMeta = &animeMeta
		u.Type = EntryTypeAnimeMeta
		return nil
	}

	var mangaMeta MangaMeta = MangaMeta{}
	if err := utils.UnmarshalJSON(data, &mangaMeta, "", true, true); err == nil {
		u.MangaMeta = &mangaMeta
		u.Type = EntryTypeMangaMeta
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Entry", string(data))
}

func (u Entry) MarshalJSON() ([]byte, error) {
	if u.AnimeMeta != nil {
		return utils.MarshalJSON(u.AnimeMeta, "", true)
	}

	if u.MangaMeta != nil {
		return utils.MarshalJSON(u.MangaMeta, "", true)
	}

	return nil, errors.New("could not marshal union type Entry: all fields are null")
}

type EntryRecommendationsData struct {
	Entry *Entry `json:"entry,omitempty"`
}

func (o *EntryRecommendationsData) GetEntry() *Entry {
	if o == nil {
		return nil
	}
	return o.Entry
}

// EntryRecommendations - Entry Recommendations Resource
type EntryRecommendations struct {
	Data []EntryRecommendationsData `json:"data,omitempty"`
}

func (o *EntryRecommendations) GetData() []EntryRecommendationsData {
	if o == nil {
		return nil
	}
	return o.Data
}
