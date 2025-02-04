package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *bool
}
// NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
