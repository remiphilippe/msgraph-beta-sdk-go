package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemPreviewPostRequestBody 
type ItemItemsItemPreviewPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allowEdit property
    allowEdit *bool
    // The chromeless property
    chromeless *bool
    // The page property
    page *string
    // The viewer property
    viewer *string
    // The zoom property
    zoom *float64
}
// NewItemItemsItemPreviewPostRequestBody instantiates a new ItemItemsItemPreviewPostRequestBody and sets the default values.
func NewItemItemsItemPreviewPostRequestBody()(*ItemItemsItemPreviewPostRequestBody) {
    m := &ItemItemsItemPreviewPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemPreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemPreviewPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowEdit gets the allowEdit property value. The allowEdit property
func (m *ItemItemsItemPreviewPostRequestBody) GetAllowEdit()(*bool) {
    return m.allowEdit
}
// GetChromeless gets the chromeless property value. The chromeless property
func (m *ItemItemsItemPreviewPostRequestBody) GetChromeless()(*bool) {
    return m.chromeless
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemPreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowEdit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEdit(val)
        }
        return nil
    }
    res["chromeless"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChromeless(val)
        }
        return nil
    }
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["viewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewer(val)
        }
        return nil
    }
    res["zoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoom(val)
        }
        return nil
    }
    return res
}
// GetPage gets the page property value. The page property
func (m *ItemItemsItemPreviewPostRequestBody) GetPage()(*string) {
    return m.page
}
// GetViewer gets the viewer property value. The viewer property
func (m *ItemItemsItemPreviewPostRequestBody) GetViewer()(*string) {
    return m.viewer
}
// GetZoom gets the zoom property value. The zoom property
func (m *ItemItemsItemPreviewPostRequestBody) GetZoom()(*float64) {
    return m.zoom
}
// Serialize serializes information the current object
func (m *ItemItemsItemPreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowEdit", m.GetAllowEdit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("chromeless", m.GetChromeless())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewer", m.GetViewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("zoom", m.GetZoom())
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
func (m *ItemItemsItemPreviewPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowEdit sets the allowEdit property value. The allowEdit property
func (m *ItemItemsItemPreviewPostRequestBody) SetAllowEdit(value *bool)() {
    m.allowEdit = value
}
// SetChromeless sets the chromeless property value. The chromeless property
func (m *ItemItemsItemPreviewPostRequestBody) SetChromeless(value *bool)() {
    m.chromeless = value
}
// SetPage sets the page property value. The page property
func (m *ItemItemsItemPreviewPostRequestBody) SetPage(value *string)() {
    m.page = value
}
// SetViewer sets the viewer property value. The viewer property
func (m *ItemItemsItemPreviewPostRequestBody) SetViewer(value *string)() {
    m.viewer = value
}
// SetZoom sets the zoom property value. The zoom property
func (m *ItemItemsItemPreviewPostRequestBody) SetZoom(value *float64)() {
    m.zoom = value
}
