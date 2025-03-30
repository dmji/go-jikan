package components

// MangaPictures - Manga Pictures
type MangaPictures struct {
	Data []MangaImages `json:"data,omitempty"`
}

func (o *MangaPictures) GetData() []MangaImages {
	if o == nil {
		return nil
	}
	return o.Data
}
