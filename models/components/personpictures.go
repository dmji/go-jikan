package components

// PersonPictures - Character Pictures
type PersonPictures struct {
	Data []PeopleImages `json:"data,omitempty"`
}

func (o *PersonPictures) GetData() []PeopleImages {
	if o == nil {
		return nil
	}
	return o.Data
}
