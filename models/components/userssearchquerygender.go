package components

import (
	"encoding/json"
	"fmt"
)

// UsersSearchQueryGender - Users Search Query Gender.
type UsersSearchQueryGender string

const (
	UsersSearchQueryGenderAny       UsersSearchQueryGender = "any"
	UsersSearchQueryGenderMale      UsersSearchQueryGender = "male"
	UsersSearchQueryGenderFemale    UsersSearchQueryGender = "female"
	UsersSearchQueryGenderNonbinary UsersSearchQueryGender = "nonbinary"
)

func (e UsersSearchQueryGender) ToPointer() *UsersSearchQueryGender {
	return &e
}

func (e *UsersSearchQueryGender) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "male":
		fallthrough
	case "female":
		fallthrough
	case "nonbinary":
		*e = UsersSearchQueryGender(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UsersSearchQueryGender: %v", v)
	}
}
