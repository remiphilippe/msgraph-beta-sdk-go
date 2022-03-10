package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyAssignment provides operations to manage the deviceManagement singleton.
type DeviceManagementConfigurationPolicyAssignment struct {
    Entity
    // The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
    source *DeviceAndAppManagementAssignmentSource;
    // The identifier of the source of the assignment.
    sourceId *string;
    // The assignment target for the DeviceManagementConfigurationPolicy.
    target DeviceAndAppManagementAssignmentTargetable;
}
// NewDeviceManagementConfigurationPolicyAssignment instantiates a new deviceManagementConfigurationPolicyAssignment and sets the default values.
func NewDeviceManagementConfigurationPolicyAssignment()(*DeviceManagementConfigurationPolicyAssignment) {
    m := &DeviceManagementConfigurationPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*DeviceAndAppManagementAssignmentSource))
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
func (m *DeviceManagementConfigurationPolicyAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *DeviceManagementConfigurationPolicyAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSource sets the source property value. The assignment source for the device compliance policy, direct or parcel/policySet. Possible values are: direct, policySets.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
func (m *DeviceManagementConfigurationPolicyAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
