package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CloudPcDeviceable 
type CloudPcDeviceable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCloudPcStatus()(*string)
    GetDisplayName()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetProvisioningPolicyId()(*string)
    GetServicePlanName()(*string)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetUserPrincipalName()(*string)
    SetCloudPcStatus(value *string)()
    SetDisplayName(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetProvisioningPolicyId(value *string)()
    SetServicePlanName(value *string)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetUserPrincipalName(value *string)()
}
