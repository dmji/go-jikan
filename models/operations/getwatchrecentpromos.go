package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetWatchRecentPromosRequest struct {
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetWatchRecentPromosRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetWatchRecentPromosResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Recently Added Promotional Videos
	WatchPromos *components.WatchPromos
}

func (o *GetWatchRecentPromosResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWatchRecentPromosResponse) GetWatchPromos() *components.WatchPromos {
	if o == nil {
		return nil
	}
	return o.WatchPromos
}
