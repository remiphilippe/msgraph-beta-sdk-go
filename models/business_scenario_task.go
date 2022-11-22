package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioTask 
type BusinessScenarioTask struct {
    PlannerTask
    // The businessScenarioProperties property
    businessScenarioProperties BusinessScenarioPropertiesable
    // The target property
    target BusinessScenarioTaskTargetBaseable
}
// NewBusinessScenarioTask instantiates a new BusinessScenarioTask and sets the default values.
func NewBusinessScenarioTask()(*BusinessScenarioTask) {
    m := &BusinessScenarioTask{
        PlannerTask: *NewPlannerTask(),
    }
    return m
}
// CreateBusinessScenarioTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioTask(), nil
}
// GetBusinessScenarioProperties gets the businessScenarioProperties property value. The businessScenarioProperties property
func (m *BusinessScenarioTask) GetBusinessScenarioProperties()(BusinessScenarioPropertiesable) {
    return m.businessScenarioProperties
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTask.GetFieldDeserializers()
    res["businessScenarioProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBusinessScenarioPropertiesFromDiscriminatorValue , m.SetBusinessScenarioProperties)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetTarget gets the target property value. The target property
func (m *BusinessScenarioTask) GetTarget()(BusinessScenarioTaskTargetBaseable) {
    return m.target
}
// Serialize serializes information the current object
func (m *BusinessScenarioTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerTask.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("businessScenarioProperties", m.GetBusinessScenarioProperties())
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
// SetBusinessScenarioProperties sets the businessScenarioProperties property value. The businessScenarioProperties property
func (m *BusinessScenarioTask) SetBusinessScenarioProperties(value BusinessScenarioPropertiesable)() {
    m.businessScenarioProperties = value
}
// SetTarget sets the target property value. The target property
func (m *BusinessScenarioTask) SetTarget(value BusinessScenarioTaskTargetBaseable)() {
    m.target = value
}
