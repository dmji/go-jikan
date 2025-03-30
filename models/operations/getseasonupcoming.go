package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-jikan/models/components"
)

// GetSeasonUpcomingQueryParamFilter - Entry types
type GetSeasonUpcomingQueryParamFilter string

const (
	GetSeasonUpcomingQueryParamFilterTv      GetSeasonUpcomingQueryParamFilter = "tv"
	GetSeasonUpcomingQueryParamFilterMovie   GetSeasonUpcomingQueryParamFilter = "movie"
	GetSeasonUpcomingQueryParamFilterOva     GetSeasonUpcomingQueryParamFilter = "ova"
	GetSeasonUpcomingQueryParamFilterSpecial GetSeasonUpcomingQueryParamFilter = "special"
	GetSeasonUpcomingQueryParamFilterOna     GetSeasonUpcomingQueryParamFilter = "ona"
	GetSeasonUpcomingQueryParamFilterMusic   GetSeasonUpcomingQueryParamFilter = "music"
)

func (e GetSeasonUpcomingQueryParamFilter) ToPointer() *GetSeasonUpcomingQueryParamFilter {
	return &e
}

func (e *GetSeasonUpcomingQueryParamFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tv":
		fallthrough
	case "movie":
		fallthrough
	case "ova":
		fallthrough
	case "special":
		fallthrough
	case "ona":
		fallthrough
	case "music":
		*e = GetSeasonUpcomingQueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSeasonUpcomingQueryParamFilter: %v", v)
	}
}

type GetSeasonUpcomingRequest struct {
	// Entry types
	Filter *GetSeasonUpcomingQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// 'Safe For Work'. This is a flag. When supplied it will filter out entries according to the SFW Policy. You do not need to pass a value to it. e.g usage: `?sfw`
	Sfw *bool `queryParam:"style=form,explode=true,name=sfw"`
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool `queryParam:"style=form,explode=true,name=unapproved"`
	// This is a flag. When supplied it will include entries which are continuing from previous seasons. MAL includes these items on the seasons view in the &#8243;TV (continuing)&#8243; section. (Example: https://myanimelist.net/anime/season/2024/winter) <br />Example usage: `?continuing`
	Continuing *bool  `queryParam:"style=form,explode=true,name=continuing"`
	Page       *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetSeasonUpcomingRequest) GetFilter() *GetSeasonUpcomingQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetSeasonUpcomingRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetSeasonUpcomingRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetSeasonUpcomingRequest) GetContinuing() *bool {
	if o == nil {
		return nil
	}
	return o.Continuing
}

func (o *GetSeasonUpcomingRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetSeasonUpcomingRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetSeasonUpcomingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns upcoming season's anime
	AnimeSearch *components.AnimeSearch
}

func (o *GetSeasonUpcomingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSeasonUpcomingResponse) GetAnimeSearch() *components.AnimeSearch {
	if o == nil {
		return nil
	}
	return o.AnimeSearch
}
