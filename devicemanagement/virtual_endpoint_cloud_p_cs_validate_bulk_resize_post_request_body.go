package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// VirtualEndpointCloudPCsValidateBulkResizePostRequestBody 
type VirtualEndpointCloudPCsValidateBulkResizePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVirtualEndpointCloudPCsValidateBulkResizePostRequestBody instantiates a new VirtualEndpointCloudPCsValidateBulkResizePostRequestBody and sets the default values.
func NewVirtualEndpointCloudPCsValidateBulkResizePostRequestBody()(*VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) {
    m := &VirtualEndpointCloudPCsValidateBulkResizePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualEndpointCloudPCsValidateBulkResizePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointCloudPCsValidateBulkResizePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointCloudPCsValidateBulkResizePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCloudPcIds gets the cloudPcIds property value. The cloudPcIds property
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) GetCloudPcIds()([]string) {
    val, err := m.GetBackingStore().Get("cloudPcIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPcIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCloudPcIds(res)
        }
        return nil
    }
    res["targetServicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetTargetServicePlanId gets the targetServicePlanId property value. The targetServicePlanId property
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) GetTargetServicePlanId()(*string) {
    val, err := m.GetBackingStore().Get("targetServicePlanId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCloudPcIds() != nil {
        err := writer.WriteCollectionOfStringValues("cloudPcIds", m.GetCloudPcIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetServicePlanId", m.GetTargetServicePlanId())
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
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCloudPcIds sets the cloudPcIds property value. The cloudPcIds property
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) SetCloudPcIds(value []string)() {
    err := m.GetBackingStore().Set("cloudPcIds", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetServicePlanId sets the targetServicePlanId property value. The targetServicePlanId property
func (m *VirtualEndpointCloudPCsValidateBulkResizePostRequestBody) SetTargetServicePlanId(value *string)() {
    err := m.GetBackingStore().Set("targetServicePlanId", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEndpointCloudPCsValidateBulkResizePostRequestBodyable 
type VirtualEndpointCloudPCsValidateBulkResizePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCloudPcIds()([]string)
    GetTargetServicePlanId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCloudPcIds(value []string)()
    SetTargetServicePlanId(value *string)()
}
