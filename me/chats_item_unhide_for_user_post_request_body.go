package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ChatsItemUnhideForUserPostRequestBody 
type ChatsItemUnhideForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tenantId property
    tenantId *string
    // The user property
    user ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkUserIdentityable
}
// NewChatsItemUnhideForUserPostRequestBody instantiates a new ChatsItemUnhideForUserPostRequestBody and sets the default values.
func NewChatsItemUnhideForUserPostRequestBody()(*ChatsItemUnhideForUserPostRequestBody) {
    m := &ChatsItemUnhideForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateChatsItemUnhideForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatsItemUnhideForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatsItemUnhideForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatsItemUnhideForUserPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatsItemUnhideForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ChatsItemUnhideForUserPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// GetUser gets the user property value. The user property
func (m *ChatsItemUnhideForUserPostRequestBody) GetUser()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ChatsItemUnhideForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ChatsItemUnhideForUserPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ChatsItemUnhideForUserPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUser sets the user property value. The user property
func (m *ChatsItemUnhideForUserPostRequestBody) SetUser(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkUserIdentityable)() {
    m.user = value
}
