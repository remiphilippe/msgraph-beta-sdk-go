package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementIntentUserStateSummary struct {
    Entity
    // Number of users in conflict
    conflictCount *int32;
    // Number of error users
    errorCount *int32;
    // Number of failed users
    failedCount *int32;
    // Number of not applicable users
    notApplicableCount *int32;
    // Number of succeeded users
    successCount *int32;
}
// Instantiates a new deviceManagementIntentUserStateSummary and sets the default values.
func NewDeviceManagementIntentUserStateSummary()(*DeviceManagementIntentUserStateSummary) {
    m := &DeviceManagementIntentUserStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the conflictCount property value. Number of users in conflict
func (m *DeviceManagementIntentUserStateSummary) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
// Gets the errorCount property value. Number of error users
func (m *DeviceManagementIntentUserStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// Gets the failedCount property value. Number of failed users
func (m *DeviceManagementIntentUserStateSummary) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
// Gets the notApplicableCount property value. Number of not applicable users
func (m *DeviceManagementIntentUserStateSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// Gets the successCount property value. Number of succeeded users
func (m *DeviceManagementIntentUserStateSummary) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
// The deserialization information for the current model
func (m *DeviceManagementIntentUserStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["successCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementIntentUserStateSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementIntentUserStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("conflictCount", m.GetConflictCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the conflictCount property value. Number of users in conflict
// Parameters:
//  - value : Value to set for the conflictCount property.
func (m *DeviceManagementIntentUserStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
// Sets the errorCount property value. Number of error users
// Parameters:
//  - value : Value to set for the errorCount property.
func (m *DeviceManagementIntentUserStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// Sets the failedCount property value. Number of failed users
// Parameters:
//  - value : Value to set for the failedCount property.
func (m *DeviceManagementIntentUserStateSummary) SetFailedCount(value *int32)() {
    m.failedCount = value
}
// Sets the notApplicableCount property value. Number of not applicable users
// Parameters:
//  - value : Value to set for the notApplicableCount property.
func (m *DeviceManagementIntentUserStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// Sets the successCount property value. Number of succeeded users
// Parameters:
//  - value : Value to set for the successCount property.
func (m *DeviceManagementIntentUserStateSummary) SetSuccessCount(value *int32)() {
    m.successCount = value
}
