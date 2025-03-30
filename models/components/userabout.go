package components

type UserAboutData struct {
	// User About. NOTE: About information is customizable by users through BBCode on MyAnimeList. This means users can add multimedia content, different text sizes, etc. Due to this freeform, Jikan returns parsed HTML. Validate on your end!
	About *string `json:"about,omitempty"`
}

func (o *UserAboutData) GetAbout() *string {
	if o == nil {
		return nil
	}
	return o.About
}

type UserAbout struct {
	Data []UserAboutData `json:"data,omitempty"`
}

func (o *UserAbout) GetData() []UserAboutData {
	if o == nil {
		return nil
	}
	return o.Data
}
