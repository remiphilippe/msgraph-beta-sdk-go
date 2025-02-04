package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBodyable 
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressableUserName()(*string)
    GetDeviceAccountPassword()(*string)
    GetDeviceAccountUpn()(*string)
    GetDeviceFriendlyName()(*string)
    GetDisplayName()(*string)
    GetGroupTag()(*string)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetDeviceAccountPassword(value *string)()
    SetDeviceAccountUpn(value *string)()
    SetDeviceFriendlyName(value *string)()
    SetDisplayName(value *string)()
    SetGroupTag(value *string)()
    SetUserPrincipalName(value *string)()
}
