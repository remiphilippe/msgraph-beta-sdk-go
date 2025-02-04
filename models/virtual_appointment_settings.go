package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualAppointmentSettings 
type VirtualAppointmentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the client can use the browser to join a virtual appointment. If set to false, the client can only use Microsoft Teams to join. Optional.
    allowClientToJoinUsingBrowser *bool
    // The OdataType property
    odataType *string
}
// NewVirtualAppointmentSettings instantiates a new virtualAppointmentSettings and sets the default values.
func NewVirtualAppointmentSettings()(*VirtualAppointmentSettings) {
    m := &VirtualAppointmentSettings{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateVirtualAppointmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualAppointmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualAppointmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualAppointmentSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowClientToJoinUsingBrowser gets the allowClientToJoinUsingBrowser property value. Indicates whether the client can use the browser to join a virtual appointment. If set to false, the client can only use Microsoft Teams to join. Optional.
func (m *VirtualAppointmentSettings) GetAllowClientToJoinUsingBrowser()(*bool) {
    return m.allowClientToJoinUsingBrowser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualAppointmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowClientToJoinUsingBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowClientToJoinUsingBrowser(val)
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
func (m *VirtualAppointmentSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *VirtualAppointmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowClientToJoinUsingBrowser", m.GetAllowClientToJoinUsingBrowser())
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
func (m *VirtualAppointmentSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowClientToJoinUsingBrowser sets the allowClientToJoinUsingBrowser property value. Indicates whether the client can use the browser to join a virtual appointment. If set to false, the client can only use Microsoft Teams to join. Optional.
func (m *VirtualAppointmentSettings) SetAllowClientToJoinUsingBrowser(value *bool)() {
    m.allowClientToJoinUsingBrowser = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualAppointmentSettings) SetOdataType(value *string)() {
    m.odataType = value
}
