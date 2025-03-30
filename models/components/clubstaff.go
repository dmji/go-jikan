package components

type ClubStaffData struct {
	// User URL
	URL *string `json:"url,omitempty"`
	// User's username
	Username *string `json:"username,omitempty"`
}

func (o *ClubStaffData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *ClubStaffData) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// ClubStaff - Club Staff Resource
type ClubStaff struct {
	Data []ClubStaffData `json:"data,omitempty"`
}

func (o *ClubStaff) GetData() []ClubStaffData {
	if o == nil {
		return nil
	}
	return o.Data
}
