package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type AccessPackageCustomExtensionStage int

const (
    ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE AccessPackageCustomExtensionStage = iota
    ASSIGNMENTREQUESTAPPROVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTREQUESTGRANTED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTREQUESTREMOVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTONEDAYBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
)

func (i AccessPackageCustomExtensionStage) String() string {
    return []string{"assignmentRequestCreated", "assignmentRequestApproved", "assignmentRequestGranted", "assignmentRequestRemoved", "assignmentFourteenDaysBeforeExpiration", "assignmentOneDayBeforeExpiration", "unknownFutureValue"}[i]
}
func ParseAccessPackageCustomExtensionStage(v string) (interface{}, error) {
    result := ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    switch v {
        case "assignmentRequestCreated":
            result = ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "assignmentRequestApproved":
            result = ASSIGNMENTREQUESTAPPROVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "assignmentRequestGranted":
            result = ASSIGNMENTREQUESTGRANTED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "assignmentRequestRemoved":
            result = ASSIGNMENTREQUESTREMOVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "assignmentFourteenDaysBeforeExpiration":
            result = ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "assignmentOneDayBeforeExpiration":
            result = ASSIGNMENTONEDAYBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        default:
            return 0, errors.New("Unknown AccessPackageCustomExtensionStage value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCustomExtensionStage(values []AccessPackageCustomExtensionStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
