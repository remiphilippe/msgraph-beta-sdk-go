package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Set 
type Set struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Children terms of set in term [store].
    children []Termable
    // Date and time of set creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description giving details on the term usage.
    description *string
    // Name of the set for each languageTag.
    localizedNames []LocalizedNameable
    // The parentGroup property
    parentGroup Groupable
    // Custom properties for the set.
    properties []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValueable
    // Indicates which terms have been pinned or reused directly under the set.
    relations []Relationable
    // All the terms under the set.
    terms []Termable
}
// NewSet instantiates a new set and sets the default values.
func NewSet()(*Set) {
    m := &Set{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSet(), nil
}
// GetChildren gets the children property value. Children terms of set in term [store].
func (m *Set) GetChildren()([]Termable) {
    return m.children
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of set creation. Read-only.
func (m *Set) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description giving details on the term usage.
func (m *Set) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Set) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTermFromDiscriminatorValue , m.SetChildren)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["localizedNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLocalizedNameFromDiscriminatorValue , m.SetLocalizedNames)
    res["parentGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupFromDiscriminatorValue , m.SetParentGroup)
    res["properties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateKeyValueFromDiscriminatorValue , m.SetProperties)
    res["relations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRelationFromDiscriminatorValue , m.SetRelations)
    res["terms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTermFromDiscriminatorValue , m.SetTerms)
    return res
}
// GetLocalizedNames gets the localizedNames property value. Name of the set for each languageTag.
func (m *Set) GetLocalizedNames()([]LocalizedNameable) {
    return m.localizedNames
}
// GetParentGroup gets the parentGroup property value. The parentGroup property
func (m *Set) GetParentGroup()(Groupable) {
    return m.parentGroup
}
// GetProperties gets the properties property value. Custom properties for the set.
func (m *Set) GetProperties()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValueable) {
    return m.properties
}
// GetRelations gets the relations property value. Indicates which terms have been pinned or reused directly under the set.
func (m *Set) GetRelations()([]Relationable) {
    return m.relations
}
// GetTerms gets the terms property value. All the terms under the set.
func (m *Set) GetTerms()([]Termable) {
    return m.terms
}
// Serialize serializes information the current object
func (m *Set) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildren())
        err = writer.WriteCollectionOfObjectValues("children", cast)
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
    if m.GetLocalizedNames() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLocalizedNames())
        err = writer.WriteCollectionOfObjectValues("localizedNames", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentGroup", m.GetParentGroup())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProperties())
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRelations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRelations())
        err = writer.WriteCollectionOfObjectValues("relations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTerms() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTerms())
        err = writer.WriteCollectionOfObjectValues("terms", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Children terms of set in term [store].
func (m *Set) SetChildren(value []Termable)() {
    m.children = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of set creation. Read-only.
func (m *Set) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description giving details on the term usage.
func (m *Set) SetDescription(value *string)() {
    m.description = value
}
// SetLocalizedNames sets the localizedNames property value. Name of the set for each languageTag.
func (m *Set) SetLocalizedNames(value []LocalizedNameable)() {
    m.localizedNames = value
}
// SetParentGroup sets the parentGroup property value. The parentGroup property
func (m *Set) SetParentGroup(value Groupable)() {
    m.parentGroup = value
}
// SetProperties sets the properties property value. Custom properties for the set.
func (m *Set) SetProperties(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValueable)() {
    m.properties = value
}
// SetRelations sets the relations property value. Indicates which terms have been pinned or reused directly under the set.
func (m *Set) SetRelations(value []Relationable)() {
    m.relations = value
}
// SetTerms sets the terms property value. All the terms under the set.
func (m *Set) SetTerms(value []Termable)() {
    m.terms = value
}
