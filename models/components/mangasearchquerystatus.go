package components

import (
	"encoding/json"
	"fmt"
)

// MangaSearchQueryStatus - Available Manga statuses
type MangaSearchQueryStatus string

const (
	MangaSearchQueryStatusPublishing   MangaSearchQueryStatus = "publishing"
	MangaSearchQueryStatusComplete     MangaSearchQueryStatus = "complete"
	MangaSearchQueryStatusHiatus       MangaSearchQueryStatus = "hiatus"
	MangaSearchQueryStatusDiscontinued MangaSearchQueryStatus = "discontinued"
	MangaSearchQueryStatusUpcoming     MangaSearchQueryStatus = "upcoming"
)

func (e MangaSearchQueryStatus) ToPointer() *MangaSearchQueryStatus {
	return &e
}

func (e *MangaSearchQueryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "publishing":
		fallthrough
	case "complete":
		fallthrough
	case "hiatus":
		fallthrough
	case "discontinued":
		fallthrough
	case "upcoming":
		*e = MangaSearchQueryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaSearchQueryStatus: %v", v)
	}
}
