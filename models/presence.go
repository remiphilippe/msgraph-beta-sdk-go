package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Presence 
type Presence struct {
    Entity
    // The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive,InAMeeting, Offline, OffWork,OutOfOffice, PresenceUnknown,Presenting, UrgentInterruptionsOnly.
    activity *string
    // The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
    availability *string
    // The out of office settings for a user.
    outOfOfficeSettings OutOfOfficeSettingsable
    // The presence status message of a user.
    statusMessage PresenceStatusMessageable
}
// NewPresence instantiates a new presence and sets the default values.
func NewPresence()(*Presence) {
    m := &Presence{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePresenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePresenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPresence(), nil
}
// GetActivity gets the activity property value. The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive,InAMeeting, Offline, OffWork,OutOfOffice, PresenceUnknown,Presenting, UrgentInterruptionsOnly.
func (m *Presence) GetActivity()(*string) {
    return m.activity
}
// GetAvailability gets the availability property value. The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
func (m *Presence) GetAvailability()(*string) {
    return m.availability
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Presence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["availability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailability(val)
        }
        return nil
    }
    res["outOfOfficeSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutOfOfficeSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutOfOfficeSettings(val.(OutOfOfficeSettingsable))
        }
        return nil
    }
    res["statusMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePresenceStatusMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusMessage(val.(PresenceStatusMessageable))
        }
        return nil
    }
    return res
}
// GetOutOfOfficeSettings gets the outOfOfficeSettings property value. The out of office settings for a user.
func (m *Presence) GetOutOfOfficeSettings()(OutOfOfficeSettingsable) {
    return m.outOfOfficeSettings
}
// GetStatusMessage gets the statusMessage property value. The presence status message of a user.
func (m *Presence) GetStatusMessage()(PresenceStatusMessageable) {
    return m.statusMessage
}
// Serialize serializes information the current object
func (m *Presence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outOfOfficeSettings", m.GetOutOfOfficeSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("statusMessage", m.GetStatusMessage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive,InAMeeting, Offline, OffWork,OutOfOffice, PresenceUnknown,Presenting, UrgentInterruptionsOnly.
func (m *Presence) SetActivity(value *string)() {
    m.activity = value
}
// SetAvailability sets the availability property value. The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
func (m *Presence) SetAvailability(value *string)() {
    m.availability = value
}
// SetOutOfOfficeSettings sets the outOfOfficeSettings property value. The out of office settings for a user.
func (m *Presence) SetOutOfOfficeSettings(value OutOfOfficeSettingsable)() {
    m.outOfOfficeSettings = value
}
// SetStatusMessage sets the statusMessage property value. The presence status message of a user.
func (m *Presence) SetStatusMessage(value PresenceStatusMessageable)() {
    m.statusMessage = value
}
