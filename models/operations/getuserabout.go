package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserAboutRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserAboutRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserAboutResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user about in raw HTML
	UserAbout *components.UserAbout
}

func (o *GetUserAboutResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserAboutResponse) GetUserAbout() *components.UserAbout {
	if o == nil {
		return nil
	}
	return o.UserAbout
}
