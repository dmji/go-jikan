package components

type PicturesVariantsData struct {
	Images *CommonImages `json:"images,omitempty"`
}

func (o *PicturesVariantsData) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

// PicturesVariants - Pictures Resource
type PicturesVariants struct {
	Data []PicturesVariantsData `json:"data,omitempty"`
}

func (o *PicturesVariants) GetData() []PicturesVariantsData {
	if o == nil {
		return nil
	}
	return o.Data
}
