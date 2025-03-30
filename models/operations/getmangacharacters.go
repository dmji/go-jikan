package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaCharactersRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaCharactersRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaCharactersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga characters resource
	MangaCharacters *components.MangaCharacters
}

func (o *GetMangaCharactersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaCharactersResponse) GetMangaCharacters() *components.MangaCharacters {
	if o == nil {
		return nil
	}
	return o.MangaCharacters
}
