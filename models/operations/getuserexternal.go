package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserExternalRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserExternalRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserExternalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user's external links
	ExternalLinks *components.ExternalLinks
}

func (o *GetUserExternalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserExternalResponse) GetExternalLinks() *components.ExternalLinks {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}
