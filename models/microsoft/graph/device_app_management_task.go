package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceAppManagementTask struct {
    Entity
    // The name or email of the admin this task is assigned to.
    assignedTo *string;
    // The category. Possible values are: unknown, advancedThreatProtection.
    category *DeviceAppManagementTaskCategory;
    // The created date.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The email address of the creator.
    creator *string;
    // Notes from the creator.
    creatorNotes *string;
    // The description.
    description *string;
    // The name.
    displayName *string;
    // The due date.
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The priority. Possible values are: none, high, low.
    priority *DeviceAppManagementTaskPriority;
    // The status. Possible values are: unknown, pending, active, completed, rejected.
    status *DeviceAppManagementTaskStatus;
}
// Instantiates a new deviceAppManagementTask and sets the default values.
func NewDeviceAppManagementTask()(*DeviceAppManagementTask) {
    m := &DeviceAppManagementTask{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignedTo property value. The name or email of the admin this task is assigned to.
func (m *DeviceAppManagementTask) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// Gets the category property value. The category. Possible values are: unknown, advancedThreatProtection.
func (m *DeviceAppManagementTask) GetCategory()(*DeviceAppManagementTaskCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the createdDateTime property value. The created date.
func (m *DeviceAppManagementTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the creator property value. The email address of the creator.
func (m *DeviceAppManagementTask) GetCreator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creator
    }
}
// Gets the creatorNotes property value. Notes from the creator.
func (m *DeviceAppManagementTask) GetCreatorNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorNotes
    }
}
// Gets the description property value. The description.
func (m *DeviceAppManagementTask) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name.
func (m *DeviceAppManagementTask) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the dueDateTime property value. The due date.
func (m *DeviceAppManagementTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// Gets the priority property value. The priority. Possible values are: none, high, low.
func (m *DeviceAppManagementTask) GetPriority()(*DeviceAppManagementTaskPriority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the status property value. The status. Possible values are: unknown, pending, active, completed, rejected.
func (m *DeviceAppManagementTask) GetStatus()(*DeviceAppManagementTaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *DeviceAppManagementTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceAppManagementTaskCategory)
            m.SetCategory(&cast)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val)
        }
        return nil
    }
    res["creatorNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatorNotes(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskPriority)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceAppManagementTaskPriority)
            m.SetPriority(&cast)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceAppManagementTaskStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *DeviceAppManagementTask) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceAppManagementTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := m.GetCategory().String()
        err = writer.WriteStringValue("category", &cast)
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
        err = writer.WriteStringValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorNotes", m.GetCreatorNotes())
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := m.GetPriority().String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedTo property value. The name or email of the admin this task is assigned to.
// Parameters:
//  - value : Value to set for the assignedTo property.
func (m *DeviceAppManagementTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// Sets the category property value. The category. Possible values are: unknown, advancedThreatProtection.
// Parameters:
//  - value : Value to set for the category property.
func (m *DeviceAppManagementTask) SetCategory(value *DeviceAppManagementTaskCategory)() {
    m.category = value
}
// Sets the createdDateTime property value. The created date.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceAppManagementTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the creator property value. The email address of the creator.
// Parameters:
//  - value : Value to set for the creator property.
func (m *DeviceAppManagementTask) SetCreator(value *string)() {
    m.creator = value
}
// Sets the creatorNotes property value. Notes from the creator.
// Parameters:
//  - value : Value to set for the creatorNotes property.
func (m *DeviceAppManagementTask) SetCreatorNotes(value *string)() {
    m.creatorNotes = value
}
// Sets the description property value. The description.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceAppManagementTask) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceAppManagementTask) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the dueDateTime property value. The due date.
// Parameters:
//  - value : Value to set for the dueDateTime property.
func (m *DeviceAppManagementTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
// Sets the priority property value. The priority. Possible values are: none, high, low.
// Parameters:
//  - value : Value to set for the priority property.
func (m *DeviceAppManagementTask) SetPriority(value *DeviceAppManagementTaskPriority)() {
    m.priority = value
}
// Sets the status property value. The status. Possible values are: unknown, pending, active, completed, rejected.
// Parameters:
//  - value : Value to set for the status property.
func (m *DeviceAppManagementTask) SetStatus(value *DeviceAppManagementTaskStatus)() {
    m.status = value
}
