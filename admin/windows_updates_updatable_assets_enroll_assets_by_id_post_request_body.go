package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody 
type WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
    // The updateCategory property
    updateCategory *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory
}
// NewWindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody instantiates a new WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody()(*WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) {
    m := &WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory))
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory) {
    return m.updateCategory
}
// Serialize serializes information the current object
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err := writer.WriteStringValue("updateCategory", &cast)
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
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesUpdatableAssetsEnrollAssetsByIdPostRequestBody) SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)() {
    m.updateCategory = value
}
