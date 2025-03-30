package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeGenresRequest struct {
	// Filter genres by type
	Filter *components.GenreQueryFilter `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetAnimeGenresRequest) GetFilter() *components.GenreQueryFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetAnimeGenresResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns entry genres, explicit_genres, themes and demographics
	Genres *components.Genres
}

func (o *GetAnimeGenresResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeGenresResponse) GetGenres() *components.Genres {
	if o == nil {
		return nil
	}
	return o.Genres
}
