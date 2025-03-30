package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetCharacterMangaRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterMangaRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetCharacterMangaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga that character is in
	CharacterManga *components.CharacterManga
}

func (o *GetCharacterMangaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterMangaResponse) GetCharacterManga() *components.CharacterManga {
	if o == nil {
		return nil
	}
	return o.CharacterManga
}
