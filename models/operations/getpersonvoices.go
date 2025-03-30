package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetPersonVoicesRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonVoicesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetPersonVoicesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns person's voice acting roles
	PersonVoiceActingRoles *components.PersonVoiceActingRoles
}

func (o *GetPersonVoicesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonVoicesResponse) GetPersonVoiceActingRoles() *components.PersonVoiceActingRoles {
	if o == nil {
		return nil
	}
	return o.PersonVoiceActingRoles
}
