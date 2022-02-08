package graph
import (
    "strings"
    "errors"
)
// 
type RequirementProvider int

const (
    USER_REQUIREMENTPROVIDER RequirementProvider = iota
    REQUEST_REQUIREMENTPROVIDER
    SERVICEPRINCIPAL_REQUIREMENTPROVIDER
    V1CONDITIONALACCESS_REQUIREMENTPROVIDER
    MULTICONDITIONALACCESS_REQUIREMENTPROVIDER
    TENANTSESSIONRISKPOLICY_REQUIREMENTPROVIDER
    ACCOUNTCOMPROMISEPOLICIES_REQUIREMENTPROVIDER
    V1CONDITIONALACCESSDEPENDENCY_REQUIREMENTPROVIDER
    V1CONDITIONALACCESSPOLICYIDREQUESTED_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY_REQUIREMENTPROVIDER
    BASELINEPROTECTION_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS_REQUIREMENTPROVIDER
    ENFORCEDFORCSPADMINS_REQUIREMENTPROVIDER
    SECURITYDEFAULTS_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS_REQUIREMENTPROVIDER
    PROOFUPCODEREQUEST_REQUIREMENTPROVIDER
    CROSSTENANTOUTBOUNDRULE_REQUIREMENTPROVIDER
    GPSLOCATIONCONDITION_REQUIREMENTPROVIDER
    RISKBASEDPOLICY_REQUIREMENTPROVIDER
    UNKNOWNFUTUREVALUE_REQUIREMENTPROVIDER
)

func (i RequirementProvider) String() string {
    return []string{"USER", "REQUEST", "SERVICEPRINCIPAL", "V1CONDITIONALACCESS", "MULTICONDITIONALACCESS", "TENANTSESSIONRISKPOLICY", "ACCOUNTCOMPROMISEPOLICIES", "V1CONDITIONALACCESSDEPENDENCY", "V1CONDITIONALACCESSPOLICYIDREQUESTED", "MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY", "BASELINEPROTECTION", "MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION", "MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS", "ENFORCEDFORCSPADMINS", "SECURITYDEFAULTS", "MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS", "PROOFUPCODEREQUEST", "CROSSTENANTOUTBOUNDRULE", "GPSLOCATIONCONDITION", "RISKBASEDPOLICY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRequirementProvider(v string) (interface{}, error) {
    result := USER_REQUIREMENTPROVIDER
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_REQUIREMENTPROVIDER
        case "REQUEST":
            result = REQUEST_REQUIREMENTPROVIDER
        case "SERVICEPRINCIPAL":
            result = SERVICEPRINCIPAL_REQUIREMENTPROVIDER
        case "V1CONDITIONALACCESS":
            result = V1CONDITIONALACCESS_REQUIREMENTPROVIDER
        case "MULTICONDITIONALACCESS":
            result = MULTICONDITIONALACCESS_REQUIREMENTPROVIDER
        case "TENANTSESSIONRISKPOLICY":
            result = TENANTSESSIONRISKPOLICY_REQUIREMENTPROVIDER
        case "ACCOUNTCOMPROMISEPOLICIES":
            result = ACCOUNTCOMPROMISEPOLICIES_REQUIREMENTPROVIDER
        case "V1CONDITIONALACCESSDEPENDENCY":
            result = V1CONDITIONALACCESSDEPENDENCY_REQUIREMENTPROVIDER
        case "V1CONDITIONALACCESSPOLICYIDREQUESTED":
            result = V1CONDITIONALACCESSPOLICYIDREQUESTED_REQUIREMENTPROVIDER
        case "MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY":
            result = MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY_REQUIREMENTPROVIDER
        case "BASELINEPROTECTION":
            result = BASELINEPROTECTION_REQUIREMENTPROVIDER
        case "MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION":
            result = MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION_REQUIREMENTPROVIDER
        case "MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS":
            result = MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS_REQUIREMENTPROVIDER
        case "ENFORCEDFORCSPADMINS":
            result = ENFORCEDFORCSPADMINS_REQUIREMENTPROVIDER
        case "SECURITYDEFAULTS":
            result = SECURITYDEFAULTS_REQUIREMENTPROVIDER
        case "MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS":
            result = MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS_REQUIREMENTPROVIDER
        case "PROOFUPCODEREQUEST":
            result = PROOFUPCODEREQUEST_REQUIREMENTPROVIDER
        case "CROSSTENANTOUTBOUNDRULE":
            result = CROSSTENANTOUTBOUNDRULE_REQUIREMENTPROVIDER
        case "GPSLOCATIONCONDITION":
            result = GPSLOCATIONCONDITION_REQUIREMENTPROVIDER
        case "RISKBASEDPOLICY":
            result = RISKBASEDPOLICY_REQUIREMENTPROVIDER
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REQUIREMENTPROVIDER
        default:
            return 0, errors.New("Unknown RequirementProvider value: " + v)
    }
    return &result, nil
}
func SerializeRequirementProvider(values []RequirementProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
