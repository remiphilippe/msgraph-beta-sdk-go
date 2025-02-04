package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExchangeConnectorsItemSyncPostRequestBody 
type ExchangeConnectorsItemSyncPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of Exchange Connector sync requested.
    syncType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeConnectorSyncType
}
// NewExchangeConnectorsItemSyncPostRequestBody instantiates a new ExchangeConnectorsItemSyncPostRequestBody and sets the default values.
func NewExchangeConnectorsItemSyncPostRequestBody()(*ExchangeConnectorsItemSyncPostRequestBody) {
    m := &ExchangeConnectorsItemSyncPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateExchangeConnectorsItemSyncPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExchangeConnectorsItemSyncPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExchangeConnectorsItemSyncPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExchangeConnectorsItemSyncPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExchangeConnectorsItemSyncPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["syncType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseDeviceManagementExchangeConnectorSyncType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeConnectorSyncType))
        }
        return nil
    }
    return res
}
// GetSyncType gets the syncType property value. The type of Exchange Connector sync requested.
func (m *ExchangeConnectorsItemSyncPostRequestBody) GetSyncType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeConnectorSyncType) {
    return m.syncType
}
// Serialize serializes information the current object
func (m *ExchangeConnectorsItemSyncPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSyncType() != nil {
        cast := (*m.GetSyncType()).String()
        err := writer.WriteStringValue("syncType", &cast)
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
func (m *ExchangeConnectorsItemSyncPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSyncType sets the syncType property value. The type of Exchange Connector sync requested.
func (m *ExchangeConnectorsItemSyncPostRequestBody) SetSyncType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementExchangeConnectorSyncType)() {
    m.syncType = value
}
