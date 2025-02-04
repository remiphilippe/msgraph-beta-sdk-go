package models
import (
    "errors"
)
// 
type PrivilegedAccessGroupMemberType int

const (
    DIRECT_PRIVILEGEDACCESSGROUPMEMBERTYPE PrivilegedAccessGroupMemberType = iota
    GROUP_PRIVILEGEDACCESSGROUPMEMBERTYPE
    UNKNOWNFUTUREVALUE_PRIVILEGEDACCESSGROUPMEMBERTYPE
)

func (i PrivilegedAccessGroupMemberType) String() string {
    return []string{"direct", "group", "unknownFutureValue"}[i]
}
func ParsePrivilegedAccessGroupMemberType(v string) (any, error) {
    result := DIRECT_PRIVILEGEDACCESSGROUPMEMBERTYPE
    switch v {
        case "direct":
            result = DIRECT_PRIVILEGEDACCESSGROUPMEMBERTYPE
        case "group":
            result = GROUP_PRIVILEGEDACCESSGROUPMEMBERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVILEGEDACCESSGROUPMEMBERTYPE
        default:
            return 0, errors.New("Unknown PrivilegedAccessGroupMemberType value: " + v)
    }
    return &result, nil
}
func SerializePrivilegedAccessGroupMemberType(values []PrivilegedAccessGroupMemberType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
