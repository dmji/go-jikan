package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetTopAnimeRequest struct {
	// Available Anime types
	Type *components.AnimeSearchQueryType `queryParam:"style=form,explode=true,name=type"`
	// Top items filter types
	Filter *components.TopAnimeFilter `queryParam:"style=form,explode=true,name=filter"`
	// Available Anime audience ratings<br><br><b>Ratings</b><br><ul><li>G - All Ages</li><li>PG - Children</li><li>PG-13 - Teens 13 or older</li><li>R - 17+ (violence & profanity)</li><li>R+ - Mild Nudity</li><li>Rx - Hentai</li></ul>
	Rating *components.AnimeSearchQueryRating `queryParam:"style=form,explode=true,name=rating"`
	// Filter out Adult entries
	Sfw   *bool  `queryParam:"style=form,explode=true,name=sfw"`
	Page  *int64 `queryParam:"style=form,explode=true,name=page"`
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
}

func (o *GetTopAnimeRequest) GetType() *components.AnimeSearchQueryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetTopAnimeRequest) GetFilter() *components.TopAnimeFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetTopAnimeRequest) GetRating() *components.AnimeSearchQueryRating {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetTopAnimeRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetTopAnimeRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetTopAnimeRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetTopAnimeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns top anime
	AnimeSearch *components.AnimeSearch
}

func (o *GetTopAnimeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopAnimeResponse) GetAnimeSearch() *components.AnimeSearch {
	if o == nil {
		return nil
	}
	return o.AnimeSearch
}
