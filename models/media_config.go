package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaConfig 
type MediaConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The removeFromDefaultAudioGroup property
    removeFromDefaultAudioGroup *bool
}
// NewMediaConfig instantiates a new mediaConfig and sets the default values.
func NewMediaConfig()(*MediaConfig) {
    m := &MediaConfig{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.appHostedMediaConfig":
                        return NewAppHostedMediaConfig(), nil
                    case "#microsoft.graph.serviceHostedMediaConfig":
                        return NewServiceHostedMediaConfig(), nil
                }
            }
        }
    }
    return NewMediaConfig(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaConfig) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["removeFromDefaultAudioGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRemoveFromDefaultAudioGroup)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MediaConfig) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoveFromDefaultAudioGroup gets the removeFromDefaultAudioGroup property value. The removeFromDefaultAudioGroup property
func (m *MediaConfig) GetRemoveFromDefaultAudioGroup()(*bool) {
    return m.removeFromDefaultAudioGroup
}
// Serialize serializes information the current object
func (m *MediaConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("removeFromDefaultAudioGroup", m.GetRemoveFromDefaultAudioGroup())
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
func (m *MediaConfig) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaConfig) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoveFromDefaultAudioGroup sets the removeFromDefaultAudioGroup property value. The removeFromDefaultAudioGroup property
func (m *MediaConfig) SetRemoveFromDefaultAudioGroup(value *bool)() {
    m.removeFromDefaultAudioGroup = value
}
