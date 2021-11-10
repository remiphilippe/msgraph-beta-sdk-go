package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsDefenderApplicationControlSupplementalPolicyAssignment struct {
    Entity
    // The target group assignment defined by the admin.
    target *DeviceAndAppManagementAssignmentTarget;
}
// Instantiates a new windowsDefenderApplicationControlSupplementalPolicyAssignment and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyAssignment()(*WindowsDefenderApplicationControlSupplementalPolicyAssignment) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the target property value. The target group assignment defined by the admin.
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the target property value. The target group assignment defined by the admin.
// Parameters:
//  - value : Value to set for the target property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
