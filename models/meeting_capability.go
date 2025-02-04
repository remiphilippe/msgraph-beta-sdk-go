package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingCapability 
type MeetingCapability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether anonymous users dialout is allowed in a meeting.
    allowAnonymousUsersToDialOut *bool
    // Indicates whether anonymous users are allowed to start a meeting.
    allowAnonymousUsersToStartMeeting *bool
    // The autoAdmittedUsers property
    autoAdmittedUsers *AutoAdmittedUsersType
    // The OdataType property
    odataType *string
}
// NewMeetingCapability instantiates a new meetingCapability and sets the default values.
func NewMeetingCapability()(*MeetingCapability) {
    m := &MeetingCapability{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateMeetingCapabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingCapabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingCapability(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAnonymousUsersToDialOut gets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToDialOut()(*bool) {
    return m.allowAnonymousUsersToDialOut
}
// GetAllowAnonymousUsersToStartMeeting gets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) GetAllowAnonymousUsersToStartMeeting()(*bool) {
    return m.allowAnonymousUsersToStartMeeting
}
// GetAutoAdmittedUsers gets the autoAdmittedUsers property value. The autoAdmittedUsers property
func (m *MeetingCapability) GetAutoAdmittedUsers()(*AutoAdmittedUsersType) {
    return m.autoAdmittedUsers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingCapability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowAnonymousUsersToDialOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToDialOut(val)
        }
        return nil
    }
    res["allowAnonymousUsersToStartMeeting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToStartMeeting(val)
        }
        return nil
    }
    res["autoAdmittedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoAdmittedUsersType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAdmittedUsers(val.(*AutoAdmittedUsersType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingCapability) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MeetingCapability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToDialOut", m.GetAllowAnonymousUsersToDialOut())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToStartMeeting", m.GetAllowAnonymousUsersToStartMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetAutoAdmittedUsers() != nil {
        cast := (*m.GetAutoAdmittedUsers()).String()
        err := writer.WriteStringValue("autoAdmittedUsers", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAnonymousUsersToDialOut sets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToDialOut(value *bool)() {
    m.allowAnonymousUsersToDialOut = value
}
// SetAllowAnonymousUsersToStartMeeting sets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToStartMeeting(value *bool)() {
    m.allowAnonymousUsersToStartMeeting = value
}
// SetAutoAdmittedUsers sets the autoAdmittedUsers property value. The autoAdmittedUsers property
func (m *MeetingCapability) SetAutoAdmittedUsers(value *AutoAdmittedUsersType)() {
    m.autoAdmittedUsers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingCapability) SetOdataType(value *string)() {
    m.odataType = value
}
