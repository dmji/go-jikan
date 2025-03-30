package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMangaSearchRequest struct {
	// This is a flag. When supplied it will include entries which are unapproved. Unapproved entries on MyAnimeList are those that are user submitted and have not yet been approved by MAL to show up on other pages. They will have their own specifc pages and are often removed resulting in a 404 error. You do not need to pass a value to it. e.g usage: `?unapproved`
	Unapproved *bool   `queryParam:"style=form,explode=true,name=unapproved"`
	Page       *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit      *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q          *string `queryParam:"style=form,explode=true,name=q"`
	// Available Manga types
	Type  *components.MangaSearchQueryType `queryParam:"style=form,explode=true,name=type"`
	Score *float64                         `queryParam:"style=form,explode=true,name=score"`
	// Set a minimum score for results.
	MinScore *float64 `queryParam:"style=form,explode=true,name=min_score"`
	// Set a maximum score for results
	MaxScore *float64 `queryParam:"style=form,explode=true,name=max_score"`
	// Available Manga statuses
	Status *components.MangaSearchQueryStatus `queryParam:"style=form,explode=true,name=status"`
	// Filter out Adult entries
	Sfw *bool `queryParam:"style=form,explode=true,name=sfw"`
	// Filter by genre(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	Genres *string `queryParam:"style=form,explode=true,name=genres"`
	// Exclude genre(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	GenresExclude *string `queryParam:"style=form,explode=true,name=genres_exclude"`
	// Available Manga order_by properties
	OrderBy *components.MangaSearchQueryOrderby `queryParam:"style=form,explode=true,name=order_by"`
	// Search query sort direction
	Sort *components.SearchQuerySort `queryParam:"style=form,explode=true,name=sort"`
	// Return entries starting with the given letter
	Letter *string `queryParam:"style=form,explode=true,name=letter"`
	// Filter by magazine(s) IDs. Can pass multiple with a comma as a delimiter. e.g 1,2,3
	Magazines *string `queryParam:"style=form,explode=true,name=magazines"`
	// Filter by starting date. Format: YYYY-MM-DD. e.g `2022`, `2005-05`, `2005-01-01`
	StartDate *string `queryParam:"style=form,explode=true,name=start_date"`
	// Filter by ending date. Format: YYYY-MM-DD. e.g `2022`, `2005-05`, `2005-01-01`
	EndDate *string `queryParam:"style=form,explode=true,name=end_date"`
}

func (o *GetMangaSearchRequest) GetUnapproved() *bool {
	if o == nil {
		return nil
	}
	return o.Unapproved
}

func (o *GetMangaSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetMangaSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetMangaSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetMangaSearchRequest) GetType() *components.MangaSearchQueryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetMangaSearchRequest) GetScore() *float64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *GetMangaSearchRequest) GetMinScore() *float64 {
	if o == nil {
		return nil
	}
	return o.MinScore
}

func (o *GetMangaSearchRequest) GetMaxScore() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxScore
}

func (o *GetMangaSearchRequest) GetStatus() *components.MangaSearchQueryStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetMangaSearchRequest) GetSfw() *bool {
	if o == nil {
		return nil
	}
	return o.Sfw
}

func (o *GetMangaSearchRequest) GetGenres() *string {
	if o == nil {
		return nil
	}
	return o.Genres
}

func (o *GetMangaSearchRequest) GetGenresExclude() *string {
	if o == nil {
		return nil
	}
	return o.GenresExclude
}

func (o *GetMangaSearchRequest) GetOrderBy() *components.MangaSearchQueryOrderby {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetMangaSearchRequest) GetSort() *components.SearchQuerySort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetMangaSearchRequest) GetLetter() *string {
	if o == nil {
		return nil
	}
	return o.Letter
}

func (o *GetMangaSearchRequest) GetMagazines() *string {
	if o == nil {
		return nil
	}
	return o.Magazines
}

func (o *GetMangaSearchRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetMangaSearchRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetMangaSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for manga
	MangaSearch *components.MangaSearch
}

func (o *GetMangaSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaSearchResponse) GetMangaSearch() *components.MangaSearch {
	if o == nil {
		return nil
	}
	return o.MangaSearch
}
