package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserFullProfileRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserFullProfileRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// GetUserFullProfileResponseBody - Returns complete user resource data
type GetUserFullProfileResponseBody struct {
	// Transform the resource into an array.
	Data *components.UserProfileFull `json:"data,omitempty"`
}

func (o *GetUserFullProfileResponseBody) GetData() *components.UserProfileFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserFullProfileResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns complete user resource data
	Object *GetUserFullProfileResponseBody
}

func (o *GetUserFullProfileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserFullProfileResponse) GetObject() *GetUserFullProfileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
