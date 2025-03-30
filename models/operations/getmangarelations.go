package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMangaRelationsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaRelationsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetMangaRelationsResponseBody - Returns manga relations
type GetMangaRelationsResponseBody struct {
	Data []components.Relation `json:"data,omitempty"`
}

func (o *GetMangaRelationsResponseBody) GetData() []components.Relation {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetMangaRelationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga relations
	Object *GetMangaRelationsResponseBody
}

func (o *GetMangaRelationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaRelationsResponse) GetObject() *GetMangaRelationsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
