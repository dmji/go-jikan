package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeCharactersRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeCharactersRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeCharactersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime characters resource
	AnimeCharacters *components.AnimeCharacters
}

func (o *GetAnimeCharactersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeCharactersResponse) GetAnimeCharacters() *components.AnimeCharacters {
	if o == nil {
		return nil
	}
	return o.AnimeCharacters
}
