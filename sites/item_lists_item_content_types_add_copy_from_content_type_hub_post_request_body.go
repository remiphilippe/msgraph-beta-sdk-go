package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody 
type ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contentTypeId property
    contentTypeId *string
}
// NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody instantiates a new ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody and sets the default values.
func NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody()(*ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) {
    m := &ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentTypeId gets the contentTypeId property value. The contentTypeId property
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetContentTypeId()(*string) {
    return m.contentTypeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentTypeId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentTypeId", m.GetContentTypeId())
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
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentTypeId sets the contentTypeId property value. The contentTypeId property
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) SetContentTypeId(value *string)() {
    m.contentTypeId = value
}
