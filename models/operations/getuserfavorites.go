package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUserFavoritesRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserFavoritesRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

// GetUserFavoritesResponseBody - Returns user favorites
type GetUserFavoritesResponseBody struct {
	Data *components.UserFavorites `json:"data,omitempty"`
}

func (o *GetUserFavoritesResponseBody) GetData() *components.UserFavorites {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserFavoritesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user favorites
	Object *GetUserFavoritesResponseBody
}

func (o *GetUserFavoritesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserFavoritesResponse) GetObject() *GetUserFavoritesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
