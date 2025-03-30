package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetCharacterPicturesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterPicturesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetCharacterPicturesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns pictures related to the entry
	CharacterPictures *components.CharacterPictures
}

func (o *GetCharacterPicturesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterPicturesResponse) GetCharacterPictures() *components.CharacterPictures {
	if o == nil {
		return nil
	}
	return o.CharacterPictures
}
