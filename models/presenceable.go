package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Presenceable 
type Presenceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetAvailability()(*string)
    GetOutOfOfficeSettings()(OutOfOfficeSettingsable)
    GetStatusMessage()(PresenceStatusMessageable)
    SetActivity(value *string)()
    SetAvailability(value *string)()
    SetOutOfOfficeSettings(value OutOfOfficeSettingsable)()
    SetStatusMessage(value PresenceStatusMessageable)()
}
