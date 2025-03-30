package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeThemesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeThemesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeThemesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime themes
	AnimeThemes *components.AnimeThemes
}

func (o *GetAnimeThemesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeThemesResponse) GetAnimeThemes() *components.AnimeThemes {
	if o == nil {
		return nil
	}
	return o.AnimeThemes
}
