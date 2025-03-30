package components

import (
	"encoding/json"
	"fmt"
)

// UserMangaListStatusFilter - User's anime list status filter options
type UserMangaListStatusFilter string

const (
	UserMangaListStatusFilterAll        UserMangaListStatusFilter = "all"
	UserMangaListStatusFilterReading    UserMangaListStatusFilter = "reading"
	UserMangaListStatusFilterCompleted  UserMangaListStatusFilter = "completed"
	UserMangaListStatusFilterOnhold     UserMangaListStatusFilter = "onhold"
	UserMangaListStatusFilterDropped    UserMangaListStatusFilter = "dropped"
	UserMangaListStatusFilterPlantoread UserMangaListStatusFilter = "plantoread"
)

func (e UserMangaListStatusFilter) ToPointer() *UserMangaListStatusFilter {
	return &e
}

func (e *UserMangaListStatusFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "reading":
		fallthrough
	case "completed":
		fallthrough
	case "onhold":
		fallthrough
	case "dropped":
		fallthrough
	case "plantoread":
		*e = UserMangaListStatusFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserMangaListStatusFilter: %v", v)
	}
}
