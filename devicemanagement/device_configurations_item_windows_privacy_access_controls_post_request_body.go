package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody 
type DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The windowsPrivacyAccessControls property
    windowsPrivacyAccessControls []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable
}
// NewDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody instantiates a new DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody and sets the default values.
func NewDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody()(*DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) {
    m := &DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["windowsPrivacyAccessControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable)
            }
            m.SetWindowsPrivacyAccessControls(res)
        }
        return nil
    }
    return res
}
// GetWindowsPrivacyAccessControls gets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetWindowsPrivacyAccessControls()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable) {
    return m.windowsPrivacyAccessControls
}
// Serialize serializes information the current object
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWindowsPrivacyAccessControls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsPrivacyAccessControls()))
        for i, v := range m.GetWindowsPrivacyAccessControls() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("windowsPrivacyAccessControls", cast)
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
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWindowsPrivacyAccessControls sets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) SetWindowsPrivacyAccessControls(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsPrivacyDataAccessControlItemable)() {
    m.windowsPrivacyAccessControls = value
}
