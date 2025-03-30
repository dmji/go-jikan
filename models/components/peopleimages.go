package components

// PeopleImagesJpg - Available images in JPG
type PeopleImagesJpg struct {
	// Image URL JPG
	ImageURL *string `json:"image_url,omitempty"`
}

func (o *PeopleImagesJpg) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

type PeopleImages struct {
	// Available images in JPG
	Jpg *PeopleImagesJpg `json:"jpg,omitempty"`
}

func (o *PeopleImages) GetJpg() *PeopleImagesJpg {
	if o == nil {
		return nil
	}
	return o.Jpg
}
