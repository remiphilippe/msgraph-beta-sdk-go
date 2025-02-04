package models
import (
    "errors"
)
// 
type LayoutTemplateType int

const (
    DEFAULT_ESCAPED_LAYOUTTEMPLATETYPE LayoutTemplateType = iota
    VERTICALSPLIT_LAYOUTTEMPLATETYPE
    UNKNOWNFUTUREVALUE_LAYOUTTEMPLATETYPE
)

func (i LayoutTemplateType) String() string {
    return []string{"default", "verticalSplit", "unknownFutureValue"}[i]
}
func ParseLayoutTemplateType(v string) (any, error) {
    result := DEFAULT_ESCAPED_LAYOUTTEMPLATETYPE
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_LAYOUTTEMPLATETYPE
        case "verticalSplit":
            result = VERTICALSPLIT_LAYOUTTEMPLATETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LAYOUTTEMPLATETYPE
        default:
            return 0, errors.New("Unknown LayoutTemplateType value: " + v)
    }
    return &result, nil
}
func SerializeLayoutTemplateType(values []LayoutTemplateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
