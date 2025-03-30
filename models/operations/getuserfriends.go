package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserFriendsRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	Page     *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetUserFriendsRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserFriendsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetUserFriendsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user friends
	UserFriends *components.UserFriends
}

func (o *GetUserFriendsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserFriendsResponse) GetUserFriends() *components.UserFriends {
	if o == nil {
		return nil
	}
	return o.UserFriends
}
