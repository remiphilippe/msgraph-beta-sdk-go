package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MediaConfig provides operations to manage the commsApplication singleton.
type MediaConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    removeFromDefaultAudioGroup *bool;
}
// NewMediaConfig instantiates a new mediaConfig and sets the default values.
func NewMediaConfig()(*MediaConfig) {
    m := &MediaConfig{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaConfigFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMediaConfig(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaConfig) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaConfig) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["removeFromDefaultAudioGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveFromDefaultAudioGroup(val)
        }
        return nil
    }
    return res
}
// GetRemoveFromDefaultAudioGroup gets the removeFromDefaultAudioGroup property value. 
func (m *MediaConfig) GetRemoveFromDefaultAudioGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.removeFromDefaultAudioGroup
    }
}
func (m *MediaConfig) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MediaConfig) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.additionalData = value
    }
}
// SetRemoveFromDefaultAudioGroup sets the removeFromDefaultAudioGroup property value. 
func (m *MediaConfig) SetRemoveFromDefaultAudioGroup(value *bool)() {
    if m != nil {
        m.removeFromDefaultAudioGroup = value
    }
}
