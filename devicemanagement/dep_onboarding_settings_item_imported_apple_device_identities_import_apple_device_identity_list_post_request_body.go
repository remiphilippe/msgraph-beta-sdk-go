package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody 
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The importedAppleDeviceIdentities property
    importedAppleDeviceIdentities []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable
    // The overwriteImportedDeviceIdentities property
    overwriteImportedDeviceIdentities *bool
}
// NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody instantiates a new DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody and sets the default values.
func NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody()(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) {
    m := &DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedAppleDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["overwriteImportedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteImportedDeviceIdentities(val)
        }
        return nil
    }
    return res
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetImportedAppleDeviceIdentities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable) {
    return m.importedAppleDeviceIdentities
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    return m.overwriteImportedDeviceIdentities
}
// Serialize serializes information the current object
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteImportedDeviceIdentities", m.GetOverwriteImportedDeviceIdentities())
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
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The importedAppleDeviceIdentities property
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetImportedAppleDeviceIdentities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable)() {
    m.importedAppleDeviceIdentities = value
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. The overwriteImportedDeviceIdentities property
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}
