package components

import (
	"encoding/json"
	"fmt"
)

// UserAnimeListStatusFilter - User's anime list status filter options
type UserAnimeListStatusFilter string

const (
	UserAnimeListStatusFilterAll         UserAnimeListStatusFilter = "all"
	UserAnimeListStatusFilterWatching    UserAnimeListStatusFilter = "watching"
	UserAnimeListStatusFilterCompleted   UserAnimeListStatusFilter = "completed"
	UserAnimeListStatusFilterOnhold      UserAnimeListStatusFilter = "onhold"
	UserAnimeListStatusFilterDropped     UserAnimeListStatusFilter = "dropped"
	UserAnimeListStatusFilterPlantowatch UserAnimeListStatusFilter = "plantowatch"
)

func (e UserAnimeListStatusFilter) ToPointer() *UserAnimeListStatusFilter {
	return &e
}

func (e *UserAnimeListStatusFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "watching":
		fallthrough
	case "completed":
		fallthrough
	case "onhold":
		fallthrough
	case "dropped":
		fallthrough
	case "plantowatch":
		*e = UserAnimeListStatusFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserAnimeListStatusFilter: %v", v)
	}
}
