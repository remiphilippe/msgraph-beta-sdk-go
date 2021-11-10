package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceComplianceScript struct {
    Entity
    // The list of group assignments for the device compliance script
    assignments []DeviceHealthScriptAssignment;
    // The timestamp of when the device compliance script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the device compliance script
    description *string;
    // The entire content of the detection powershell script
    detectionScriptContent []byte;
    // List of run states for the device compliance script across all devices
    deviceRunStates []DeviceComplianceScriptDeviceState;
    // Name of the device compliance script
    displayName *string;
    // Indicate whether the script signature needs be checked
    enforceSignatureCheck *bool;
    // The timestamp of when the device compliance script was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the device compliance script publisher
    publisher *string;
    // List of Scope Tag IDs for the device compliance script
    roleScopeTagIds []string;
    // Indicate whether PowerShell script(s) should run as 32-bit
    runAs32Bit *bool;
    // Indicates the type of execution context. Possible values are: system, user.
    runAsAccount *RunAsAccountType;
    // High level run summary for device compliance script.
    runSummary *DeviceComplianceScriptRunSummary;
    // Version of the device compliance script
    version *string;
}
// Instantiates a new deviceComplianceScript and sets the default values.
func NewDeviceComplianceScript()(*DeviceComplianceScript) {
    m := &DeviceComplianceScript{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the device compliance script
func (m *DeviceComplianceScript) GetAssignments()([]DeviceHealthScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
func (m *DeviceComplianceScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description of the device compliance script
func (m *DeviceComplianceScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceComplianceScript) GetDetectionScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptContent
    }
}
// Gets the deviceRunStates property value. List of run states for the device compliance script across all devices
func (m *DeviceComplianceScript) GetDeviceRunStates()([]DeviceComplianceScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// Gets the displayName property value. Name of the device compliance script
func (m *DeviceComplianceScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceComplianceScript) GetEnforceSignatureCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceSignatureCheck
    }
}
// Gets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
func (m *DeviceComplianceScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the publisher property value. Name of the device compliance script publisher
func (m *DeviceComplianceScript) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
func (m *DeviceComplianceScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceComplianceScript) GetRunAs32Bit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runAs32Bit
    }
}
// Gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceComplianceScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// Gets the runSummary property value. High level run summary for device compliance script.
func (m *DeviceComplianceScript) GetRunSummary()(*DeviceComplianceScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// Gets the version property value. Version of the device compliance script
func (m *DeviceComplianceScript) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceComplianceScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptAssignment))
            }
            m.SetAssignments(res)
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
    res["detectionScriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptContent(val)
        }
        return nil
    }
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScriptDeviceState))
            }
            m.SetDeviceRunStates(res)
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
    res["enforceSignatureCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["runAs32Bit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RunAsAccountType)
            m.SetRunAsAccount(&cast)
        }
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRunSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(*DeviceComplianceScriptRunSummary))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *DeviceComplianceScript) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceComplianceScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        err = writer.WriteByteArrayValue("detectionScriptContent", m.GetDetectionScriptContent())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
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
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := m.GetRunAsAccount().String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for the device compliance script
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceComplianceScript) SetAssignments(value []DeviceHealthScriptAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceComplianceScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description of the device compliance script
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceComplianceScript) SetDescription(value *string)() {
    m.description = value
}
// Sets the detectionScriptContent property value. The entire content of the detection powershell script
// Parameters:
//  - value : Value to set for the detectionScriptContent property.
func (m *DeviceComplianceScript) SetDetectionScriptContent(value []byte)() {
    m.detectionScriptContent = value
}
// Sets the deviceRunStates property value. List of run states for the device compliance script across all devices
// Parameters:
//  - value : Value to set for the deviceRunStates property.
func (m *DeviceComplianceScript) SetDeviceRunStates(value []DeviceComplianceScriptDeviceState)() {
    m.deviceRunStates = value
}
// Sets the displayName property value. Name of the device compliance script
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceComplianceScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
// Parameters:
//  - value : Value to set for the enforceSignatureCheck property.
func (m *DeviceComplianceScript) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// Sets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceComplianceScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the publisher property value. Name of the device compliance script publisher
// Parameters:
//  - value : Value to set for the publisher property.
func (m *DeviceComplianceScript) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceComplianceScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
// Parameters:
//  - value : Value to set for the runAs32Bit property.
func (m *DeviceComplianceScript) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// Sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
// Parameters:
//  - value : Value to set for the runAsAccount property.
func (m *DeviceComplianceScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// Sets the runSummary property value. High level run summary for device compliance script.
// Parameters:
//  - value : Value to set for the runSummary property.
func (m *DeviceComplianceScript) SetRunSummary(value *DeviceComplianceScriptRunSummary)() {
    m.runSummary = value
}
// Sets the version property value. Version of the device compliance script
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceComplianceScript) SetVersion(value *string)() {
    m.version = value
}
