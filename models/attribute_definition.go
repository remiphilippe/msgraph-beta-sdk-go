package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeDefinition 
type AttributeDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
    anchor *bool
    // The apiExpressions property
    apiExpressions []StringKeyStringValuePairable
    // true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
    caseExact *bool
    // The defaultValue property
    defaultValue *string
    // 'true' to allow null values for attributes.
    flowNullValues *bool
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntryable
    // true if an attribute can have multiple values. Default is false.
    multivalued *bool
    // The mutability property
    mutability *Mutability
    // Name of the attribute. Must be unique within the object definition. Not nullable.
    name *string
    // The OdataType property
    odataType *string
    // For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
    referencedObjects []ReferencedObjectable
    // true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
    required *bool
    // The type property
    type_escaped *AttributeType
}
// NewAttributeDefinition instantiates a new attributeDefinition and sets the default values.
func NewAttributeDefinition()(*AttributeDefinition) {
    m := &AttributeDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttributeDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeDefinition) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAnchor gets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) GetAnchor()(*bool) {
    return m.anchor
}
// GetApiExpressions gets the apiExpressions property value. The apiExpressions property
func (m *AttributeDefinition) GetApiExpressions()([]StringKeyStringValuePairable) {
    return m.apiExpressions
}
// GetCaseExact gets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) GetCaseExact()(*bool) {
    return m.caseExact
}
// GetDefaultValue gets the defaultValue property value. The defaultValue property
func (m *AttributeDefinition) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["anchor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAnchor)
    res["apiExpressions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateStringKeyStringValuePairFromDiscriminatorValue , m.SetApiExpressions)
    res["caseExact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCaseExact)
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultValue)
    res["flowNullValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFlowNullValues)
    res["metadata"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMetadataEntryFromDiscriminatorValue , m.SetMetadata)
    res["multivalued"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetMultivalued)
    res["mutability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMutability , m.SetMutability)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["referencedObjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateReferencedObjectFromDiscriminatorValue , m.SetReferencedObjects)
    res["required"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequired)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAttributeType , m.SetType)
    return res
}
// GetFlowNullValues gets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) GetFlowNullValues()(*bool) {
    return m.flowNullValues
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) GetMetadata()([]MetadataEntryable) {
    return m.metadata
}
// GetMultivalued gets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) GetMultivalued()(*bool) {
    return m.multivalued
}
// GetMutability gets the mutability property value. The mutability property
func (m *AttributeDefinition) GetMutability()(*Mutability) {
    return m.mutability
}
// GetName gets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttributeDefinition) GetOdataType()(*string) {
    return m.odataType
}
// GetReferencedObjects gets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) GetReferencedObjects()([]ReferencedObjectable) {
    return m.referencedObjects
}
// GetRequired gets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) GetRequired()(*bool) {
    return m.required
}
// GetType gets the type property value. The type property
func (m *AttributeDefinition) GetType()(*AttributeType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *AttributeDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("anchor", m.GetAnchor())
        if err != nil {
            return err
        }
    }
    if m.GetApiExpressions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApiExpressions())
        err := writer.WriteCollectionOfObjectValues("apiExpressions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("caseExact", m.GetCaseExact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("flowNullValues", m.GetFlowNullValues())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMetadata())
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("multivalued", m.GetMultivalued())
        if err != nil {
            return err
        }
    }
    if m.GetMutability() != nil {
        cast := (*m.GetMutability()).String()
        err := writer.WriteStringValue("mutability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetReferencedObjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReferencedObjects())
        err := writer.WriteCollectionOfObjectValues("referencedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAnchor sets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) SetAnchor(value *bool)() {
    m.anchor = value
}
// SetApiExpressions sets the apiExpressions property value. The apiExpressions property
func (m *AttributeDefinition) SetApiExpressions(value []StringKeyStringValuePairable)() {
    m.apiExpressions = value
}
// SetCaseExact sets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) SetCaseExact(value *bool)() {
    m.caseExact = value
}
// SetDefaultValue sets the defaultValue property value. The defaultValue property
func (m *AttributeDefinition) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetFlowNullValues sets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) SetFlowNullValues(value *bool)() {
    m.flowNullValues = value
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) SetMetadata(value []MetadataEntryable)() {
    m.metadata = value
}
// SetMultivalued sets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) SetMultivalued(value *bool)() {
    m.multivalued = value
}
// SetMutability sets the mutability property value. The mutability property
func (m *AttributeDefinition) SetMutability(value *Mutability)() {
    m.mutability = value
}
// SetName sets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttributeDefinition) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReferencedObjects sets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) SetReferencedObjects(value []ReferencedObjectable)() {
    m.referencedObjects = value
}
// SetRequired sets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) SetRequired(value *bool)() {
    m.required = value
}
// SetType sets the type property value. The type property
func (m *AttributeDefinition) SetType(value *AttributeType)() {
    m.type_escaped = value
}
