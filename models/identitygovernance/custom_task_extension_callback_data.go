package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CustomTaskExtensionCallbackData 
type CustomTaskExtensionCallbackData struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionData
    // The operationStatus property
    operationStatus *CustomTaskExtensionOperationStatus
}
// NewCustomTaskExtensionCallbackData instantiates a new CustomTaskExtensionCallbackData and sets the default values.
func NewCustomTaskExtensionCallbackData()(*CustomTaskExtensionCallbackData) {
    m := &CustomTaskExtensionCallbackData{
        CustomExtensionData: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewCustomExtensionData(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtensionCallbackData";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomTaskExtensionCallbackDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionCallbackDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtensionCallbackData(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtensionCallbackData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionData.GetFieldDeserializers()
    res["operationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomTaskExtensionOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationStatus(val.(*CustomTaskExtensionOperationStatus))
        }
        return nil
    }
    return res
}
// GetOperationStatus gets the operationStatus property value. The operationStatus property
func (m *CustomTaskExtensionCallbackData) GetOperationStatus()(*CustomTaskExtensionOperationStatus) {
    return m.operationStatus
}
// Serialize serializes information the current object
func (m *CustomTaskExtensionCallbackData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionData.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOperationStatus() != nil {
        cast := (*m.GetOperationStatus()).String()
        err = writer.WriteStringValue("operationStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOperationStatus sets the operationStatus property value. The operationStatus property
func (m *CustomTaskExtensionCallbackData) SetOperationStatus(value *CustomTaskExtensionOperationStatus)() {
    m.operationStatus = value
}
