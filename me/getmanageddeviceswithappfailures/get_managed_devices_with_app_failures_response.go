package getmanageddeviceswithappfailures

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetManagedDevicesWithAppFailuresResponse provides operations to call the getManagedDevicesWithAppFailures method.
type GetManagedDevicesWithAppFailuresResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value []string;
}
// NewGetManagedDevicesWithAppFailuresResponse instantiates a new getManagedDevicesWithAppFailuresResponse and sets the default values.
func NewGetManagedDevicesWithAppFailuresResponse()(*GetManagedDevicesWithAppFailuresResponse) {
    m := &GetManagedDevicesWithAppFailuresResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetManagedDevicesWithAppFailuresResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetManagedDevicesWithAppFailuresResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetManagedDevicesWithAppFailuresResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *GetManagedDevicesWithAppFailuresResponse) GetValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *GetManagedDevicesWithAppFailuresResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetManagedDevicesWithAppFailuresResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        err := writer.WriteCollectionOfStringValues("value", m.GetValue())
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
func (m *GetManagedDevicesWithAppFailuresResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *GetManagedDevicesWithAppFailuresResponse) SetValue(value []string)() {
    if m != nil {
        m.value = value
    }
}
