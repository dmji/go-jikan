package components

// Genres Collection Resource
type Genres struct {
	Data []Genre `json:"data,omitempty"`
}

func (o *Genres) GetData() []Genre {
	if o == nil {
		return nil
	}
	return o.Data
}
