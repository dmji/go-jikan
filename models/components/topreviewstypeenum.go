package components

import (
	"encoding/json"
	"fmt"
)

// TopReviewsTypeEnum - The type of reviews to filter by. Defaults to anime.
type TopReviewsTypeEnum string

const (
	TopReviewsTypeEnumAnime TopReviewsTypeEnum = "anime"
	TopReviewsTypeEnumManga TopReviewsTypeEnum = "manga"
)

func (e TopReviewsTypeEnum) ToPointer() *TopReviewsTypeEnum {
	return &e
}

func (e *TopReviewsTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "anime":
		fallthrough
	case "manga":
		*e = TopReviewsTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TopReviewsTypeEnum: %v", v)
	}
}
