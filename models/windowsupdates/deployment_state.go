package windowsupdates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentState 
type DeploymentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Specifies the reasons the deployment has its state value. Read-only.
    reasons []DeploymentStateReasonable
    // The requestedValue property
    requestedValue *RequestedDeploymentStateValue
    // The value property
    value *DeploymentStateValue
}
// NewDeploymentState instantiates a new deploymentState and sets the default values.
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeploymentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["reasons"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeploymentStateReasonFromDiscriminatorValue , m.SetReasons)
    res["requestedValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRequestedDeploymentStateValue , m.SetRequestedValue)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeploymentStateValue , m.SetValue)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeploymentState) GetOdataType()(*string) {
    return m.odataType
}
// GetReasons gets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) GetReasons()([]DeploymentStateReasonable) {
    return m.reasons
}
// GetRequestedValue gets the requestedValue property value. The requestedValue property
func (m *DeploymentState) GetRequestedValue()(*RequestedDeploymentStateValue) {
    return m.requestedValue
}
// GetValue gets the value property value. The value property
func (m *DeploymentState) GetValue()(*DeploymentStateValue) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeploymentState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetReasons() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReasons())
        err := writer.WriteCollectionOfObjectValues("reasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedValue() != nil {
        cast := (*m.GetRequestedValue()).String()
        err := writer.WriteStringValue("requestedValue", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *DeploymentState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReasons sets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) SetReasons(value []DeploymentStateReasonable)() {
    m.reasons = value
}
// SetRequestedValue sets the requestedValue property value. The requestedValue property
func (m *DeploymentState) SetRequestedValue(value *RequestedDeploymentStateValue)() {
    m.requestedValue = value
}
// SetValue sets the value property value. The value property
func (m *DeploymentState) SetValue(value *DeploymentStateValue)() {
    m.value = value
}
