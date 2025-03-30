package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserProfileRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserProfileRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// GetUserProfileResponseBody - Returns user profile
type GetUserProfileResponseBody struct {
	Data *components.UserProfile `json:"data,omitempty"`
}

func (o *GetUserProfileResponseBody) GetData() *components.UserProfile {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserProfileResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user profile
	Object *GetUserProfileResponseBody
}

func (o *GetUserProfileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserProfileResponse) GetObject() *GetUserProfileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
