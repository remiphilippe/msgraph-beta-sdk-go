package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the compliance singleton.
type PrintMultipageLayout int

const (
    CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT PrintMultipageLayout = iota
    COUNTERCLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
    COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
    CLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
)

func (i PrintMultipageLayout) String() string {
    return []string{"CLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPLEFT", "COUNTERCLOCKWISEFROMTOPRIGHT", "CLOCKWISEFROMTOPRIGHT", "COUNTERCLOCKWISEFROMBOTTOMLEFT", "CLOCKWISEFROMBOTTOMLEFT", "COUNTERCLOCKWISEFROMBOTTOMRIGHT", "CLOCKWISEFROMBOTTOMRIGHT"}[i]
}
func ParsePrintMultipageLayout(v string) (interface{}, error) {
    result := CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
    switch strings.ToUpper(v) {
        case "CLOCKWISEFROMTOPLEFT":
            result = CLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
        case "COUNTERCLOCKWISEFROMTOPLEFT":
            result = COUNTERCLOCKWISEFROMTOPLEFT_PRINTMULTIPAGELAYOUT
        case "COUNTERCLOCKWISEFROMTOPRIGHT":
            result = COUNTERCLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
        case "CLOCKWISEFROMTOPRIGHT":
            result = CLOCKWISEFROMTOPRIGHT_PRINTMULTIPAGELAYOUT
        case "COUNTERCLOCKWISEFROMBOTTOMLEFT":
            result = COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
        case "CLOCKWISEFROMBOTTOMLEFT":
            result = CLOCKWISEFROMBOTTOMLEFT_PRINTMULTIPAGELAYOUT
        case "COUNTERCLOCKWISEFROMBOTTOMRIGHT":
            result = COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
        case "CLOCKWISEFROMBOTTOMRIGHT":
            result = CLOCKWISEFROMBOTTOMRIGHT_PRINTMULTIPAGELAYOUT
        default:
            return 0, errors.New("Unknown PrintMultipageLayout value: " + v)
    }
    return &result, nil
}
func SerializePrintMultipageLayout(values []PrintMultipageLayout) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
