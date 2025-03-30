package components

type UserHistory struct {
	Data []History `json:"data,omitempty"`
}

func (o *UserHistory) GetData() []History {
	if o == nil {
		return nil
	}
	return o.Data
}
