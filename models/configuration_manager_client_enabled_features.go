package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerClientEnabledFeatures configuration Manager client enabled features
type ConfigurationManagerClientEnabledFeatures struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether compliance policy is managed by Intune
    compliancePolicy *bool
    // Whether device configuration is managed by Intune
    deviceConfiguration *bool
    // Whether Endpoint Protection is managed by Intune
    endpointProtection *bool
    // Whether inventory is managed by Intune
    inventory *bool
    // Whether modern application is managed by Intune
    modernApps *bool
    // The OdataType property
    odataType *string
    // Whether Office application is managed by Intune
    officeApps *bool
    // Whether resource access is managed by Intune
    resourceAccess *bool
    // Whether Windows Update for Business is managed by Intune
    windowsUpdateForBusiness *bool
}
// NewConfigurationManagerClientEnabledFeatures instantiates a new configurationManagerClientEnabledFeatures and sets the default values.
func NewConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    m := &ConfigurationManagerClientEnabledFeatures{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerClientEnabledFeatures(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientEnabledFeatures) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompliancePolicy gets the compliancePolicy property value. Whether compliance policy is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetCompliancePolicy()(*bool) {
    return m.compliancePolicy
}
// GetDeviceConfiguration gets the deviceConfiguration property value. Whether device configuration is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetDeviceConfiguration()(*bool) {
    return m.deviceConfiguration
}
// GetEndpointProtection gets the endpointProtection property value. Whether Endpoint Protection is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetEndpointProtection()(*bool) {
    return m.endpointProtection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientEnabledFeatures) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compliancePolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCompliancePolicy)
    res["deviceConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceConfiguration)
    res["endpointProtection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEndpointProtection)
    res["inventory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetInventory)
    res["modernApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetModernApps)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["officeApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOfficeApps)
    res["resourceAccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetResourceAccess)
    res["windowsUpdateForBusiness"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWindowsUpdateForBusiness)
    return res
}
// GetInventory gets the inventory property value. Whether inventory is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetInventory()(*bool) {
    return m.inventory
}
// GetModernApps gets the modernApps property value. Whether modern application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetModernApps()(*bool) {
    return m.modernApps
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConfigurationManagerClientEnabledFeatures) GetOdataType()(*string) {
    return m.odataType
}
// GetOfficeApps gets the officeApps property value. Whether Office application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetOfficeApps()(*bool) {
    return m.officeApps
}
// GetResourceAccess gets the resourceAccess property value. Whether resource access is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetResourceAccess()(*bool) {
    return m.resourceAccess
}
// GetWindowsUpdateForBusiness gets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusiness()(*bool) {
    return m.windowsUpdateForBusiness
}
// Serialize serializes information the current object
func (m *ConfigurationManagerClientEnabledFeatures) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("compliancePolicy", m.GetCompliancePolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("endpointProtection", m.GetEndpointProtection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inventory", m.GetInventory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("modernApps", m.GetModernApps())
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
        err := writer.WriteBoolValue("officeApps", m.GetOfficeApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("resourceAccess", m.GetResourceAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("windowsUpdateForBusiness", m.GetWindowsUpdateForBusiness())
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
func (m *ConfigurationManagerClientEnabledFeatures) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompliancePolicy sets the compliancePolicy property value. Whether compliance policy is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetCompliancePolicy(value *bool)() {
    m.compliancePolicy = value
}
// SetDeviceConfiguration sets the deviceConfiguration property value. Whether device configuration is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetDeviceConfiguration(value *bool)() {
    m.deviceConfiguration = value
}
// SetEndpointProtection sets the endpointProtection property value. Whether Endpoint Protection is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetEndpointProtection(value *bool)() {
    m.endpointProtection = value
}
// SetInventory sets the inventory property value. Whether inventory is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetInventory(value *bool)() {
    m.inventory = value
}
// SetModernApps sets the modernApps property value. Whether modern application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetModernApps(value *bool)() {
    m.modernApps = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConfigurationManagerClientEnabledFeatures) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfficeApps sets the officeApps property value. Whether Office application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetOfficeApps(value *bool)() {
    m.officeApps = value
}
// SetResourceAccess sets the resourceAccess property value. Whether resource access is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetResourceAccess(value *bool)() {
    m.resourceAccess = value
}
// SetWindowsUpdateForBusiness sets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetWindowsUpdateForBusiness(value *bool)() {
    m.windowsUpdateForBusiness = value
}
