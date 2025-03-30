package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetPersonAnimeRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonAnimeRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetPersonAnimeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns person's anime staff positions
	PersonAnime *components.PersonAnime
}

func (o *GetPersonAnimeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonAnimeResponse) GetPersonAnime() *components.PersonAnime {
	if o == nil {
		return nil
	}
	return o.PersonAnime
}
