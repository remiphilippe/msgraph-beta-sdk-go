package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemChatsItemSendActivityNotificationPostRequestBody 
type ItemChatsItemSendActivityNotificationPostRequestBody struct {
    // The activityType property
    activityType *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chainId property
    chainId *int64
    // The previewText property
    previewText ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable
    // The recipient property
    recipient ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkNotificationRecipientable
    // The templateParameters property
    templateParameters []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable
    // The topic property
    topic ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkActivityTopicable
}
// NewItemChatsItemSendActivityNotificationPostRequestBody instantiates a new ItemChatsItemSendActivityNotificationPostRequestBody and sets the default values.
func NewItemChatsItemSendActivityNotificationPostRequestBody()(*ItemChatsItemSendActivityNotificationPostRequestBody) {
    m := &ItemChatsItemSendActivityNotificationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemChatsItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemSendActivityNotificationPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetActivityType()(*string) {
    return m.activityType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChainId gets the chainId property value. The chainId property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetChainId()(*int64) {
    return m.chainId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable))
        }
        return nil
    }
    res["recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkNotificationRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkNotificationRecipientable))
        }
        return nil
    }
    res["templateParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetPreviewText()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable) {
    return m.previewText
}
// GetRecipient gets the recipient property value. The recipient property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetRecipient()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkNotificationRecipientable) {
    return m.recipient
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetTemplateParameters()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable) {
    return m.templateParameters
}
// GetTopic gets the topic property value. The topic property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) GetTopic()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkActivityTopicable) {
    return m.topic
}
// Serialize serializes information the current object
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// SetActivityType sets the activityType property value. The activityType property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetPreviewText(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable)() {
    m.previewText = value
}
// SetRecipient sets the recipient property value. The recipient property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetRecipient(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkNotificationRecipientable)() {
    m.recipient = value
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetTemplateParameters(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)() {
    m.templateParameters = value
}
// SetTopic sets the topic property value. The topic property
func (m *ItemChatsItemSendActivityNotificationPostRequestBody) SetTopic(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkActivityTopicable)() {
    m.topic = value
}
