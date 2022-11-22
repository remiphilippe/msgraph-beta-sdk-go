package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalPlanSource 
type PlannerExternalPlanSource struct {
    PlannerPlanCreation
    // The contextScenarioId property
    contextScenarioId *string
    // The externalContextId property
    externalContextId *string
    // The externalObjectId property
    externalObjectId *string
}
// NewPlannerExternalPlanSource instantiates a new PlannerExternalPlanSource and sets the default values.
func NewPlannerExternalPlanSource()(*PlannerExternalPlanSource) {
    m := &PlannerExternalPlanSource{
        PlannerPlanCreation: *NewPlannerPlanCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerExternalPlanSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerExternalPlanSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerExternalPlanSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerExternalPlanSource(), nil
}
// GetContextScenarioId gets the contextScenarioId property value. The contextScenarioId property
func (m *PlannerExternalPlanSource) GetContextScenarioId()(*string) {
    return m.contextScenarioId
}
// GetExternalContextId gets the externalContextId property value. The externalContextId property
func (m *PlannerExternalPlanSource) GetExternalContextId()(*string) {
    return m.externalContextId
}
// GetExternalObjectId gets the externalObjectId property value. The externalObjectId property
func (m *PlannerExternalPlanSource) GetExternalObjectId()(*string) {
    return m.externalObjectId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerExternalPlanSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPlanCreation.GetFieldDeserializers()
    res["contextScenarioId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContextScenarioId)
    res["externalContextId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalContextId)
    res["externalObjectId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalObjectId)
    return res
}
// Serialize serializes information the current object
func (m *PlannerExternalPlanSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPlanCreation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contextScenarioId", m.GetContextScenarioId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalContextId", m.GetExternalContextId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalObjectId", m.GetExternalObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextScenarioId sets the contextScenarioId property value. The contextScenarioId property
func (m *PlannerExternalPlanSource) SetContextScenarioId(value *string)() {
    m.contextScenarioId = value
}
// SetExternalContextId sets the externalContextId property value. The externalContextId property
func (m *PlannerExternalPlanSource) SetExternalContextId(value *string)() {
    m.externalContextId = value
}
// SetExternalObjectId sets the externalObjectId property value. The externalObjectId property
func (m *PlannerExternalPlanSource) SetExternalObjectId(value *string)() {
    m.externalObjectId = value
}
