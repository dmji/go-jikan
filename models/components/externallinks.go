package components

type ExternalLinksData struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

func (o *ExternalLinksData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ExternalLinksData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// ExternalLinks - External links
type ExternalLinks struct {
	Data []ExternalLinksData `json:"data,omitempty"`
}

func (o *ExternalLinks) GetData() []ExternalLinksData {
	if o == nil {
		return nil
	}
	return o.Data
}
