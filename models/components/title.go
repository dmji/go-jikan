package components

type Title struct {
	// Title type
	Type *string `json:"type,omitempty"`
	// Title value
	Title *string `json:"title,omitempty"`
}

func (o *Title) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Title) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
