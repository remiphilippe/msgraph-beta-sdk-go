package externalconnectors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlToItemResolverBase 
type UrlToItemResolverBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The priority which defines the sequence in which the urlToItemResolverBase instances are evaluated.
    priority *int32
}
// NewUrlToItemResolverBase instantiates a new urlToItemResolverBase and sets the default values.
func NewUrlToItemResolverBase()(*UrlToItemResolverBase) {
    m := &UrlToItemResolverBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUrlToItemResolverBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUrlToItemResolverBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.externalConnectors.itemIdResolver":
                        return NewItemIdResolver(), nil
                }
            }
        }
    }
    return NewUrlToItemResolverBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UrlToItemResolverBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UrlToItemResolverBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UrlToItemResolverBase) GetOdataType()(*string) {
    return m.odataType
}
// GetPriority gets the priority property value. The priority which defines the sequence in which the urlToItemResolverBase instances are evaluated.
func (m *UrlToItemResolverBase) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *UrlToItemResolverBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *UrlToItemResolverBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UrlToItemResolverBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPriority sets the priority property value. The priority which defines the sequence in which the urlToItemResolverBase instances are evaluated.
func (m *UrlToItemResolverBase) SetPriority(value *int32)() {
    m.priority = value
}
