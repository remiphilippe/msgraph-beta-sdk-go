package setscheduledretirestate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SetScheduledRetireStatePostRequestBody provides operations to call the setScheduledRetireState method.
type SetScheduledRetireStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceIds property
    managedDeviceIds []string
    // The scopedToAllDevices property
    scopedToAllDevices *bool
    // Cancel or confirm scheduled retire 
    state *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState
}
// NewSetScheduledRetireStatePostRequestBody instantiates a new setScheduledRetireStatePostRequestBody and sets the default values.
func NewSetScheduledRetireStatePostRequestBody()(*SetScheduledRetireStatePostRequestBody) {
    m := &SetScheduledRetireStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetScheduledRetireStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetScheduledRetireStatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetScheduledRetireStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetManagedDeviceIds)
    res["scopedToAllDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScopedToAllDevices)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseScheduledRetireState , m.SetState)
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *SetScheduledRetireStatePostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetScopedToAllDevices gets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *SetScheduledRetireStatePostRequestBody) GetScopedToAllDevices()(*bool) {
    return m.scopedToAllDevices
}
// GetState gets the state property value. Cancel or confirm scheduled retire 
func (m *SetScheduledRetireStatePostRequestBody) GetState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState) {
    return m.state
}
// Serialize serializes information the current object
func (m *SetScheduledRetireStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("scopedToAllDevices", m.GetScopedToAllDevices())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *SetScheduledRetireStatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *SetScheduledRetireStatePostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetScopedToAllDevices sets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *SetScheduledRetireStatePostRequestBody) SetScopedToAllDevices(value *bool)() {
    m.scopedToAllDevices = value
}
// SetState sets the state property value. Cancel or confirm scheduled retire 
func (m *SetScheduledRetireStatePostRequestBody) SetState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)() {
    m.state = value
}
