package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetCharacterVoiceActorsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterVoiceActorsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetCharacterVoiceActorsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns the character's voice actors
	CharacterVoiceActors *components.CharacterVoiceActors
}

func (o *GetCharacterVoiceActorsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterVoiceActorsResponse) GetCharacterVoiceActors() *components.CharacterVoiceActors {
	if o == nil {
		return nil
	}
	return o.CharacterVoiceActors
}
