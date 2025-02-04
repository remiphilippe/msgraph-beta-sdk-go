package models
import (
    "errors"
)
// 
type IdentityProviderState int

const (
    ENABLED_IDENTITYPROVIDERSTATE IdentityProviderState = iota
    DISABLED_IDENTITYPROVIDERSTATE
    UNKNOWNFUTUREVALUE_IDENTITYPROVIDERSTATE
)

func (i IdentityProviderState) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseIdentityProviderState(v string) (any, error) {
    result := ENABLED_IDENTITYPROVIDERSTATE
    switch v {
        case "enabled":
            result = ENABLED_IDENTITYPROVIDERSTATE
        case "disabled":
            result = DISABLED_IDENTITYPROVIDERSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IDENTITYPROVIDERSTATE
        default:
            return 0, errors.New("Unknown IdentityProviderState value: " + v)
    }
    return &result, nil
}
func SerializeIdentityProviderState(values []IdentityProviderState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
