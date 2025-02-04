package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody 
type ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceTag property
    deviceTag *string
}
// NewItemWipeManagedAppRegistrationByDeviceTagPostRequestBody instantiates a new ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody and sets the default values.
func NewItemWipeManagedAppRegistrationByDeviceTagPostRequestBody()(*ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) {
    m := &ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemWipeManagedAppRegistrationByDeviceTagPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemWipeManagedAppRegistrationByDeviceTagPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWipeManagedAppRegistrationByDeviceTagPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceTag gets the deviceTag property value. The deviceTag property
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) GetDeviceTag()(*string) {
    return m.deviceTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceTag", m.GetDeviceTag())
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
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceTag sets the deviceTag property value. The deviceTag property
func (m *ItemWipeManagedAppRegistrationByDeviceTagPostRequestBody) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
