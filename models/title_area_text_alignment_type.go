package models
import (
    "errors"
)
// 
type TitleAreaTextAlignmentType int

const (
    LEFT_TITLEAREATEXTALIGNMENTTYPE TitleAreaTextAlignmentType = iota
    CENTER_TITLEAREATEXTALIGNMENTTYPE
    UNKNOWNFUTUREVALUE_TITLEAREATEXTALIGNMENTTYPE
)

func (i TitleAreaTextAlignmentType) String() string {
    return []string{"left", "center", "unknownFutureValue"}[i]
}
func ParseTitleAreaTextAlignmentType(v string) (any, error) {
    result := LEFT_TITLEAREATEXTALIGNMENTTYPE
    switch v {
        case "left":
            result = LEFT_TITLEAREATEXTALIGNMENTTYPE
        case "center":
            result = CENTER_TITLEAREATEXTALIGNMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TITLEAREATEXTALIGNMENTTYPE
        default:
            return 0, errors.New("Unknown TitleAreaTextAlignmentType value: " + v)
    }
    return &result, nil
}
func SerializeTitleAreaTextAlignmentType(values []TitleAreaTextAlignmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
