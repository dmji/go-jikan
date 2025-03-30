package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetCharacterAnimeRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterAnimeRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetCharacterAnimeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime that character is in
	CharacterAnime *components.CharacterAnime
}

func (o *GetCharacterAnimeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterAnimeResponse) GetCharacterAnime() *components.CharacterAnime {
	if o == nil {
		return nil
	}
	return o.CharacterAnime
}
