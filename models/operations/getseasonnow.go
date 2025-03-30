package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-jikan/models/components"
)

// GetSeasonNowQueryParamFilter - Entry types
type GetSeasonNowQueryParamFilter string

const (
	GetSeasonNowQueryParamFilterTv      GetSeasonNowQueryParamFilter = "tv"
	GetSeasonNowQueryParamFilterMovie   GetSeasonNowQueryParamFilter = "movie"
	GetSeasonNowQueryParamFilterOva     GetSeasonNowQueryParamFilter = "ova"
	GetSeasonNowQueryParamFilterSpecial GetSeasonNowQueryParamFilter = "special"
	GetSeasonNowQueryParamFilterOna     GetSeasonNowQueryParamFilter = "ona"
	GetSeasonNowQueryParamFilterMusic   GetSeasonNowQueryParamFilter = "music"
)

func (e GetSeasonNowQueryParamFilter) ToPointer() *GetSeasonNowQueryParamFilter {
	return &e
}

func (e *GetSeasonNowQueryParamFilter) UnmarshalJSON(data []byte) error {
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
		*e = GetSeasonNowQueryParamFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSeasonNowQueryParamFilter: %v", v)
	}
}

type GetSeasonNowRequest struct {
	// Entry types
	Filter *GetSeasonNowQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// 'Safe For Work'. This is a flag. When supplied it will filter out entries according to the SFW Policy. You do not need to pass a value to it. e.g usage: `?sfw`
	Sfw *bool `queryParam:"style=form,explode=true,name=sfw"`
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool `queryParam:"style=form,explode=true,name=unapproved"`
	// This is a flag. When supplied it will include entries which are continuing from previous seasons. MAL includes these items on the seasons view in the &#8243;TV (continuing)&#8243; section. (Example: https://myanimelist.net/anime/season/2024/winter) <br />Example usage: `?continuing`
	Continuing *bool  `queryParam:"style=form,explode=true,name=continuing"`
	Page       *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetSeasonNowRequest) GetFilter() *GetSeasonNowQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetSeasonNowRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetSeasonNowRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetSeasonNowRequest) GetContinuing() *bool {
	if o == nil {
		return nil
	}
	return o.Continuing
}

func (o *GetSeasonNowRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetSeasonNowRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetSeasonNowResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns current seasonal anime
	AnimeSearch *components.AnimeSearch
}

func (o *GetSeasonNowResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSeasonNowResponse) GetAnimeSearch() *components.AnimeSearch {
	if o == nil {
		return nil
	}
	return o.AnimeSearch
}
