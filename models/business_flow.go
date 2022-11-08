package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlow provides operations to manage the collection of approvalWorkflowProvider entities.
type BusinessFlow struct {
    Entity
    // The customData property
    customData *string
    // The deDuplicationId property
    deDuplicationId *string
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The policy property
    policy GovernancePolicyable
    // The policyTemplateId property
    policyTemplateId *string
    // The recordVersion property
    recordVersion *string
    // The schemaId property
    schemaId *string
    // The settings property
    settings BusinessFlowSettingsable
}
// NewBusinessFlow instantiates a new businessFlow and sets the default values.
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.businessFlow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBusinessFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessFlow(), nil
}
// GetCustomData gets the customData property value. The customData property
func (m *BusinessFlow) GetCustomData()(*string) {
    return m.customData
}
// GetDeDuplicationId gets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    return m.deDuplicationId
}
// GetDescription gets the description property value. The description property
func (m *BusinessFlow) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *BusinessFlow) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomData)
    res["deDuplicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeDuplicationId)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["policy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGovernancePolicyFromDiscriminatorValue , m.SetPolicy)
    res["policyTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPolicyTemplateId)
    res["recordVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecordVersion)
    res["schemaId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSchemaId)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBusinessFlowSettingsFromDiscriminatorValue , m.SetSettings)
    return res
}
// GetPolicy gets the policy property value. The policy property
func (m *BusinessFlow) GetPolicy()(GovernancePolicyable) {
    return m.policy
}
// GetPolicyTemplateId gets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    return m.policyTemplateId
}
// GetRecordVersion gets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) GetRecordVersion()(*string) {
    return m.recordVersion
}
// GetSchemaId gets the schemaId property value. The schemaId property
func (m *BusinessFlow) GetSchemaId()(*string) {
    return m.schemaId
}
// GetSettings gets the settings property value. The settings property
func (m *BusinessFlow) GetSettings()(BusinessFlowSettingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *BusinessFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customData", m.GetCustomData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deDuplicationId", m.GetDeDuplicationId())
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
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyTemplateId", m.GetPolicyTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordVersion", m.GetRecordVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schemaId", m.GetSchemaId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomData sets the customData property value. The customData property
func (m *BusinessFlow) SetCustomData(value *string)() {
    m.customData = value
}
// SetDeDuplicationId sets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    m.deDuplicationId = value
}
// SetDescription sets the description property value. The description property
func (m *BusinessFlow) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *BusinessFlow) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPolicy sets the policy property value. The policy property
func (m *BusinessFlow) SetPolicy(value GovernancePolicyable)() {
    m.policy = value
}
// SetPolicyTemplateId sets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    m.policyTemplateId = value
}
// SetRecordVersion sets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    m.recordVersion = value
}
// SetSchemaId sets the schemaId property value. The schemaId property
func (m *BusinessFlow) SetSchemaId(value *string)() {
    m.schemaId = value
}
// SetSettings sets the settings property value. The settings property
func (m *BusinessFlow) SetSettings(value BusinessFlowSettingsable)() {
    m.settings = value
}
