package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody 
type ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The userPrincipalName property
    userPrincipalName *string
}
// NewComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody instantiates a new ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody and sets the default values.
func NewComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody()(*ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) {
    m := &ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *ComanagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
