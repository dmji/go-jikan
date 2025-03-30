package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetWatchPopularPromosResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Popular Promotional Videos
	WatchPromos *components.WatchPromos
}

func (o *GetWatchPopularPromosResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWatchPopularPromosResponse) GetWatchPromos() *components.WatchPromos {
	if o == nil {
		return nil
	}
	return o.WatchPromos
}
