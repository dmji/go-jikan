package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetPersonMangaRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonMangaRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetPersonMangaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns person's published manga works
	PersonManga *components.PersonManga
}

func (o *GetPersonMangaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonMangaResponse) GetPersonManga() *components.PersonManga {
	if o == nil {
		return nil
	}
	return o.PersonManga
}
