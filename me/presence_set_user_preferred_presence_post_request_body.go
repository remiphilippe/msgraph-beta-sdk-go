package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PresenceSetUserPreferredPresencePostRequestBody 
type PresenceSetUserPreferredPresencePostRequestBody struct {
    // The activity property
    activity *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availability property
    availability *string
    // The expirationDuration property
    expirationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewPresenceSetUserPreferredPresencePostRequestBody instantiates a new PresenceSetUserPreferredPresencePostRequestBody and sets the default values.
func NewPresenceSetUserPreferredPresencePostRequestBody()(*PresenceSetUserPreferredPresencePostRequestBody) {
    m := &PresenceSetUserPreferredPresencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreatePresenceSetUserPreferredPresencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePresenceSetUserPreferredPresencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPresenceSetUserPreferredPresencePostRequestBody(), nil
}
// GetActivity gets the activity property value. The activity property
func (m *PresenceSetUserPreferredPresencePostRequestBody) GetActivity()(*string) {
    return m.activity
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PresenceSetUserPreferredPresencePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailability gets the availability property value. The availability property
func (m *PresenceSetUserPreferredPresencePostRequestBody) GetAvailability()(*string) {
    return m.availability
}
// GetExpirationDuration gets the expirationDuration property value. The expirationDuration property
func (m *PresenceSetUserPreferredPresencePostRequestBody) GetExpirationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.expirationDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PresenceSetUserPreferredPresencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["expirationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDuration(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PresenceSetUserPreferredPresencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("expirationDuration", m.GetExpirationDuration())
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
// SetActivity sets the activity property value. The activity property
func (m *PresenceSetUserPreferredPresencePostRequestBody) SetActivity(value *string)() {
    m.activity = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PresenceSetUserPreferredPresencePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailability sets the availability property value. The availability property
func (m *PresenceSetUserPreferredPresencePostRequestBody) SetAvailability(value *string)() {
    m.availability = value
}
// SetExpirationDuration sets the expirationDuration property value. The expirationDuration property
func (m *PresenceSetUserPreferredPresencePostRequestBody) SetExpirationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.expirationDuration = value
}
