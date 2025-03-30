package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeNewsRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAnimeNewsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeNewsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetAnimeNewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a list of news articles related to the entry
	AnimeNews *components.AnimeNews
}

func (o *GetAnimeNewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeNewsResponse) GetAnimeNews() *components.AnimeNews {
	if o == nil {
		return nil
	}
	return o.AnimeNews
}
