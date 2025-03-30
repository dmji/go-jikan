package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserUpdatesRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserUpdatesRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserUpdatesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user updates
	UserUpdates *components.UserUpdates
}

func (o *GetUserUpdatesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserUpdatesResponse) GetUserUpdates() *components.UserUpdates {
	if o == nil {
		return nil
	}
	return o.UserUpdates
}
