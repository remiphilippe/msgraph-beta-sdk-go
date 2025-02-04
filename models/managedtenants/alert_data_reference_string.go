package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertDataReferenceString 
type AlertDataReferenceString struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The displayName property
    displayName *string
    // The OdataType property
    odataType *string
}
// NewAlertDataReferenceString instantiates a new alertDataReferenceString and sets the default values.
func NewAlertDataReferenceString()(*AlertDataReferenceString) {
    m := &AlertDataReferenceString{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAlertDataReferenceStringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertDataReferenceStringFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertDataReferenceString(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertDataReferenceString) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AlertDataReferenceString) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertDataReferenceString) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
func (m *AlertDataReferenceString) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AlertDataReferenceString) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *AlertDataReferenceString) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AlertDataReferenceString) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertDataReferenceString) SetOdataType(value *string)() {
    m.odataType = value
}
