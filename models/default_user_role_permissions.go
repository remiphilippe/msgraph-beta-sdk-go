package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRolePermissions 
type DefaultUserRolePermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the default user role can create applications.
    allowedToCreateApps *bool
    // Indicates whether the default user role can create security groups.
    allowedToCreateSecurityGroups *bool
    // Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
    allowedToReadBitlockerKeysForOwnedDevice *bool
    // Indicates whether the default user role can read other users.
    allowedToReadOtherUsers *bool
    // The OdataType property
    odataType *string
}
// NewDefaultUserRolePermissions instantiates a new defaultUserRolePermissions and sets the default values.
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateDefaultUserRolePermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRolePermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRolePermissions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedToCreateApps gets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    return m.allowedToCreateApps
}
// GetAllowedToCreateSecurityGroups gets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    return m.allowedToCreateSecurityGroups
}
// GetAllowedToReadBitlockerKeysForOwnedDevice gets the allowedToReadBitlockerKeysForOwnedDevice property value. Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
func (m *DefaultUserRolePermissions) GetAllowedToReadBitlockerKeysForOwnedDevice()(*bool) {
    return m.allowedToReadBitlockerKeysForOwnedDevice
}
// GetAllowedToReadOtherUsers gets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    return m.allowedToReadOtherUsers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRolePermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedToCreateApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateApps(val)
        }
        return nil
    }
    res["allowedToCreateSecurityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateSecurityGroups(val)
        }
        return nil
    }
    res["allowedToReadBitlockerKeysForOwnedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadBitlockerKeysForOwnedDevice(val)
        }
        return nil
    }
    res["allowedToReadOtherUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadOtherUsers(val)
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
func (m *DefaultUserRolePermissions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DefaultUserRolePermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowedToCreateApps", m.GetAllowedToCreateApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToCreateSecurityGroups", m.GetAllowedToCreateSecurityGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadBitlockerKeysForOwnedDevice", m.GetAllowedToReadBitlockerKeysForOwnedDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadOtherUsers", m.GetAllowedToReadOtherUsers())
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
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedToCreateApps sets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    m.allowedToCreateApps = value
}
// SetAllowedToCreateSecurityGroups sets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    m.allowedToCreateSecurityGroups = value
}
// SetAllowedToReadBitlockerKeysForOwnedDevice sets the allowedToReadBitlockerKeysForOwnedDevice property value. Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user role.
func (m *DefaultUserRolePermissions) SetAllowedToReadBitlockerKeysForOwnedDevice(value *bool)() {
    m.allowedToReadBitlockerKeysForOwnedDevice = value
}
// SetAllowedToReadOtherUsers sets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    m.allowedToReadOtherUsers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DefaultUserRolePermissions) SetOdataType(value *string)() {
    m.odataType = value
}
