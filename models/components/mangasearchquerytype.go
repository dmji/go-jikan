package components

import (
	"encoding/json"
	"fmt"
)

// MangaSearchQueryType - Available Manga types
type MangaSearchQueryType string

const (
	MangaSearchQueryTypeManga      MangaSearchQueryType = "manga"
	MangaSearchQueryTypeNovel      MangaSearchQueryType = "novel"
	MangaSearchQueryTypeLightnovel MangaSearchQueryType = "lightnovel"
	MangaSearchQueryTypeOneshot    MangaSearchQueryType = "oneshot"
	MangaSearchQueryTypeDoujin     MangaSearchQueryType = "doujin"
	MangaSearchQueryTypeManhwa     MangaSearchQueryType = "manhwa"
	MangaSearchQueryTypeManhua     MangaSearchQueryType = "manhua"
)

func (e MangaSearchQueryType) ToPointer() *MangaSearchQueryType {
	return &e
}

func (e *MangaSearchQueryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manga":
		fallthrough
	case "novel":
		fallthrough
	case "lightnovel":
		fallthrough
	case "oneshot":
		fallthrough
	case "doujin":
		fallthrough
	case "manhwa":
		fallthrough
	case "manhua":
		*e = MangaSearchQueryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaSearchQueryType: %v", v)
	}
}
