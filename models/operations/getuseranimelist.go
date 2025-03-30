package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserAnimelistRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	// User's anime list status filter options
	Status *components.UserAnimeListStatusFilter `queryParam:"style=form,explode=true,name=status"`
}

func (o *GetUserAnimelistRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserAnimelistRequest) GetStatus() *components.UserAnimeListStatusFilter {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetUserAnimelistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user anime list
	Any any
}

func (o *GetUserAnimelistResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserAnimelistResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}
