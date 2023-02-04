package models
import (
    "errors"
    "strings"
)
// 
type ConditionalAccessConditions int

const (
    NONE_CONDITIONALACCESSCONDITIONS ConditionalAccessConditions = iota
    APPLICATION_CONDITIONALACCESSCONDITIONS
    USERS_CONDITIONALACCESSCONDITIONS
    DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS
    LOCATION_CONDITIONALACCESSCONDITIONS
    CLIENTTYPE_CONDITIONALACCESSCONDITIONS
    SIGNINRISK_CONDITIONALACCESSCONDITIONS
    USERRISK_CONDITIONALACCESSCONDITIONS
    TIME_CONDITIONALACCESSCONDITIONS
    DEVICESTATE_CONDITIONALACCESSCONDITIONS
    CLIENT_CONDITIONALACCESSCONDITIONS
    IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS
    IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS
    SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS
    SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS
)

func (i ConditionalAccessConditions) String() string {
    return []string{"none", "application", "users", "devicePlatform", "location", "clientType", "signInRisk", "userRisk", "time", "deviceState", "client", "ipAddressSeenByAzureAD", "ipAddressSeenByResourceProvider", "unknownFutureValue", "servicePrincipals", "servicePrincipalRisk"}[i]
}
func ParseConditionalAccessConditions(v string) (any, error) {
    var result []*ConditionalAccessConditions
    spl := strings.Split(v, ",")
    for _, s := range spl {
        var val ConditionalAccessConditions
        switch s {
            case "none":
                val = NONE_CONDITIONALACCESSCONDITIONS
            case "application":
                val = APPLICATION_CONDITIONALACCESSCONDITIONS
            case "users":
                val = USERS_CONDITIONALACCESSCONDITIONS
            case "devicePlatform":
                val = DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS
            case "location":
                val = LOCATION_CONDITIONALACCESSCONDITIONS
            case "clientType":
                val = CLIENTTYPE_CONDITIONALACCESSCONDITIONS
            case "signInRisk":
                val = SIGNINRISK_CONDITIONALACCESSCONDITIONS
            case "userRisk":
                val = USERRISK_CONDITIONALACCESSCONDITIONS
            case "time":
                val = TIME_CONDITIONALACCESSCONDITIONS
            case "deviceState":
                val = DEVICESTATE_CONDITIONALACCESSCONDITIONS
            case "client":
                val = CLIENT_CONDITIONALACCESSCONDITIONS
            case "ipAddressSeenByAzureAD":
                val = IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS
            case "ipAddressSeenByResourceProvider":
                val = IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS
            case "unknownFutureValue":
                val = UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS
            case "servicePrincipals":
                val = SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS
            case "servicePrincipalRisk":
                val = SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS
            default:
                return result, errors.New("Unknown ConditionalAccessConditions value: " + s)
        }
        result = append(result, &val)
    }
    if len(result) == 0 {
        val := NONE_CONDITIONALACCESSCONDITIONS
        result = append(result, &val)
    }
    return result, nil
}
func SerializeConditionalAccessConditions(values []ConditionalAccessConditions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
