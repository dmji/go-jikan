package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserClubsRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	Page     *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetUserClubsRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserClubsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetUserClubsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user clubs
	UserClubs *components.UserClubs
}

func (o *GetUserClubsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserClubsResponse) GetUserClubs() *components.UserClubs {
	if o == nil {
		return nil
	}
	return o.UserClubs
}
