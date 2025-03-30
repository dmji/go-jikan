package components

type CharacterPicturesData struct {
	// Default JPG Image Size URL
	ImageURL *string `json:"image_url,omitempty"`
	// Large JPG Image Size URL
	LargeImageURL *string `json:"large_image_url,omitempty"`
}

func (o *CharacterPicturesData) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *CharacterPicturesData) GetLargeImageURL() *string {
	if o == nil {
		return nil
	}
	return o.LargeImageURL
}

// CharacterPictures - Character Pictures
type CharacterPictures struct {
	Data []CharacterPicturesData `json:"data,omitempty"`
}

func (o *CharacterPictures) GetData() []CharacterPicturesData {
	if o == nil {
		return nil
	}
	return o.Data
}
