package models
import (
    "errors"
)
// 
type PlannerTaskTargetKind int

const (
    GROUP_PLANNERTASKTARGETKIND PlannerTaskTargetKind = iota
    UNKNOWNFUTUREVALUE_PLANNERTASKTARGETKIND
)

func (i PlannerTaskTargetKind) String() string {
    return []string{"group", "unknownFutureValue"}[i]
}
func ParsePlannerTaskTargetKind(v string) (any, error) {
    result := GROUP_PLANNERTASKTARGETKIND
    switch v {
        case "group":
            result = GROUP_PLANNERTASKTARGETKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERTASKTARGETKIND
        default:
            return 0, errors.New("Unknown PlannerTaskTargetKind value: " + v)
    }
    return &result, nil
}
func SerializePlannerTaskTargetKind(values []PlannerTaskTargetKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
