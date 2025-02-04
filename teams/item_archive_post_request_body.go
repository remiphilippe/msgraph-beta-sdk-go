package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemArchivePostRequestBody 
type ItemArchivePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The shouldSetSpoSiteReadOnlyForMembers property
    shouldSetSpoSiteReadOnlyForMembers *bool
}
// NewItemArchivePostRequestBody instantiates a new ItemArchivePostRequestBody and sets the default values.
func NewItemArchivePostRequestBody()(*ItemArchivePostRequestBody) {
    m := &ItemArchivePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemArchivePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemArchivePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemArchivePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemArchivePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemArchivePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["shouldSetSpoSiteReadOnlyForMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldSetSpoSiteReadOnlyForMembers(val)
        }
        return nil
    }
    return res
}
// GetShouldSetSpoSiteReadOnlyForMembers gets the shouldSetSpoSiteReadOnlyForMembers property value. The shouldSetSpoSiteReadOnlyForMembers property
func (m *ItemArchivePostRequestBody) GetShouldSetSpoSiteReadOnlyForMembers()(*bool) {
    return m.shouldSetSpoSiteReadOnlyForMembers
}
// Serialize serializes information the current object
func (m *ItemArchivePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("shouldSetSpoSiteReadOnlyForMembers", m.GetShouldSetSpoSiteReadOnlyForMembers())
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
func (m *ItemArchivePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetShouldSetSpoSiteReadOnlyForMembers sets the shouldSetSpoSiteReadOnlyForMembers property value. The shouldSetSpoSiteReadOnlyForMembers property
func (m *ItemArchivePostRequestBody) SetShouldSetSpoSiteReadOnlyForMembers(value *bool)() {
    m.shouldSetSpoSiteReadOnlyForMembers = value
}
