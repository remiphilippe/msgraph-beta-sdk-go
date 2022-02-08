package graph
import (
    "strings"
    "errors"
)
// 
type DeviceGuardVirtualizationBasedSecurityHardwareRequirementState int

const (
    MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = iota
    SECUREBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    DMAPROTECTIONREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    HYPERVNOTSUPPORTEDFORGUESTVM_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    HYPERVNOTAVAILABLE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
)

func (i DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) String() string {
    return []string{"MEETHARDWAREREQUIREMENTS", "SECUREBOOTREQUIRED", "DMAPROTECTIONREQUIRED", "HYPERVNOTSUPPORTEDFORGUESTVM", "HYPERVNOTAVAILABLE"}[i]
}
func ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(v string) (interface{}, error) {
    result := MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    switch strings.ToUpper(v) {
        case "MEETHARDWAREREQUIREMENTS":
            result = MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "SECUREBOOTREQUIRED":
            result = SECUREBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "DMAPROTECTIONREQUIRED":
            result = DMAPROTECTIONREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "HYPERVNOTSUPPORTEDFORGUESTVM":
            result = HYPERVNOTSUPPORTEDFORGUESTVM_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "HYPERVNOTAVAILABLE":
            result = HYPERVNOTAVAILABLE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        default:
            return 0, errors.New("Unknown DeviceGuardVirtualizationBasedSecurityHardwareRequirementState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(values []DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
