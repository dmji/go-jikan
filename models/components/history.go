package components

// History - Transform the resource into an array.
type History struct {
	// Parsed URL Data
	Entry *MalURL `json:"entry,omitempty"`
	// Number of episodes/chapters watched/read
	Increment *int64 `json:"increment,omitempty"`
	// Date ISO8601
	Date *string `json:"date,omitempty"`
}

func (o *History) GetEntry() *MalURL {
	if o == nil {
		return nil
	}
	return o.Entry
}

func (o *History) GetIncrement() *int64 {
	if o == nil {
		return nil
	}
	return o.Increment
}

func (o *History) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}
