package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentConfiguration the Base Class of Device Enrollment Configuration
type DeviceEnrollmentConfiguration struct {
    Entity
    // The list of group assignments for the device configuration profile
    assignments []EnrollmentConfigurationAssignmentable
    // Created date time in UTC of the device enrollment configuration
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the device enrollment configuration
    description *string
    // Describes the TemplateFamily for the Template entity
    deviceEnrollmentConfigurationType *DeviceEnrollmentConfigurationType
    // The display name of the device enrollment configuration
    displayName *string
    // Last modified date time in UTC of the device enrollment configuration
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value.
    priority *int32
    // Optional role scope tags for the enrollment restrictions.
    roleScopeTagIds []string
    // The version of the device enrollment configuration
    version *int32
}
// NewDeviceEnrollmentConfiguration instantiates a new deviceEnrollmentConfiguration and sets the default values.
func NewDeviceEnrollmentConfiguration()(*DeviceEnrollmentConfiguration) {
    m := &DeviceEnrollmentConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceEnrollmentConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceComanagementAuthorityConfiguration":
                        return NewDeviceComanagementAuthorityConfiguration(), nil
                    case "#microsoft.graph.deviceEnrollmentLimitConfiguration":
                        return NewDeviceEnrollmentLimitConfiguration(), nil
                    case "#microsoft.graph.deviceEnrollmentNotificationConfiguration":
                        return NewDeviceEnrollmentNotificationConfiguration(), nil
                    case "#microsoft.graph.deviceEnrollmentPlatformRestrictionConfiguration":
                        return NewDeviceEnrollmentPlatformRestrictionConfiguration(), nil
                    case "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfiguration":
                        return NewDeviceEnrollmentPlatformRestrictionsConfiguration(), nil
                    case "#microsoft.graph.deviceEnrollmentWindowsHelloForBusinessConfiguration":
                        return NewDeviceEnrollmentWindowsHelloForBusinessConfiguration(), nil
                    case "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration":
                        return NewWindows10EnrollmentCompletionPageConfiguration(), nil
                }
            }
        }
    }
    return NewDeviceEnrollmentConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the device configuration profile
func (m *DeviceEnrollmentConfiguration) GetAssignments()([]EnrollmentConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. Created date time in UTC of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDeviceEnrollmentConfigurationType gets the deviceEnrollmentConfigurationType property value. Describes the TemplateFamily for the Template entity
func (m *DeviceEnrollmentConfiguration) GetDeviceEnrollmentConfigurationType()(*DeviceEnrollmentConfigurationType) {
    return m.deviceEnrollmentConfigurationType
}
// GetDisplayName gets the displayName property value. The display name of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceEnrollmentConfigurationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceEnrollmentConfigurationType , m.SetDeviceEnrollmentConfigurationType)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date time in UTC of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPriority gets the priority property value. Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value.
func (m *DeviceEnrollmentConfiguration) GetPriority()(*int32) {
    return m.priority
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. Optional role scope tags for the enrollment restrictions.
func (m *DeviceEnrollmentConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetVersion gets the version property value. The version of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurationType() != nil {
        cast := (*m.GetDeviceEnrollmentConfigurationType()).String()
        err = writer.WriteStringValue("deviceEnrollmentConfigurationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the device configuration profile
func (m *DeviceEnrollmentConfiguration) SetAssignments(value []EnrollmentConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Created date time in UTC of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceEnrollmentConfigurationType sets the deviceEnrollmentConfigurationType property value. Describes the TemplateFamily for the Template entity
func (m *DeviceEnrollmentConfiguration) SetDeviceEnrollmentConfigurationType(value *DeviceEnrollmentConfigurationType)() {
    m.deviceEnrollmentConfigurationType = value
}
// SetDisplayName sets the displayName property value. The display name of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date time in UTC of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPriority sets the priority property value. Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value.
func (m *DeviceEnrollmentConfiguration) SetPriority(value *int32)() {
    m.priority = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. Optional role scope tags for the enrollment restrictions.
func (m *DeviceEnrollmentConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetVersion sets the version property value. The version of the device enrollment configuration
func (m *DeviceEnrollmentConfiguration) SetVersion(value *int32)() {
    m.version = value
}
