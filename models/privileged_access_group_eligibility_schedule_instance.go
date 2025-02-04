package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupEligibilityScheduleInstance 
type PrivilegedAccessGroupEligibilityScheduleInstance struct {
    PrivilegedAccessScheduleInstance
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
    // The eligibilityScheduleId property
    eligibilityScheduleId *string
    // The group property
    group Groupable
    // The groupId property
    groupId *string
    // The memberType property
    memberType *PrivilegedAccessGroupMemberType
    // The principal property
    principal DirectoryObjectable
    // The principalId property
    principalId *string
}
// NewPrivilegedAccessGroupEligibilityScheduleInstance instantiates a new privilegedAccessGroupEligibilityScheduleInstance and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleInstance()(*PrivilegedAccessGroupEligibilityScheduleInstance) {
    m := &PrivilegedAccessGroupEligibilityScheduleInstance{
        PrivilegedAccessScheduleInstance: *NewPrivilegedAccessScheduleInstance(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilityScheduleInstance(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetEligibilityScheduleId gets the eligibilityScheduleId property value. The eligibilityScheduleId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetEligibilityScheduleId()(*string) {
    return m.eligibilityScheduleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleInstance.GetFieldDeserializers()
    res["accessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupRelationships)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessId(val.(*PrivilegedAccessGroupRelationships))
        }
        return nil
    }
    res["eligibilityScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibilityScheduleId(val)
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val.(*PrivilegedAccessGroupMemberType))
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(DirectoryObjectable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetGroupId()(*string) {
    return m.groupId
}
// GetMemberType gets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetMemberType()(*PrivilegedAccessGroupMemberType) {
    return m.memberType
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetPrincipalId()(*string) {
    return m.principalId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessId() != nil {
        cast := (*m.GetAccessId()).String()
        err = writer.WriteStringValue("accessId", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eligibilityScheduleId", m.GetEligibilityScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberType() != nil {
        cast := (*m.GetMemberType()).String()
        err = writer.WriteStringValue("memberType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessId sets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetEligibilityScheduleId sets the eligibilityScheduleId property value. The eligibilityScheduleId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetEligibilityScheduleId(value *string)() {
    m.eligibilityScheduleId = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetGroupId(value *string)() {
    m.groupId = value
}
// SetMemberType sets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetMemberType(value *PrivilegedAccessGroupMemberType)() {
    m.memberType = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetPrincipalId(value *string)() {
    m.principalId = value
}
