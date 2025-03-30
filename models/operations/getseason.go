package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/jikan/models/components"
)

// GetSeasonQueryParamFilter - Entry types
type GetSeasonQueryParamFilter string

const (
	GetSeasonQueryParamFilterTv      GetSeasonQueryParamFilter = "tv"
	GetSeasonQueryParamFilterMovie   GetSeasonQueryParamFilter = "movie"
	GetSeasonQueryParamFilterOva     GetSeasonQueryParamFilter = "ova"
	GetSeasonQueryParamFilterSpecial GetSeasonQueryParamFilter = "special"
	GetSeasonQueryParamFilterOna     GetSeasonQueryParamFilter = "ona"
	GetSeasonQueryParamFilterMusic   GetSeasonQueryParamFilter = "music"
)

func (e GetSeasonQueryParamFilter) ToPointer() *GetSeasonQueryParamFilter {
	return &e
}

func (e *GetSeasonQueryParamFilter) UnmarshalJSON(data []byte) error {
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
		*e = GetSeasonQueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSeasonQueryParamFilter: %v", v)
	}
}

type GetSeasonRequest struct {
	Year   int64  `pathParam:"style=simple,explode=false,name=year"`
	Season string `pathParam:"style=simple,explode=false,name=season"`
	// Entry types
	Filter *GetSeasonQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// 'Safe For Work'. This is a flag. When supplied it will filter out entries according to the SFW Policy. You do not need to pass a value to it. e.g usage: `?sfw`
	Sfw *bool `queryParam:"style=form,explode=true,name=sfw"`
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool `queryParam:"style=form,explode=true,name=unapproved"`
	// This is a flag. When supplied it will include entries which are continuing from previous seasons. MAL includes these items on the seasons view in the &#8243;TV (continuing)&#8243; section. (Example: https://myanimelist.net/anime/season/2024/winter) <br />Example usage: `?continuing`
	Continuing *bool  `queryParam:"style=form,explode=true,name=continuing"`
	Page       *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetSeasonRequest) GetYear() int64 {
	if o == nil {
		return 0
	}
	return o.Year
}

func (o *GetSeasonRequest) GetSeason() string {
	if o == nil {
		return ""
	}
	return o.Season
}

func (o *GetSeasonRequest) GetFilter() *GetSeasonQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetSeasonRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetSeasonRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetSeasonRequest) GetContinuing() *bool {
	if o == nil {
		return nil
	}
	return o.Continuing
}

func (o *GetSeasonRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetSeasonRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetSeasonResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns seasonal anime
	AnimeSearch *components.AnimeSearch
}

func (o *GetSeasonResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSeasonResponse) GetAnimeSearch() *components.AnimeSearch {
	if o == nil {
		return nil
	}
	return o.AnimeSearch
}
