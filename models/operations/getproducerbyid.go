package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetProducerByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetProducerByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetProducerByIDResponseBody - Returns producer resource
type GetProducerByIDResponseBody struct {
	// Producers Resource
	Data *components.Producer `json:"data,omitempty"`
}

func (o *GetProducerByIDResponseBody) GetData() *components.Producer {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetProducerByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns producer resource
	Object *GetProducerByIDResponseBody
}

func (o *GetProducerByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProducerByIDResponse) GetObject() *GetProducerByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
