package operations

import (
	"github.com/dmji/jikan/models/components"
)

// GetRandomUsersResponseBody - Returns a random user profile resource
type GetRandomUsersResponseBody struct {
	Data *components.UserProfile `json:"data,omitempty"`
}

func (o *GetRandomUsersResponseBody) GetData() *components.UserProfile {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRandomUsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a random user profile resource
	Object *GetRandomUsersResponseBody
}

func (o *GetRandomUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRandomUsersResponse) GetObject() *GetRandomUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
