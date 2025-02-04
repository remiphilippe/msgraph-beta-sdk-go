package models
import (
    "errors"
)
// 
type ShiftWorkCloudPcAccessState int

const (
    UNASSIGNED_SHIFTWORKCLOUDPCACCESSSTATE ShiftWorkCloudPcAccessState = iota
    NOLICENSESAVAILABLE_SHIFTWORKCLOUDPCACCESSSTATE
    ACTIVATIONFAILED_SHIFTWORKCLOUDPCACCESSSTATE
    ACTIVE_SHIFTWORKCLOUDPCACCESSSTATE
    ACTIVATING_SHIFTWORKCLOUDPCACCESSSTATE
    WAITLISTED_SHIFTWORKCLOUDPCACCESSSTATE
    UNKNOWNFUTUREVALUE_SHIFTWORKCLOUDPCACCESSSTATE
)

func (i ShiftWorkCloudPcAccessState) String() string {
    return []string{"unassigned", "noLicensesAvailable", "activationFailed", "active", "activating", "waitlisted", "unknownFutureValue"}[i]
}
func ParseShiftWorkCloudPcAccessState(v string) (any, error) {
    result := UNASSIGNED_SHIFTWORKCLOUDPCACCESSSTATE
    switch v {
        case "unassigned":
            result = UNASSIGNED_SHIFTWORKCLOUDPCACCESSSTATE
        case "noLicensesAvailable":
            result = NOLICENSESAVAILABLE_SHIFTWORKCLOUDPCACCESSSTATE
        case "activationFailed":
            result = ACTIVATIONFAILED_SHIFTWORKCLOUDPCACCESSSTATE
        case "active":
            result = ACTIVE_SHIFTWORKCLOUDPCACCESSSTATE
        case "activating":
            result = ACTIVATING_SHIFTWORKCLOUDPCACCESSSTATE
        case "waitlisted":
            result = WAITLISTED_SHIFTWORKCLOUDPCACCESSSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SHIFTWORKCLOUDPCACCESSSTATE
        default:
            return 0, errors.New("Unknown ShiftWorkCloudPcAccessState value: " + v)
    }
    return &result, nil
}
func SerializeShiftWorkCloudPcAccessState(values []ShiftWorkCloudPcAccessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
