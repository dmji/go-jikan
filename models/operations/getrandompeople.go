package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

// GetRandomPeopleResponseBody - Returns a random person resource
type GetRandomPeopleResponseBody struct {
	// Person Resource
	Data *components.Person `json:"data,omitempty"`
}

func (o *GetRandomPeopleResponseBody) GetData() *components.Person {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRandomPeopleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a random person resource
	Object *GetRandomPeopleResponseBody
}

func (o *GetRandomPeopleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRandomPeopleResponse) GetObject() *GetRandomPeopleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
