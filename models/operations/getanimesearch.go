package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeSearchRequest struct {
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool   `queryParam:"style=form,explode=true,name=unapproved"`
	Page       *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q          *string `queryParam:"style=form,explode=true,name=q"`
	// Available Anime types
	Type  *components.AnimeSearchQueryType `queryParam:"style=form,explode=true,name=type"`
	Score *float64                         `queryParam:"style=form,explode=true,name=score"`
	// Set a minimum score for results.
	MinScore *float64 `queryParam:"style=form,explode=true,name=min_score"`
	// Set a maximum score for results
	MaxScore *float64 `queryParam:"style=form,explode=true,name=max_score"`
	// Available Anime statuses
	Status *components.AnimeSearchQueryStatus `queryParam:"style=form,explode=true,name=status"`
	// Available Anime audience ratings<br><br><b>Ratings</b><br><ul><li>G - All Ages</li><li>PG - Children</li><li>PG-13 - Teens 13 or older</li><li>R - 17+ (violence & profanity)</li><li>R+ - Mild Nudity</li><li>Rx - Hentai</li></ul>
	Rating *components.AnimeSearchQueryRating `queryParam:"style=form,explode=true,name=rating"`
	// Filter out Adult entries
	Sfw *bool `queryParam:"style=form,explode=true,name=sfw"`
	// Filter by genre(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	Genres *string `queryParam:"style=form,explode=true,name=genres"`
	// Exclude genre(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	GenresExclude *string `queryParam:"style=form,explode=true,name=genres_exclude"`
	// Available Anime order_by properties
	OrderBy *components.AnimeSearchQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
	// Filter by producer(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	Producers *string `queryParam:"style=form,explode=true,name=producers"`
	// Filter by starting date. Format: YYYY-MM-DD. e.g `2022`, `2005-05`, `2005-01-01`
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Filter by ending date. Format: YYYY-MM-DD. e.g `2022`, `2005-05`, `2005-01-01`
	EndDate *string `queryParam:"style=form,explode=true,name=end_date"`
}

func (o *GetAnimeSearchRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetAnimeSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetAnimeSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetAnimeSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetAnimeSearchRequest) GetType() *components.AnimeSearchQueryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetAnimeSearchRequest) GetScore() *float64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *GetAnimeSearchRequest) GetMinScore() *float64 {
	if o == nil {
		return nil
	}
	return o.MinScore
}

func (o *GetAnimeSearchRequest) GetMaxScore() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxScore
}

func (o *GetAnimeSearchRequest) GetStatus() *components.AnimeSearchQueryStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetAnimeSearchRequest) GetRating() *components.AnimeSearchQueryRating {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetAnimeSearchRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetAnimeSearchRequest) GetGenres() *string {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *GetAnimeSearchRequest) GetGenresExclude() *string {
	if o == nil {
		return nil
	}
	return o.GenresExclude
}

func (o *GetAnimeSearchRequest) GetOrderBy() *components.AnimeSearchQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetAnimeSearchRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetAnimeSearchRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

func (o *GetAnimeSearchRequest) GetProducers() *string {
	if o == nil {
		return nil
	}
	return o.Producers
}

func (o *GetAnimeSearchRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetAnimeSearchRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetAnimeSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for anime
	AnimeSearch *components.AnimeSearch
}

func (o *GetAnimeSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeSearchResponse) GetAnimeSearch() *components.AnimeSearch {
	if o == nil {
		return nil
	}
	return o.AnimeSearch
}
