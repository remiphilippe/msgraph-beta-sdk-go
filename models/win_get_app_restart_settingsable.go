package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetAppRestartSettingsable 
type WinGetAppRestartSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountdownDisplayBeforeRestartInMinutes()(*int32)
    GetGracePeriodInMinutes()(*int32)
    GetOdataType()(*string)
    GetRestartNotificationSnoozeDurationInMinutes()(*int32)
    SetCountdownDisplayBeforeRestartInMinutes(value *int32)()
    SetGracePeriodInMinutes(value *int32)()
    SetOdataType(value *string)()
    SetRestartNotificationSnoozeDurationInMinutes(value *int32)()
}
