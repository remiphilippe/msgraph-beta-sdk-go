package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DowngradeJustification 
type DowngradeJustification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the downgrade is or is not justified.
    isDowngradeJustified *bool
    // Message that indicates why a downgrade is justified. The message will appear in administrative logs.
    justificationMessage *string
    // The OdataType property
    odataType *string
}
// NewDowngradeJustification instantiates a new downgradeJustification and sets the default values.
func NewDowngradeJustification()(*DowngradeJustification) {
    m := &DowngradeJustification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDowngradeJustificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDowngradeJustificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDowngradeJustification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DowngradeJustification) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DowngradeJustification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDowngradeJustified"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDowngradeJustified)
    res["justificationMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJustificationMessage)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsDowngradeJustified gets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
func (m *DowngradeJustification) GetIsDowngradeJustified()(*bool) {
    return m.isDowngradeJustified
}
// GetJustificationMessage gets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
func (m *DowngradeJustification) GetJustificationMessage()(*string) {
    return m.justificationMessage
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DowngradeJustification) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DowngradeJustification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDowngradeJustified", m.GetIsDowngradeJustified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationMessage", m.GetJustificationMessage())
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
func (m *DowngradeJustification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsDowngradeJustified sets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
func (m *DowngradeJustification) SetIsDowngradeJustified(value *bool)() {
    m.isDowngradeJustified = value
}
// SetJustificationMessage sets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
func (m *DowngradeJustification) SetJustificationMessage(value *string)() {
    m.justificationMessage = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DowngradeJustification) SetOdataType(value *string)() {
    m.odataType = value
}
