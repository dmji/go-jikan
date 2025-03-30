package components

type MoreinfoData struct {
	// Additional information on the entry
	Moreinfo *string `json:"moreinfo,omitempty"`
}

func (o *MoreinfoData) GetMoreinfo() *string {
	if o == nil {
		return nil
	}
	return o.Moreinfo
}

// Moreinfo - More Info Resource
type Moreinfo struct {
	Data *MoreinfoData `json:"data,omitempty"`
}

func (o *Moreinfo) GetData() *MoreinfoData {
	if o == nil {
		return nil
	}
	return o.Data
}
